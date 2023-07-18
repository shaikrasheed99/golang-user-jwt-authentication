package services

import (
	"fmt"

	"github.com/shaikrasheed99/golang-user-jwt-authentication/repositories"
)

type AuthService interface {
	SaveTokensByUsername(string, string, string) error
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
