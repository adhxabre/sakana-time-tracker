package api

import (
	"absence-click/handlers"
	"absence-click/packages/databases"
	"absence-click/repositories"

	"github.com/gin-gonic/gin"
)

func UserRoutes(e *gin.RouterGroup) {
	userRepository := repositories.RepositoryUser(databases.DB)
	h := handlers.HandlerUser(userRepository)

	e.GET("/users", h.FindUsers)
	idUsage := e.Group("/users")
	{
		idUsage.GET("/:id", h.GetUser)
		idUsage.PATCH("/:id", h.UpdateEmail)
		idUsage.PATCH("/:id", h.UpdateName)
		idUsage.PATCH("/:id", h.UpdatePassword)
		idUsage.DELETE("/:id", h.DeleteUser)
	}
}
