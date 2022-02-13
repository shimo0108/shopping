package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func NewDB() *gorm.DB {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load(".env")

		if err != nil {
			fmt.Printf("読み込み出来ませんでした: %v", err)
		}
	}

	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")

	var connection string = "host=" + DB_HOST + "\n" + "user=" + DB_USER + "\n" + "password=" + DB_PASS + "\n" + "dbname=" + DB_NAME + "\n" + "port=5432" + "\n" + "sslmode=disable" + "\n"

	fmt.Println(connection)
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
