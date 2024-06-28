package debug

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Debug(e *gin.Engine) {
	e.GET("/debug", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status":  http.StatusOK,
				"message": "debug",
			},
		)
	})
}
