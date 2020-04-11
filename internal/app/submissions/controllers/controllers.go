package controllers

import (
	"github.com/Infinity-OJ/Server/internal/pkg/transports/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func CreateInitControllersFn(pc *SubmissionController) http.InitControllers {
	return func(r *gin.Engine) {
		//r.POST("/session", pc.CreateSession)
		//r.GET("/session", jwt.JWT(), pc.GetSession)
	}
}

var ProviderSet = wire.NewSet(NewSubmissionsController, CreateInitControllersFn)
