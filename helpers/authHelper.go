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
	accessTokenClaims := &authClaims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "user1",
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Minute * 1),
			},
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims).SignedString([]byte("thisismysecretkey"))
	if err != nil {
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
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func ValidateToken(signedToken string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&authClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
				message := fmt.Sprintln("invalid token ", token.Header["alg"])
				return nil, errors.New(message)
			}
			return []byte("thisismysecretkey"), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*authClaims)
	if !ok {
		return nil, err
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("token has expired")
	}

	return token, nil
}
