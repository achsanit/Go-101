package router

import (
	"github.com/achsanit/my-gram/internal/handler"
	"github.com/achsanit/my-gram/internal/middleware"
	"github.com/gin-gonic/gin"
)

type PhotoRouter interface {
	Mount()
}

type photoRouterImpl struct {
	v *gin.RouterGroup
	h handler.PhotoHandler
}

// Mount implements PhotoRouter.
func (p *photoRouterImpl) Mount() {
	p.v.Use(middleware.CheckAuthBearer)

	p.v.POST("", p.h.PostPhoto)
	p.v.GET("", p.h.GetPhotosUser)

	p.v.GET("/:id", p.h.GetPhotoByID)
}

func NewPhotoRouter(v *gin.RouterGroup, h handler.PhotoHandler) PhotoRouter {
	return &photoRouterImpl{
		v: v,
		h: h,
	}
}
