package repositories

import (
	"fmt"

	"github.com/shaikrasheed99/golang-user-jwt-authentication/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(*models.User) (models.User, error)
	FindUserByUsername(username string) (models.User, error)
	FindAllUsers() ([]models.User, error)
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
	fmt.Println("[FindUserByUsernamRepository] Hitting find user details by username function in user repository")

	var user models.User
	res := ur.db.Where("username = ?", username).Find(&user)
	if res.Error != nil {
		fmt.Println("[FindUserByUsernamRepository] Error while finding user by username")
		return models.User{}, res.Error
	}

	if res.RowsAffected == 0 {
		fmt.Println("[FindUserByUsernamRepository] User is not found with username")
		return models.User{}, gorm.ErrRecordNotFound
	}

	fmt.Println("[FindUserByUsernamRepository] Found user deatils of a user by username")
	return user, nil
}

func (ur *userRepository) FindAllUsers() ([]models.User, error) {
	fmt.Println("[FindAllUsersRepository] Hitting find all user function in user repository")

	var users []models.User
	res := ur.db.Find(&users)
	if res.Error != nil {
		fmt.Println("[FindAllUsersRepository] Error while finding list of users")
		return nil, res.Error
	}

	fmt.Println("[FindAllUsersRepository] Found list of users")
	return users, nil
}
