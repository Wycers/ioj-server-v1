package controllers

import (
	"github.com/Infinity-OJ/Server/internal/app/server/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
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

	p, err := usersController.service.Create(c.Request.Context(), user.Username, user.Password, user.Email)
	if err != nil {
		usersController.logger.Error("register error", zap.Error(err))
		c.String(http.StatusInternalServerError, "%+v", err)
		return
	}

	c.JSON(http.StatusOK, p)
}
