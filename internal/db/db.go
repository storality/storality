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

func Connect(dataDir string) *DB {
	database, err := openDB(dataDir)
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.Exec("PRAGMA journal_mode=WAL")
	if err != nil {
		log.Fatal(err)
	}

	db := &DB{
		Collections: &models.CollectionModel{DB: database},
	}
	db.Collections.CreateTable()
	return db
}

func openDB(dataDir string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dataDir + "/stor_db")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}