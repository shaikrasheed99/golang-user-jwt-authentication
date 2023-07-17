package models

type User struct {
	ID        uint `gorm:"autoIncrement"`
	FirstName string
	LastName  string
	Username  string `gorm:"primaryKey"`
	Password  string
	Role      string
	Email     string
}
