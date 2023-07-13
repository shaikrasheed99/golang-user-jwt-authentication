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
	Login(*requests.LoginRequest) (*models.User, error)
	UserByUsername(username string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
}

type userService struct {
	ur repositories.UserRepository
}

func NewUserService(ur repositories.UserRepository) UserService {
	fmt.Println("[NewUserService] Initiating New User Service")
	return &userService{
		ur: ur,
	}
}

func (us *userService) Save(userReq *requests.SignupRequest) (*models.User, error) {
	fmt.Println("[SaveService] Hitting save function in user service")

	hashedPass, err := utils.GenerateHashedPassword(userReq.Password)
	if err != nil {
		fmt.Println("[SaveService]", err.Error())
		return nil, err
	}

	_, err = us.ur.FindUserByUsername(userReq.Username)

	if err != gorm.ErrRecordNotFound {
		fmt.Println("[SaveService] User is already exists with username")
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
		fmt.Println("[LoginService] User is not found with username")
		return nil, errors.New("user is not found with username")
	}

	if err != nil {
		fmt.Println("[LoginService] Error while fetching user details with username")
		return nil, err
	}

	isValidPassword := utils.CheckPassword(user.Password, userReq.Password)

	if !isValidPassword {
		fmt.Println("[LoginService] Password is wrong")
		return nil, errors.New("password is wrong")
	}

	fmt.Println("[LoginService] Returned logined user deatils from repository")
	return &user, nil
}

func (us *userService) UserByUsername(username string) (*models.User, error) {
	fmt.Println("[UserByUsernameService] Hitting user by username function in user service")

	user, err := us.ur.FindUserByUsername(username)
	if err == gorm.ErrRecordNotFound {
		fmt.Println("[UserByUsernameService] User is not found with username")
		return nil, errors.New("user is not found with username")
	}

	if err != nil {
		fmt.Println("[UserByUsernameService] Error while fetching user details with username")
		return nil, err
	}

	fmt.Println("[UserByUsernameService] Returned user details by username from repository")
	return &user, nil
}

func (us *userService) GetAllUsers() ([]models.User, error) {
	fmt.Println("[GetAllUsersService] Hitting get all users function in user service")

	users, err := us.ur.FindAllUsers()
	if err != nil {
		fmt.Println("[GetAllUsersService] Error while fetching list of users")
		return nil, err
	}

	fmt.Println("[GetAllUsersService] Returned list of users from repository")
	return users, nil
}
