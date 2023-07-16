package helpers

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type authClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(username string, role string) (string, string, error) {
	fmt.Println("[GenerateToken] Generating access and refresh tokens")

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
		fmt.Println("[GenerateToken]", err.Error())
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
		fmt.Println("[GenerateToken]", err.Error())
		return "", "", err
	}

	fmt.Println("[GenerateToken] Generated access and refresh tokens")
	return accessToken, refreshToken, nil
}

func ValidateToken(signedToken string) (*authClaims, error) {
	fmt.Println("[ValidateToken] Validating access tokens")

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
		fmt.Println("[ValidateToken]", err.Error())
		return nil, err
	}

	claims, ok := token.Claims.(*authClaims)
	if !ok {
		errMessage := "invalid token"
		fmt.Println("[ValidateToken]", errMessage)
		return nil, errors.New(errMessage)
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		errMessage := "token has expired"
		fmt.Println("[ValidateToken]", errMessage)
		return nil, errors.New(errMessage)
	}

	fmt.Println("[ValidateToken] Validated access tokens")
	return claims, nil
}
