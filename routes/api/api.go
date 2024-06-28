package api

import "github.com/gin-gonic/gin"

func API(e *gin.Engine) {
	routes := e.Group("/api")
	{
		UserRoutes(routes)
	}
}
