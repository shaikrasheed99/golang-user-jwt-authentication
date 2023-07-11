package repositories

import (
	"fmt"

	"github.com/shaikrasheed99/golang-user-jwt-authentication/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(*models.User) (models.User, error)
	FindUserByUsername(username string) (models.User, error)
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
		fmt.Println("[SaveRepository] Error while saving the user")
		return models.User{}, res.Error
	}

	return *user, nil
}

func (ur *userRepository) FindUserByUsername(username string) (models.User, error) {
	var user models.User

	res := ur.db.Where("username = ?", username).Find(&user)
	if res.Error != nil {
		fmt.Println("[SaveRepository] Error while finding user by username")
		return models.User{}, res.Error
	}

	if res.RowsAffected == 0 {
		fmt.Println("[SaveRepository] No users found by username")
		return models.User{}, gorm.ErrRecordNotFound
	}

	return user, nil
}
