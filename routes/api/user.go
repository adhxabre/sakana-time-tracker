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
		idUsage.PATCH("/email/:id", h.UpdateEmail)
		idUsage.PATCH("/name/:id", h.UpdateName)
		idUsage.PATCH("/password/:id", h.UpdatePassword)
		idUsage.DELETE("/:id", h.DeleteUser)
	}
}
