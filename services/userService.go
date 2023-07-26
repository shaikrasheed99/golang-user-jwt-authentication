package services

import (
	"errors"
	"fmt"

	"github.com/shaikrasheed99/golang-user-jwt-authentication/constants"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/helpers"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/models"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/repositories"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/requests"
	"gorm.io/gorm"
)

type IUserService interface {
	Save(*requests.SignupRequest) (*models.User, error)
	Login(*requests.LoginRequest) (*models.User, error)
	UserByUsername(username string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
}

type userService struct {
	ur repositories.IUserRepository
}

func NewUserService(ur repositories.IUserRepository) IUserService {
	fmt.Println("[NewUserService] Initiating New User Service")
	return &userService{
		ur: ur,
	}
}

func (us *userService) Save(userReq *requests.SignupRequest) (*models.User, error) {
	fmt.Println("[SaveService] Hitting save function in user service")

	hashedPass, err := helpers.GenerateHashedPassword(userReq.Password)
	if err != nil {
		fmt.Println("[SaveService]", err.Error())
		return nil, err
	}

	_, err = us.ur.FindUserByUsername(userReq.Username)

	if err != gorm.ErrRecordNotFound {
		errMessage := constants.ErrUserAlreadyExists
		fmt.Println("[SaveService]", errMessage)
		return nil, errors.New(errMessage)
	}

	newUser := &models.User{
		FirstName: userReq.FirstName,
		LastName:  userReq.LastName,
		Username:  userReq.Username,
		Password:  hashedPass,
		Email:     userReq.Email,
		Role:      constants.User,
	}

	savedUser, err := us.ur.Save(newUser)
	if err != nil {
		fmt.Println("[SaveService]", err.Error())
		return nil, err
	}

	fmt.Println("[SaveService] Returned saved user deatils from repository")
	return &savedUser, nil
}

func (us *userService) Login(userReq *requests.LoginRequest) (*models.User, error) {
	fmt.Println("[LoginService] Hitting login function in user service")

	user, err := us.ur.FindUserByUsername(userReq.Username)
	if err == gorm.ErrRecordNotFound {
		errMessage := constants.ErrUserNotFound
		fmt.Println("[LoginService]", errMessage)
		return nil, errors.New(errMessage)
	}

	if err != nil {
		fmt.Println("[LoginService]", err.Error())
		return nil, err
	}

	isValidPassword := helpers.CheckPassword(user.Password, userReq.Password)

	if !isValidPassword {
		errMessage := constants.ErrWrongPassword
		fmt.Println("[LoginService]", errMessage)
		return nil, errors.New(errMessage)
	}

	fmt.Println("[LoginService] Returned logged in user deatils from repository")
	return &user, nil
}

func (us *userService) UserByUsername(username string) (*models.User, error) {
	fmt.Println("[UserByUsernameService] Hitting user by username function in user service")

	user, err := us.ur.FindUserByUsername(username)
	if err == gorm.ErrRecordNotFound {
		errMessage := constants.ErrUserNotFound
		fmt.Println("[UserByUsernameService]", errMessage)
		return nil, errors.New(errMessage)
	}

	if err != nil {
		fmt.Println("[UserByUsernameService]", err.Error())
		return nil, err
	}

	fmt.Println("[UserByUsernameService] Returned user details from repository")
	return &user, nil
}

func (us *userService) GetAllUsers() ([]models.User, error) {
	fmt.Println("[GetAllUsersService] Hitting get all users function in user service")

	users, err := us.ur.FindAllUsers()
	if err != nil {
		fmt.Println("[GetAllUsersService]", err.Error())
		return nil, err
	}

	fmt.Println("[GetAllUsersService] Returned list of users from repository")
	return users, nil
}
