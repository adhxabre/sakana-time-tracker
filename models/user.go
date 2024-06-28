package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" form:"-" gorm:"type:uuid;default:uuid_generate_v4();primary_key;column:id"`
	Name      string    `json:"name" form:"name" gorm:"type:varchar(255);column:name"`
	Email     string    `json:"email" form:"email" gorm:"type:varchar(255);column:email"`
	Username  string    `json:"username" form:"username" gorm:"type:varchar(255);column:username"`
	Password  string    `json:"-" form:"password" gorm:"type:varchar(255);column:password"`
	CreatedAt time.Time `json:"-" form:"-" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"-" form:"-" gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "users"
}
