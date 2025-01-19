package schemas

import (
	"errors"

	"gorm.io/gorm"
)

type Database struct {
	Connection *gorm.DB
}

var (
	ErrorHashingPassword = errors.New("error while hashing password")
)
