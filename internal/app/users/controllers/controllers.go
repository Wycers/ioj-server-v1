package controllers

import (
	"github.com/Infinity-OJ/Server/internal/pkg/transports/http"
	"github.com/Infinity-OJ/Server/internal/pkg/transports/http/middlewares/jwt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func CreateInitControllersFn(pc *UsersController) http.InitControllers {
	return func(r *gin.Engine) {
		r.POST("/session", pc.CreateSession)
		r.GET("/session", jwt.JWT(), pc.GetSession)
	}
}

var ProviderSet = wire.NewSet(NewUsersController, CreateInitControllersFn)
