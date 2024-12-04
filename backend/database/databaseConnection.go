package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	host := os.Getenv("POSTGRES_DB")
	if host == "" {
		panic("POSTGRES_DB is not set")
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		panic("DB_PORT is not set")
	}
	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		panic("POSTGRES_USER is not set")
	}
	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		panic("POSTGRES_PASSWORD is not set")
	}
	dbname := os.Getenv("POSTGRES_DB")
	if dbname == "" {
		panic("POSTGRES_DB is not set")
	}

	dsn := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable TimeZone=Europe/Paris"
	dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	println("Database connection established")

	os.Getenv("BACKEND_MODE")
	if os.Getenv("BACKEND_MODE") != "release" {
		dbConnection = dbConnection.Debug()
	}
	return dbConnection
}
