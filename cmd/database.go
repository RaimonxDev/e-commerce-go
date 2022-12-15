package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
)

var (
	// DefaultDbMinConnexion is the minimum number of connections in the pool
	DefaultDbMinConnexion = 3
	//DefaultDbMaxConnexion is the maximum number of connections in the pool
	DefaultDbMaxConnexion = 100
)

func newDbConnection() (*sqlx.DB, error) {
	minConnEnv := os.Getenv("DB_MIN_CONNEXION")
	maxConnEnv := os.Getenv("DB_MAX_CONNEXION")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	url := makeDNS(user, password, host, port, name)
	DB, err := sqlx.ConnectContext(context.Background(), "postgres", url)

	if minConnEnv != "" {
		v, err := strconv.Atoi(minConnEnv)

		if err != nil {
			DB.SetMaxIdleConns(DefaultDbMinConnexion)
			log.Println("warning: DB_MIN_CONN is not a number, using default value", DefaultDbMinConnexion)
		}
		if v > DefaultDbMaxConnexion {
			DB.SetMaxIdleConns(DefaultDbMinConnexion)
			log.Println("warning: DB_MIN_CONN is greater than DB_MAX_CONN, using default value", DefaultDbMinConnexion)
		}
		if v >= DefaultDbMinConnexion && v <= DefaultDbMaxConnexion {
			DB.SetMaxIdleConns(v)
		}

	}
	if maxConnEnv != "" {
		v, err := strconv.Atoi(maxConnEnv)
		if err != nil {
			DB.SetMaxOpenConns(DefaultDbMaxConnexion)
			log.Println("warning: DB_MAX_CONN is not a number, using default value", DefaultDbMaxConnexion)
		}
		if v < DefaultDbMinConnexion {
			DB.SetMaxOpenConns(DefaultDbMaxConnexion)
			log.Println("warning: DB_MAX_CONN is less than DB_MIN_CONN, using default value", DefaultDbMaxConnexion)
		}
		if v >= DefaultDbMinConnexion && v <= DefaultDbMaxConnexion {
			DB.SetMaxOpenConns(v)
		}
	}
	return DB, err
}

func makeDNS(user, password, host, port, name string) string {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		name,
	)
	return connectionString
}
