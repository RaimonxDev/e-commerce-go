package main

import (
	"fmt"
	"os"
	"strings"

	// WithInstance
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	// Empty
	// "github.com/golang-migrate/migrate/v4"
	// _ "github.com/golang-migrate/migrate/v4/database/postgres"
	// _ "github.com/golang-migrate/migrate/v4/source/file"
)

func runMigrationDB(db *sqlx.DB) error {
	filePath := os.Getenv("MIGRATIONS_PATH")
	if filePath == "" {
		return fmt.Errorf("MIGRATIONS_PATH is not set")
	}
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("error while creating migration driver: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://"+filePath, "ecommerce", driver)
	if err != nil {
		return fmt.Errorf("error while creating migration instance: %w", err)
	}

	action := strings.ToLower(os.Getenv("MIGRATION_ACTION"))

	if action != "up" && action != "down" {
		return fmt.Errorf("invalid action %s", action)
	}

	if action == "" {
		action = "up"
	}
	// fixme Not working
	if action == "down" {
		m.Down()
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error while running migration: %w", err)
	}
	return nil

	// EMPTY
	// m, err := migrate.New(
	// 	"file://"+filePath,
	// 	"postgres://ramon:1234@localhost:5434/ecommerce?sslmode=disable")

	// if err != nil {
	// 	return fmt.Errorf("error while creating migration instance: %w", err)
	// }

	// return m.Up()
}
