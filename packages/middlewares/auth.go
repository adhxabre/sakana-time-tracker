package middlewares

import (
	dtos "absence-click/dtos/result"
	"absence-click/packages/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(
				http.StatusUnauthorized,
				dtos.ErrorResult{
					Code:    http.StatusUnauthorized,
					Message: "unauthorized",
				},
			)
			return
		}

		token = strings.SplitN(token, " ", 2)[1]
		claims, err := jwt.DecodeToken(token)

		if err != nil {
			c.JSON(
				http.StatusUnauthorized,
				dtos.ErrorResult{
					Code:    http.StatusUnauthorized,
					Message: "unauthorized",
				},
			)
			return
		}

		c.Set("user_login", claims)
	}
}
