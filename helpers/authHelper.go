package helpers

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/configs"
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
			Issuer: configs.JWT_ISSUER,
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Minute * time.Duration(configs.JWT_ACCESS_TOKEN_EXPIRATION_IN_MINUTES)),
			},
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims).SignedString([]byte(configs.JWT_SECRET))
	if err != nil {
		fmt.Println("[GenerateTokenHelper]", err.Error())
		return "", "", err
	}

	refreshTokenClaims := &authClaims{
		Username: accessTokenClaims.Username,
		Role:     accessTokenClaims.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: accessTokenClaims.Issuer,
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Minute * time.Duration(configs.JWT_REFRESH_TOKEN_EXPIRATION_IN_MINUTES)),
			},
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
		},
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims).SignedString([]byte(configs.JWT_SECRET))
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
			return []byte(configs.JWT_SECRET), nil
		},
	)

	if err != nil || !token.Valid {
		fmt.Println("[ValidateTokenHelper]", err.Error())
		return nil, err
	}

	claims, ok := token.Claims.(*authClaims)
	if !ok {
		errMessage := constants.ErrInvalidToken
		fmt.Println("[ValidateTokenHelper]", errMessage)
		return nil, errors.New(errMessage)
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		errMessage := constants.ErrExpiredToken
		fmt.Println("[ValidateTokenHelper]", errMessage)
		return nil, errors.New(errMessage)
	}

	fmt.Println("[ValidateTokenHelper] Validated access token")
	return claims, nil
}

func AreTokensEqual(tokenOne string, tokenTwo string) bool {
	return tokenOne == tokenTwo
}

func IsUserMatchesWith(c *gin.Context, inputUsername string) bool {
	if IsAdmin(c) {
		return true
	}

	if !IsUser(c) {
		return false
	}

	username := c.GetString(constants.Username)
	if IsEmpty(username) || !IsEqual(username, inputUsername) {
		return false
	}

	return true
}

func IsAdmin(c *gin.Context) bool {
	role := c.GetString(constants.Role)
	if IsEmpty(role) || !IsEqual(role, constants.Admin) {
		return false
	}

	return true
}

func IsUser(c *gin.Context) bool {
	role := c.GetString(constants.Role)
	if IsEmpty(role) || !IsEqual(role, constants.User) {
		return false
	}

	return true
}

func IsEmpty(value string) bool {
	return value == ""
}

func IsEqual(valueOne string, valueTwo string) bool {
	valueOne = strings.ToLower(valueOne)
	valueTwo = strings.ToLower(valueTwo)
	return valueOne == valueTwo
}

func SetAccessAndRefreshTokenCookies(c *gin.Context, accessToken, refreshToken string) {
	c.SetCookie(
		constants.AccessTokenCookie,
		accessToken,
		int(configs.JWT_ACCESS_TOKEN_EXPIRATION_IN_MINUTES),
		constants.HomePath,
		constants.LocalHost,
		true,
		true,
	)

	c.SetCookie(
		constants.RefreshTokenCookie,
		refreshToken,
		int(configs.JWT_REFRESH_TOKEN_EXPIRATION_IN_MINUTES),
		constants.HomePath,
		constants.LocalHost,
		true,
		true,
	)
}
