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
	// DBMinConn is the minimum number of connections in the pool
	DBMinConn = 3
	//DBMaxConn is the maximum number of connections in the pool
	DBMaxConn = 100
)

func newDbConnection() (*sqlx.DB, error) {
	minConnEnv := os.Getenv("DB_MIN_CONN")
	maxConnEnv := os.Getenv("DB_MAX_CONN")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	url := makeDNS(user, password, host, port, name)
	DB, err := sqlx.ConnectContext(context.Background(), "psql", url)
	if minConnEnv != "" {
		v, err := strconv.Atoi(minConnEnv)
		if err != nil {
			log.Println("warning: DB_MIN_CONN is not a number, using default value", DBMinConn)
		}
		if v > DBMaxConn {
			log.Println("warning: DB_MIN_CONN is greater than DB_MAX_CONN, using default value", DBMinConn)
		}
		if v >= DBMinConn && v <= DBMaxConn {
			DBMinConn = v
		}

	}
	if maxConnEnv != "" {
		v, err := strconv.Atoi(maxConnEnv)
		if err != nil {
			log.Println("warning: DB_MAX_CONN is not a number, using default value", DBMaxConn)
		}
		if v < DBMinConn {
			log.Println("warning: DB_MAX_CONN is less than DB_MIN_CONN, using default value", DBMaxConn)
		}
		if v >= DBMinConn && v <= DBMaxConn {
			DBMaxConn = v
		}
	}

	DB.SetMaxIdleConns(DBMinConn)
	DB.SetMaxOpenConns(DBMaxConn)
	return DB, err
}

func makeDNS(user, password, host, port, name string) string {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user,
		password,
		host,
		port,
		name,
	)
	return connectionString
}
