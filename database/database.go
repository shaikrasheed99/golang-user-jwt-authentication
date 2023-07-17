package database

import (
	"fmt"

	"github.com/shaikrasheed99/golang-user-jwt-authentication/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	fmt.Println("[InitDatabase] Initiating database")
	dsn := configs.GetDSNString()

	fmt.Println("[InitDatabase] Opening the database connection")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("[InitDatabase]", err.Error())
		return nil
	}

	fmt.Println("[InitDatabase] Database connection has established")
	return db
}
