package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func connect() (*gorm.DB, error) {
	dbType := os.Getenv("DB_TYPE")
	dsn := os.Getenv("DB_DSN")

	var openedDb gorm.Dialector
	if dbType == "postgres" {
		openedDb = postgres.Open(dsn)
	} else {
		openedDb = sqlite.Open("store.db")
	}

	db, err := gorm.Open(openedDb, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
