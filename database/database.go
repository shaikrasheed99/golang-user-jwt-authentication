package database

import (
	"github.com/shaikrasheed99/golang-user-jwt-authentication/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dsn := GetDSNString()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	helper.LogError(err)

	return db
}
