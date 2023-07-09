package models

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

type Collection struct {
	ID 				int
	Title 		string
	Plural 		string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CollectionModel struct {
	DB *sql.DB
	Driver string
}

func (m *CollectionModel) CreateTable() {
	stmt := `CREATE TABLE collections (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title TEXT NOT_NULL,
		plural TEXT NOT_NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := m.DB.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *CollectionModel) Insert(title string, plural string) (int, error) {
	stmt := `INSERT INTO collections (title, plural)
	VALUES (?, ?, UTC_TIMESTAMP()))`
	result, err := m.DB.Exec(stmt, title, plural)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *CollectionModel) FindById(id int) (*Collection, error) {
	stmt := `SELECT id, title, plural, createdAt FROM collections WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)
	collection := &Collection{}
	err := row.Scan(&collection.ID, &collection.Title, &collection.Plural, &collection.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	return collection, nil
}

func (m *CollectionModel) FindAll() ([]*Collection, error) {
	stmt := `SELECT id, title, plural, createdAt FROM collections`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	collections := []*Collection{}
	for rows.Next() {
		collection := &Collection{}
		err = rows.Scan(&collection.ID, &collection.Title, &collection.Plural, &collection.CreatedAt)
		if err != nil {
			return nil , err
		}
		collections = append(collections, collection)
	}
	if rows.Err(); err != nil {
		return nil, err
	}
	return collections, nil
}