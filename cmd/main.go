package main

import (
	"github.com/RaimonxDev/e-commerce-go.git/infrastucture/handler/response"
	"log"
	"os"
)

func main() {

	err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}
	err = validateEnvironments()
	if err != nil {
		log.Fatal(err)
	}
	e := newHTTP(response.HTTPErrorHandler)

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
