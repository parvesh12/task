package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {

	er := godotenv.Load()

	if er != nil {
		log.Fatalf("Error loading .env file")
	}

	var (
		db  *gorm.DB
		err error
	)

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD")) //Build connection string
	
	db, err = gorm.Open(postgres.Open(dbUri), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	return db
}
