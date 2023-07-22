package services

import (
	"errors"
	"fmt"

	"github.com/shaikrasheed99/golang-user-jwt-authentication/constants"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/models"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/repositories"
	"gorm.io/gorm"
)

type AuthService interface {
	SaveTokensByUsername(string, string, string) error
	FindTokensByUsername(string) (models.Tokens, error)
	DeleteTokensByUsername(string) error
}

type authService struct {
	ar repositories.AuthRepository
}

func NewAuthService(ar repositories.AuthRepository) AuthService {
	fmt.Println("[NewAuthService] Initiating New Auth Service")
	return &authService{
		ar: ar,
	}
}

func (as *authService) SaveTokensByUsername(username string, accessToken string, refreshToken string) error {
	fmt.Println("[SaveTokensByUsernameService] Hitting save tokens function in auth service")

	err := as.ar.SaveTokens(username, accessToken, refreshToken)
	if err != nil {
		fmt.Println("[SaveTokensByUsernameService]", err.Error())
		return err
	}

	fmt.Println("[SaveTokensByUsernameService] Tokens are saved")
	return nil
}

func (as *authService) FindTokensByUsername(username string) (models.Tokens, error) {
	fmt.Println("[FindTokensByUsernameService] Hitting tokens by username function in auth service")

	tokens, err := as.ar.FindTokensByUsername(username)
	if err == gorm.ErrRecordNotFound {
		errMessage := constants.TokensNotFoundErrorMessage
		fmt.Println("[FindTokensByUsernameService]", errMessage)
		return models.Tokens{}, errors.New(errMessage)
	}

	if err != nil {
		fmt.Println("[FindTokensByUsernameService]", err.Error())
		return models.Tokens{}, err
	}

	fmt.Println("[FindTokensByUsernameService] Returned tokens from repository")
	return tokens, nil
}

func (as *authService) DeleteTokensByUsername(username string) error {
	fmt.Println("[DeleteTokensByUsername] Hitting delete tokens function in auth service")

	err := as.ar.DeleteTokensByUsername(username)
	if err != nil {
		fmt.Println("[DeleteTokensByUsername]", err.Error())
		return err
	}

	fmt.Println("[DeleteTokensByUsername] Tokens are deleted")
	return nil
}
