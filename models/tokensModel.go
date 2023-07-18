package models

import "time"

type Tokens struct {
	ID           uint   `gorm:"autoIncrement"`
	Username     string `gorm:"primaryKey"`
	AccessToken  string
	RefreshToken string
	CreatedAt    time.Time
}
