package repositories

import (
	"fmt"
	"time"

	"github.com/shaikrasheed99/golang-user-jwt-authentication/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	SaveTokens(string, string, string) error
	FindTokensByUsername(string) (models.Tokens, error)
	DeleteTokensByUsername(string) error
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
		fmt.Println("[SaveTokensRepository]", res.Error.Error())
		return res.Error
	}

	fmt.Println("[SaveTokensRepository] Tokens are saved")
	return nil
}

func (ar *authRepository) FindTokensByUsername(username string) (models.Tokens, error) {
	fmt.Println("[FindTokensByUsernameRepository] Hitting find tokens by username function in auth repository")

	var tokens models.Tokens
	res := ar.db.Where("username = ?", username).Find(&tokens)
	if res.Error != nil {
		fmt.Println("[FindTokensByUsernameRepository] Error while finding tokens by username")
		return models.Tokens{}, res.Error
	}

	if res.RowsAffected == 0 {
		fmt.Println("[FindTokensByUsernameRepository] Tokens are not found with username")
		return models.Tokens{}, gorm.ErrRecordNotFound
	}

	fmt.Println("[FindTokensByUsernameRepository] Tokens are found")
	return tokens, nil
}

func (ar *authRepository) DeleteTokensByUsername(username string) error {
	fmt.Println("[DeleteTokensByUsername] Hitting delete function in auth repository")

	res := ar.db.Where("username = ?", username).Delete(&models.Tokens{})
	if res.Error != nil {
		fmt.Println("[DeleteTokensByUsername]", res.Error.Error())
		return res.Error
	}

	fmt.Println("[DeleteTokensByUsername] Tokens are deleted")
	return nil
}
