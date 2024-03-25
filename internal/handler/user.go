package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/achsanit/my-gram/internal/model"
	"github.com/achsanit/my-gram/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	UserRegister(ctx *gin.Context)
	UserLogin(ctx *gin.Context)

	GetUserByID(ctx *gin.Context)
}

type userHandlerImpl struct {
	svc service.UserService
}

// GetUserByID implements UserHandler.
func (uh *userHandlerImpl) GetUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := uh.svc.GetUserByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user":    user,
		"message": "ok",
	})
}

// UserLogin implements UserHandler.
func (uh *userHandlerImpl) UserLogin(ctx *gin.Context) {
	userLogin := model.UserLogin{}

	if err := ctx.Bind(&userLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := uh.svc.Login(ctx, userLogin.Email, userLogin.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	token, err := uh.svc.GenerateUserAccessToken(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}

// UserRegister implements UserHandler.
func (uh *userHandlerImpl) UserRegister(ctx *gin.Context) {

	userRegist := model.UserRegister{}

	if err := ctx.Bind(&userRegist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := userRegist.ValidateInput(); err != nil {
		var errData map[string]string
		err := json.Unmarshal([]byte(err.Error()), &errData)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"message": errData})
		return
	}

	user, err := uh.svc.Register(ctx, userRegist)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	token, err := uh.svc.GenerateUserAccessToken(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}

func NewUserHandler(service service.UserService) UserHandler {
	return &userHandlerImpl{
		svc: service,
	}
}
