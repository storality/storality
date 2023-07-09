package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"storality.com/storality/internal/models"
)

type DB struct {
	Collections *models.CollectionModel
}

func Connect(driver string, connection string) *DB {
	database, err := openDB(driver, connection)
	if err != nil {
		log.Fatal(err)
	}

	if driver == "sqlite3" {
		_, err := database.Exec("PRAGMA journal_mode=WAL")
		if err != nil {
			log.Fatal(err)
		}
	}

	db := &DB{
		Collections: &models.CollectionModel{DB: database, Driver: driver},
	}
	db.Collections.CreateTable()
	return db
}

func openDB(driver string, connection string) (*sql.DB, error) {
	db, err := sql.Open(driver, connection)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}