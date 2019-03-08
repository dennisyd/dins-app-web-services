package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

// CreateConnection opens the connection to the postgres db
func CreateConnection() (*gorm.DB, error) {

	// Get connection details from environmental variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	// return connection
	return gorm.Open(
		"postgres",
		fmt.Sprintf(
			"postgres://%s:%s@%s/%s?sslmode=disable",
			user,
			password,
			host,
			DBName,
		),
	)
}
