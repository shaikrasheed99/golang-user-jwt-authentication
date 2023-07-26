package configs

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/shaikrasheed99/golang-user-jwt-authentication/constants"
)

var (
	DB_NAME                                 string
	DB_PORT                                 string
	DB_HOST                                 string
	DB_USER                                 string
	DB_PASSWORD                             string
	JWT_SECRET                              string
	JWT_ISSUER                              string
	JWT_ACCESS_TOKEN_EXPIRATION_IN_MINUTES  uint
	JWT_REFRESH_TOKEN_EXPIRATION_IN_MINUTES uint
)

func LoadConfigs() error {
	fmt.Println("[LoadConfigs] Loading env variables")

	err := godotenv.Load()
	if err != nil {
		fmt.Println("[GetDSNString]", err.Error())
		return err
	}

	DB_NAME = os.Getenv("DB_NAME")
	DB_PORT = os.Getenv("DB_PORT")
	DB_HOST = os.Getenv("DB_HOST")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	JWT_SECRET = os.Getenv("JWT_SECRET")
	JWT_ISSUER = os.Getenv("JWT_ISSUER")

	jat, err := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_EXPIRATION_IN_MINUTES"))
	if err != nil {
		errMessage := constants.ErrInvalidTokenExpiration
		fmt.Println("[LoadConfigs]", errMessage)
		return errors.New(errMessage)
	}
	JWT_ACCESS_TOKEN_EXPIRATION_IN_MINUTES = uint(jat)

	jrt, err := strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_EXPIRATION_IN_MINUTES"))
	if err != nil {
		errMessage := constants.ErrInvalidTokenExpiration
		fmt.Println("[LoadConfigs]", errMessage)
		return errors.New(errMessage)
	}
	JWT_REFRESH_TOKEN_EXPIRATION_IN_MINUTES = uint(jrt)

	fmt.Println("[LoadConfigs] Env variables have loaded")
	return nil
}

func GetDSNString() string {
	fmt.Println("[GetDSNString] Getting dns string")

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	fmt.Println("[GetDSNString] Returning dns string value")
	return dsn
}
