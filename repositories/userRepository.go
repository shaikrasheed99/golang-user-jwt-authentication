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
	fmt.Println("[NewUserRepository] Initiating New User Repository")
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Save(user *models.User) (models.User, error) {
	fmt.Println("[SaveRepository] Hitting save function in user repository")

	res := ur.db.Create(&user)
	if res.Error != nil {
		fmt.Println("[SaveRepository] Error while saving the user")
		return models.User{}, res.Error
	}

	fmt.Println("[SaveRepository] User details have created")
	return *user, nil
}

func (ur *userRepository) FindUserByUsername(username string) (models.User, error) {
	fmt.Println("[FindUserByUsername] Hitting find user details by username function in user repository")

	var user models.User
	res := ur.db.Where("username = ?", username).Find(&user)
	if res.Error != nil {
		fmt.Println("[FindUserByUsername] Error while finding user by username")
		return models.User{}, res.Error
	}

	if res.RowsAffected == 0 {
		fmt.Println("[FindUserByUsername] User is not found with username")
		return models.User{}, gorm.ErrRecordNotFound
	}

	fmt.Println("[FindUserByUsername] Found user deatils of a user by username")
	return user, nil
}
