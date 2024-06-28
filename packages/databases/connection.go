package databases

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConn() {
	var err error

	var DBHost = os.Getenv("DB_HOST")
	var DBPort = os.Getenv("DB_PORT")
	var DBUser = os.Getenv("DB_USER")
	var DBPass = os.Getenv("DB_PASS")
	var DBName = os.Getenv("DB_NAME")

	DSN := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", DBHost, DBPort, DBUser, DBPass, DBName)
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("[DB] Database connected successfully.")
}
