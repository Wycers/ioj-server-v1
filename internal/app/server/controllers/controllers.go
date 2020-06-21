package controllers

import (
	"github.com/infinity-oj/server/internal/pkg/transports/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func CreateInitControllersFn(uc *UsersController) http.InitControllers {
	return func(res *gin.Engine) {
		api := res.Group("/api")

		user := api.Group("/users")
		user.POST("/", uc.Post)
		user.PUT("/:id")
		//user.DELETE("/:id")

		session := api.Group("/sessions")
		session.GET("/")
		session.POST("/", uc.SignIn)
		//session.PUT("/")
		session.DELETE("/")
	}
}

var ProviderSet = wire.NewSet(NewUsersController, CreateInitControllersFn)
