package database

import "golang.org/x/crypto/bcrypt"

type Password interface {
	HashPassword(password string) (string, error)
	CompareHashAndPassword(hashedPassword, password string) bool
}

func HashPassword(password string) (string, error) {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	return string(passwordHash), err
}

func CompareHashAndPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
