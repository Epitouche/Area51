package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"area51/toolbox"
)

// func getInEnv(varWanted string) (envVar string) {
// 	envVar = os.Getenv(varWanted)
// 	if envVar == "" {
// 		panic(varWanted + " is not set")
// 	}
// 	return envVar
// }

func Connection() *gorm.DB {
	host := toolbox.GetInEnv("POSTGRES_DB_HOST")
	user := toolbox.GetInEnv("POSTGRES_USER")
	password := toolbox.GetInEnv("POSTGRES_PASSWORD")
	dbname := toolbox.GetInEnv("POSTGRES_DB")
	port := toolbox.GetInEnv("DB_PORT")

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
