package controllers

import (
	"fmt"

	"github.com/infinity-oj/server/internal/app/judgements/services"
	"github.com/infinity-oj/server/internal/pkg/jwt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type JudgementController struct {
	logger  *zap.Logger
	service services.JudgementsService
}

func NewJudgementsController(logger *zap.Logger, s services.JudgementsService) *JudgementController {
	return &JudgementController{
		logger:  logger,
		service: s,
	}
}

func (pc *JudgementController) CreateSession(c *gin.Context) {
	username := c.PostForm("username")
	fmt.Println(username)
	password := c.PostForm("password")
	fmt.Println(password)
	fmt.Println(jwt.GenerateToken(username))
}

func (pc *JudgementController) GetSession(c *gin.Context) {
	claims := c.MustGet("claims").(*jwt.Claims)
	fmt.Printf(claims.Username)
}
