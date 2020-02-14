package controllers

import (
	"github.com/Infinity-OJ/Server/internal/pkg/transports/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func CreateInitControllersFn(pc *UsersController) http.InitControllers {
	return func(r *gin.Engine) {
		//r.GET("/detail/:id", pc.Get)
		r.POST("/session", pc.GetSession)
	}
}

var ProviderSet = wire.NewSet(NewUsersController, CreateInitControllersFn)
