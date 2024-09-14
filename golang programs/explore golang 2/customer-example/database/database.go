package database

import (
	"fmt"
	"log"
	"os"
	"postgres-crud/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// connection with postgres database
func InitDB() (*gorm.DB, error) {
	err := godotenv.Load()

	if err != nil {
		log.Println("ERROR LOADING ENVIRONMENT VARIABLES")
	}

	DB_USER := os.Getenv("DB_USER")
	DB_USER_PASSWORD := os.Getenv("DB_USER_PASSWORD")
	HOST := os.Getenv("HOST")
	DATABASE := os.Getenv("DATABASE")

	fmt.Println(DB_USER, DB_USER_PASSWORD, HOST, DATABASE)

	URL := fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable", DB_USER, DB_USER_PASSWORD, HOST, DATABASE)

	connection, err := gorm.Open(postgres.Open(URL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqldb, err := connection.DB()
	if err != nil {
		return nil, err
	}

	err = sqldb.Ping()
	if err != nil {
		return nil, err
	}

	connection.AutoMigrate(&model.Customer{})

	log.Println("Successfully connected to database!")
	return connection, nil
}
