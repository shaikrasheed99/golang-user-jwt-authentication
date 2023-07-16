package helpers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashedPassword(password string) (string, error) {
	fmt.Println("[GenerateHashedPasswordHelper] Generating hash of the password")

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("[GenerateHashedPasswordHelper] Error while generating hashed password")
		return "", err
	}

	fmt.Println("[GenerateHashedPasswordHelper] Generated hash of the password")
	return string(hash), nil
}

func CheckPassword(hashedPass string, inputPass string) bool {
	fmt.Println("[CheckPasswordHelper] Checking hashed password and input password")

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(inputPass)); err != nil {
		fmt.Println("[CheckPasswordHelper] Passwords are not equal")
		return false
	}

	fmt.Println("[CheckPasswordHelper] Password is correct")
	return true
}
