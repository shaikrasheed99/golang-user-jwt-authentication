package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetDSNString() string {
	fmt.Println("[GetDSNString] Getting dns string")
	fmt.Println("[GetDSNString] Loading env variables")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("[GetDSNString]", err.Error())
		return ""
	}
	fmt.Println("[GetDSNString] Env variables have loaded")

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v", host, port, user, password, dbName)

	fmt.Println("[GetDSNString] Returning dns string value")
	return dsn
}
