package controllers

import (
	"fmt"
	"github.com/Infinity-OJ/Server/internal/app/users/services"
	"github.com/Infinity-OJ/Server/internal/pkg/jwt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UsersController struct {
	logger  *zap.Logger
	service services.UsersService
}

func NewUsersController(logger *zap.Logger, s services.UsersService) *UsersController {
	return &UsersController{
		logger:  logger,
		service: s,
	}
}

func (pc *UsersController) GetSession(c *gin.Context) {
	username := c.PostForm("username")
	fmt.Println(username)
	password := c.PostForm("password")
	fmt.Println(password)
	fmt.Println(jwt.GenerateToken(username))
}

func (pc *UsersController) CheckSession(c *gin.Context) {
	claims := c.MustGet("claims").(*jwt.Claims)
	fmt.Printf(claims.Username)
}
