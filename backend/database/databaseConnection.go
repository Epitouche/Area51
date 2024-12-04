package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getInEnv(varWanted string) (envVar string) {
	envVar = os.Getenv(varWanted)
	if envVar == "" {
		panic(varWanted + " is not set")
	}
	return envVar
}

func Connection() *gorm.DB {
	host := getInEnv("POSTGRES_DB")
	user := getInEnv("POSTGRES_USER")
	password := getInEnv("POSTGRES_PASSWORD")
	dbname := getInEnv("POSTGRES_DB")
	port := getInEnv("DB_PORT")

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
