package services

import (
	"errors"
	"fmt"

	"github.com/shaikrasheed99/golang-user-jwt-authentication/models"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/repositories"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/requests"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/utils"
	"gorm.io/gorm"
)

type UserService interface {
	Save(*requests.SignupRequest) (*models.User, error)
}

type userService struct {
	ur repositories.UserRepository
}

func NewUserService(ur repositories.UserRepository) UserService {
	return &userService{
		ur: ur,
	}
}

func (us *userService) Save(userReq *requests.SignupRequest) (*models.User, error) {
	hashedPass, err := utils.GenerateHashedPassword(userReq.Password)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	_, err = us.ur.FindUserByUsername(userReq.Username)

	if err != gorm.ErrRecordNotFound {
		fmt.Println("User is already exists with username")
		return nil, errors.New("user is already exists with username")
	}

	newUser := &models.User{
		FirstName: userReq.FirstName,
		LastName:  userReq.LastName,
		Username:  userReq.Username,
		Password:  hashedPass,
		Email:     userReq.Email,
	}

	savedUser, err := us.ur.Save(newUser)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &savedUser, nil
}
