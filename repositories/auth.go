package repositories

import (
	"absence-click/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepositories interface {
	Register(user models.User) (models.User, error)
	Login(email string) (models.User, error)
	CheckAuth(ID int) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) Login(username string) (models.User, error) {
	var user models.User
	err := r.db.Where("username=?", username).First(&user).Error

	return user, err
}

func (r *repository) CheckAuth(ID uuid.UUID) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", ID).First(&user).Error

	return user, err
}
