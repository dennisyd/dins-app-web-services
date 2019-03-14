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
			os.Getenv("IDENTITY_DB_HOST"),
			os.Getenv("IDENTITY_DB_PORT"),
			os.Getenv("IDENTITY_DB_USER"),
			os.Getenv("IDENTITY_DB_PASSWORD"),
			os.Getenv("IDENTITY_DB_NAME"),
		),
	)
}
