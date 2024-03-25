package repository

import (
	"context"
	"errors"

	"github.com/achsanit/my-gram/internal/infrastructure"
	"github.com/achsanit/my-gram/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type UserQuery interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	Login(ctx context.Context, email string, password string) (model.User, error)

	GetUserByID(ctx context.Context, userId int) (model.User, error)
}

type userQueryImpl struct {
	db infrastructure.GormPostgres
}

// GetUserByID implements UserQuery.
func (u *userQueryImpl) GetUserByID(ctx context.Context, id int) (model.User, error) {
	db := u.db.GetConnection()

	var user model.User
	result := db.WithContext(ctx).Where("id = ?", id).First(&user)
	if result.Error != nil {
		return model.User{}, errors.New("user not found")
	}

	return user, nil
}

// Login implements UserQuery.
func (u *userQueryImpl) Login(ctx context.Context, email string, password string) (model.User, error) {
	db := u.db.GetConnection()

	var user model.User
	result := db.WithContext(ctx).Where("email = ?", email).First(&user)
	if result.Error != nil {
		return model.User{}, errors.New("credential not match")
	}

	errPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errPass != nil {
		return model.User{}, errors.New("credential not match")
	}

	return user, nil
}

// Register implements UserQuery.
func (u *userQueryImpl) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	db := u.db.GetConnection()

	if err := db.
		WithContext(ctx).
		Table("users").
		Save(&user).
		Error; err != nil {

		return model.User{}, err

	}
	return user, nil
}

func NewUserQuery(db infrastructure.GormPostgres) UserQuery {
	return &userQueryImpl{
		db: db,
	}
}
