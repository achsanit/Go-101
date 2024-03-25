package router

import (
	"github.com/achsanit/my-gram/internal/handler"
	"github.com/gin-gonic/gin"
)

type UserRouter interface {
	Mount()
}

type userRouterImpl struct {
	version *gin.RouterGroup
	handler handler.UserHandler
}

func NewUserRouter(v *gin.RouterGroup, h handler.UserHandler) UserRouter {
	return &userRouterImpl{
		version: v,
		handler: h,
	}
}

func (u *userRouterImpl) Mount() {
	// User register
	u.version.POST("/register", u.handler.UserRegister)
	u.version.POST("/login", u.handler.UserLogin)

	u.version.GET("/:id", u.handler.GetUserByID)
}
