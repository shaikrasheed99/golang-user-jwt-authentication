package helpers

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/constants"
)

type authClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(username string, role string) (string, string, error) {
	fmt.Println("[GenerateTokenHelper] Generating access and refresh tokens")

	accessTokenClaims := &authClaims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "user1",
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Minute * 5),
			},
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims).SignedString([]byte("thisismysecretkey"))
	if err != nil {
		fmt.Println("[GenerateTokenHelper]", err.Error())
		return "", "", err
	}

	refreshTokenClaims := &authClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Hour * 1),
			},
		},
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims).SignedString([]byte("thisismysecretkey"))
	if err != nil {
		fmt.Println("[GenerateTokenHelper]", err.Error())
		return "", "", err
	}

	fmt.Println("[GenerateTokenHelper] Generated access and refresh tokens")
	return accessToken, refreshToken, nil
}

func ValidateToken(signedToken string) (*authClaims, error) {
	fmt.Println("[ValidateTokenHelper] Validating access token")

	token, err := jwt.ParseWithClaims(
		signedToken,
		&authClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
				message := fmt.Sprintln("invalid token ", token.Header["alg"])
				fmt.Println("[ValidateToken]", message)
				return nil, errors.New(message)
			}
			return []byte("thisismysecretkey"), nil
		},
	)

	if err != nil || !token.Valid {
		fmt.Println("[ValidateTokenHelper]", err.Error())
		return nil, err
	}

	claims, ok := token.Claims.(*authClaims)
	if !ok {
		errMessage := constants.InvalidTokenErrorMessage
		fmt.Println("[ValidateTokenHelper]", errMessage)
		return nil, errors.New(errMessage)
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		errMessage := constants.ExpiredTokenErrorMessage
		fmt.Println("[ValidateTokenHelper]", errMessage)
		return nil, errors.New(errMessage)
	}

	fmt.Println("[ValidateTokenHelper] Validated access token")
	return claims, nil
}
