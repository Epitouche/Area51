package database

import "golang.org/x/crypto/bcrypt"

type Password interface {
	HashPassword(password string) (string, error)
}

func HashPassword(password string) (string, error) {
	var passwordBytes = []byte(password)

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)
	return string(hashedPasswordBytes), err
}