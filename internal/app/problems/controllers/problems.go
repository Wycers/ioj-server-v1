package controllers

import (
	"fmt"

	"github.com/infinity-oj/server/internal/app/problems/services"
	"github.com/infinity-oj/server/internal/pkg/jwt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProblemController struct {
	logger  *zap.Logger
	service services.ProblemsService
}

func NewProblemsController(logger *zap.Logger, s services.ProblemsService) *ProblemController {
	return &ProblemController{
		logger:  logger,
		service: s,
	}
}

func (pc *ProblemController) CreateSession(c *gin.Context) {
	username := c.PostForm("username")
	fmt.Println(username)
	password := c.PostForm("password")
	fmt.Println(password)
	fmt.Println(jwt.GenerateToken(username))
}

func (pc *ProblemController) GetSession(c *gin.Context) {
	claims := c.MustGet("claims").(*jwt.Claims)
	fmt.Printf(claims.Username)
}
