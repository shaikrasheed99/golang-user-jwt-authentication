package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	fmt.Println("[InitDatabase] Initiating database")
	dsn := GetDSNString()

	fmt.Println("[InitDatabase] Opening the database connection")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("[InitDatabase]", err.Error())
		return nil
	}

	fmt.Println("[InitDatabase] Database connect has established")
	return db
}
