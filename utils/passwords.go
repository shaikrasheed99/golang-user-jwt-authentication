package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashedPassword(password string) (string, error) {
	fmt.Println("[GenerateHashedPassword] Generating hash of the password")

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("[GenerateHashedPassword] Error while generating hashed password")
		return "", err
	}

	fmt.Println("[GenerateHashedPassword] Generated hash of the password")
	return string(hash), nil
}

func CheckPassword(hashedPass string, inputPass string) bool {
	fmt.Println("[CheckPassword] Checking hashed password and input password")

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(inputPass)); err != nil {
		fmt.Println("[CheckPassword] Passwords are not equal")
		return false
	}

	fmt.Println("[CheckPassword] Password is correct")
	return true
}
