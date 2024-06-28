package main

import (
	"absence-click/packages/databases"
	"absence-click/routes"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatalf("Error loading .env file: %v", errEnv.Error())
	}

	// declare port from .env
	PORT := os.Getenv("PORT")

	// initialize gin
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	// initialize databases
	databases.DBConn()
	databases.DBMigration()

	// initialize cors
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	// initialize routes
	routes.RouteInit(r)

	servErr := r.Run("localhost:" + PORT)
	if servErr != nil {
		log.Fatalf("Error starting server: %v", servErr.Error())
	}
}
