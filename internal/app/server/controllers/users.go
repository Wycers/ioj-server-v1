package controllers

import (
	"net/http"

	"github.com/Infinity-OJ/Server/internal/app/server/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UsersController struct {
	logger  *zap.Logger
	service services.UserService
}

func NewUsersController(logger *zap.Logger, s services.UserService) *UsersController {
	return &UsersController{
		logger:  logger,
		service: s,
	}
}

type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Email    string `form:"email"`
}

func (usersController *UsersController) Post(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		usersController.logger.Error("Missing fields", zap.Error(err))
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	p, err := usersController.service.CreateUser(c.Request.Context(), user.Username, user.Password, user.Email)
	if err != nil {
		usersController.logger.Error("register error", zap.Error(err))
		c.String(http.StatusInternalServerError, "%+v", err)
		return
	}

	c.JSON(http.StatusOK, p)
}

type Session struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (usersController *UsersController) SignIn(c *gin.Context) {
	var session Session
	if err := c.ShouldBind(&session); err != nil {
		usersController.logger.Error("Missing fields", zap.Error(err))
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	p, err := usersController.service.CreateSession(c.Request.Context(), session.Username, session.Password)
	if err != nil {
		usersController.logger.Error("sign in error", zap.Error(err))
		//c.String(http.StatusInternalServerError, "%+v", err)
		c.String(http.StatusInternalServerError, "")
		return
	}

	if p == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"token": nil})
	} else {
		c.JSON(http.StatusOK, gin.H{"token": p})
	}
}
