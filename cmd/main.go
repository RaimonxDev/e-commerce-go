package main

import (
	"github.com/RaimonxDev/e-commerce-go.git/infrastructure/handler/response"
	"log"
	"os"
)

func main() {
	// Load .env file
	err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}
	// Validate environment variables
	err = validateEnvironments()
	if err != nil {
		log.Fatal(err)
	}
	// Init echo
	e := newHTTP(response.HTTPErrorHandler)

	// Init Db connection
	db, err := newDbConnection()
	if err != nil {
		log.Fatal(err)
	}
	_ = db
	err = e.Start(":" + os.Getenv("SERVER_PORT"))

	if err != nil {
		log.Fatal(err)
	}
}
