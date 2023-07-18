package repositories

import (
	"fmt"
	"time"

	"github.com/shaikrasheed99/golang-user-jwt-authentication/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	SaveTokens(string, string, string) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	fmt.Println("[NewAuthRepository] Initiating New Auth Repository")
	return &authRepository{
		db: db,
	}
}

func (ar *authRepository) SaveTokens(username string, accessToken string, refreshToken string) error {
	fmt.Println("[SaveTokensRepository] Hitting save function in auth repository")

	tokens := models.Tokens{
		Username:     username,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		CreatedAt:    time.Now(),
	}

	res := ar.db.Save(&tokens)
	if res.Error != nil {
		fmt.Println("[SaveTokensRepository] Error while saving the tokens")
		return res.Error
	}

	fmt.Println("[SaveTokensRepository] Tokens are saved")
	return nil
}
