package models

type User struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	FirstName string
	LastName  string
	Username  string
	Password  string
	Email     string
}
