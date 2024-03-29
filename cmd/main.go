package main

import (
	"log"
	"os"

	"github.com/RaimonxDev/e-commerce-go.git/infrastructure/handler/response"
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
	// init migration
	err = runMigrationDB(db)
	if err != nil {
		log.Fatalf("error while running migration: %v", err)
	}

	// Run Server
	err = e.Start(":" + os.Getenv("SERVER_PORT"))

	if err != nil {
		log.Fatal(err)
	}
}
