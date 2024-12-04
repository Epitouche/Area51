package database

import (
	"area51/tools"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	host := tools.GetInEnv("DB_HOST")
	user := tools.GetInEnv("POSTGRES_USER")
	password := tools.GetInEnv("POSTGRES_PASSWORD")
	dbname := tools.GetInEnv("POSTGRES_DB")
	port := tools.GetInEnv("DB_PORT")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Europe/Paris"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if os.Getenv("GIN_MODE") != "release" {
		conn = conn.Debug()
	}
	return conn
}