package jwt

import (
	"net/http"
	"time"

	"github.com/Infinity-OJ/Server/internal/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = 200
		token := c.Query("token")
		if token == "" {
			code = 400
		} else {
			claims, err := jwt.ParseToken(token)
			if err != nil {
				code = 401
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 401
			}
			c.Set("claims", claims)
		}

		if code != 200 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "GG",
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
