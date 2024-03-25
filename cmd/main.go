package main

import (
	"github.com/achsanit/my-gram/internal/handler"
	"github.com/achsanit/my-gram/internal/infrastructure"
	"github.com/achsanit/my-gram/internal/repository"
	"github.com/achsanit/my-gram/internal/router"
	"github.com/achsanit/my-gram/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	g.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "okeyy",
		})
	})

	postgres := infrastructure.NewGormPostgres()

	v1 := g.Group("/v1")
	{
		users := v1.Group("/users")
		{
			repo := repository.NewUserQuery(postgres)
			service := service.NewUserService(repo)
			handler := handler.NewUserHandler(service)
			router := router.NewUserRouter(users, handler)

			router.Mount()
		}

		photos := v1.Group("/photos")
		{
			repo := repository.NewPhotoQuery(postgres)
			service := service.NewPhotoService(repo)
			handler := handler.NewPhotoHandler(service)
			router := router.NewPhotoRouter(photos, handler)

			router.Mount()
		}
	}

	g.Run(":8001")
}
