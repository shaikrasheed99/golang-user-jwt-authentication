package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetDSNString() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error in loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v", host, port, user, password, dbName)

	return dsn
}
