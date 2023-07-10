package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashedPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error while generating hashed password")
		return "", err
	}

	return string(hash), nil
}

func CheckPassword(hashedPass string, inputPass string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(inputPass)); err != nil {
		fmt.Println("Error while comparing password")
		return false
	}

	return true
}
