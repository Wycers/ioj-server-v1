package controllers

import (
	"github.com/Infinity-OJ/Server/internal/pkg/transports/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func CreateInitControllersFn(pc *ProductsController, uc *UsersController) http.InitControllers {
	return func(res *gin.Engine) {
		api := res.Group("/api")

		user := api.Group("/users")
		user.POST("/", uc.Post)
		user.PUT("/:id")
		//user.DELETE("/:id")

		session := api.Group("/sessions")
		session.GET("/")
		session.POST("/")
		//session.PUT("/")
		session.DELETE("/")

		res.GET("/product/:id", pc.Get)
	}
}

var ProviderSet = wire.NewSet(NewProductsController, NewUsersController, CreateInitControllersFn)
