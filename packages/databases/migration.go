package databases

import (
	"absence-click/models"
	"fmt"
)

func DBMigration() {
	err := DB.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		fmt.Println(err)
		panic("[DB] Migration Failed")
	}

	fmt.Println("[DB] Migration run successfully.")
}
