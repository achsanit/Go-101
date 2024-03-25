package service

import (
	"context"
	"fmt"
	"time"

	"github.com/achsanit/my-gram/internal/model"
	"github.com/achsanit/my-gram/internal/repository"
	"github.com/achsanit/my-gram/pkg/helper"
)

type UserService interface {
	Register(ctx context.Context, user model.UserRegister) (model.User, error)
	Login(ctx context.Context, email string, password string) (model.User, error)

	GetUserByID(ctx context.Context, id int) (model.User, error)

	GenerateUserAccessToken(ctx context.Context, user model.User) (accessToken string, err error)
}

type userServiceImpl struct {
	repo repository.UserQuery
}

// GetUserByID implements UserService.
func (u *userServiceImpl) GetUserByID(ctx context.Context, id int) (model.User, error) {
	user, err := u.repo.GetUserByID(ctx, id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// Login implements UserService.
func (u *userServiceImpl) Login(ctx context.Context, email string, password string) (model.User, error) {
	user, err := u.repo.Login(ctx, email, password)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// Register implements UserService.
func (u *userServiceImpl) Register(ctx context.Context, user model.UserRegister) (model.User, error) {
	// TODO: check email first
	pass, err := helper.GenerateHash(user.Password)

	if err != nil {
		return model.User{}, err
	}

	newUser := model.User{
		Username: user.Username,
		Email:    user.Email,
		DOB:      user.DOB,
		Password: pass,
	}

	userCreated, err := u.repo.CreateUser(ctx, newUser)

	if err != nil {
		return model.User{}, err
	}

	return userCreated, nil
}

// GenerateUserAccessToken implements UserService.
func (*userServiceImpl) GenerateUserAccessToken(ctx context.Context, user model.User) (accessToken string, err error) {
	now := time.Now()

	claim := model.StandardClaim{
		Jti: fmt.Sprintf("%v", time.Now().UnixNano()),
		Iss: "my-gram",
		Aud: "my-gram",
		Sub: "access-token",
		Exp: uint64(now.Add(time.Hour).Unix()),
		Iat: uint64(now.Unix()),
		Nbf: uint64(now.Unix()),
	}

	userClaim := model.AccessClaim{
		StandardClaim: claim,
		UserID:        user.ID,
		Dob:           user.DOB,
		Username:      user.Username,
	}

	accessToken, err = helper.GenerateToken(userClaim)
	return
}

func NewUserService(repo repository.UserQuery) UserService {
	return &userServiceImpl{
		repo: repo,
	}
}
