package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

// CreateConnection opens the connection to the postgres db
func CreateConnection() (*gorm.DB, error) {

	// return connection
	return gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s "+
				"password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		),
	)
}
