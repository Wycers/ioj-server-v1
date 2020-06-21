package controllers

import (
	"fmt"

	"github.com/infinity-oj/server/internal/app/files/services"
	"github.com/infinity-oj/server/internal/pkg/jwt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FilesController struct {
	logger  *zap.Logger
	service services.FilesService
}

func NewFilesController(logger *zap.Logger, s services.FilesService) *FilesController {
	return &FilesController{
		logger:  logger,
		service: s,
	}
}

func (pc *FilesController) CreateSession(c *gin.Context) {
	username := c.PostForm("username")
	fmt.Println(username)
	password := c.PostForm("password")
	fmt.Println(password)
	fmt.Println(jwt.GenerateToken(username))
}

func (pc *FilesController) GetSession(c *gin.Context) {
	claims := c.MustGet("claims").(*jwt.Claims)
	fmt.Printf(claims.Username)
}
