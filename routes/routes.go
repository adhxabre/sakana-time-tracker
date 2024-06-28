package routes

import (
	dtos "absence-click/dtos/result"
	"absence-click/routes/api"
	"absence-click/routes/debug"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouteInit(e *gin.Engine) {
	e.NoRoute(func(c *gin.Context) {
		c.JSON(
			http.StatusNotFound,
			dtos.ErrorResult{
				Code:    http.StatusNotFound,
				Message: "Routes Not Found.",
			},
		)
	})

	e.HandleMethodNotAllowed = true
	e.NoMethod(func(c *gin.Context) {
		c.JSON(
			http.StatusMethodNotAllowed,
			dtos.ErrorResult{
				Code:    http.StatusMethodNotAllowed,
				Message: "Method Not Allowed.",
			},
		)
	})

	// debug routes
	debug.Debug(e)
	api.API(e)
}
