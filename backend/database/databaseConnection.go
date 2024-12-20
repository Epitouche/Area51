package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"area51/toolbox"
)

func Connection() *gorm.DB {
	host := toolbox.GetInEnv("POSTGRES_DB_HOST")
	user := toolbox.GetInEnv("POSTGRES_USER")
	password := toolbox.GetInEnv("POSTGRES_PASSWORD")
	dbname := toolbox.GetInEnv("POSTGRES_DB")
	port := toolbox.GetInEnv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Paris", host, user, password, dbname, port)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	if os.Getenv("GIN_MODE") != "release" {
		conn = conn.Debug()
	}

	return conn
}
