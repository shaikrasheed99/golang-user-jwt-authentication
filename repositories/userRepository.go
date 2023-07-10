package repositories

import (
	"fmt"

	"github.com/shaikrasheed99/golang-user-jwt-authentication/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(*models.User) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Save(user *models.User) (models.User, error) {
	res := ur.db.Create(&user)
	if res.Error != nil {
		fmt.Println("Error while saving the user")
		return models.User{}, res.Error
	}

	return *user, nil
}
