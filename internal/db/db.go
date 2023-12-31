package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"storality.com/storality/internal/helpers/shout"
	"storality.com/storality/internal/models"
)

type DB struct {
	Collections *models.CollectionModel
	Records *models.RecordModel
	Sessions *models.SessionModel
}

func Connect() *DB {
	database, err := openDB()
	if err != nil {
		shout.Error.Fatal(err)
	}

	_, err = database.Exec("PRAGMA journal_mode=WAL")
	if err != nil {
		shout.Error.Fatal(err)
	}

	db := &DB{
		Collections: &models.CollectionModel{DB: database},
		Records: &models.RecordModel{DB: database},
		Sessions: &models.SessionModel{DB: database},
	}

	db.Collections.Init()
	db.Records.Init()
	db.Sessions.Init()
	return db
}

func openDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "stor_data/stor_db")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}