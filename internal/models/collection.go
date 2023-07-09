package models

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

type Collection struct {
	ID 				int
	Name 		string
	Plural 		string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CollectionModel struct {
	DB *sql.DB
	Driver string
}

func (m *CollectionModel) CreateTable() {
	stmt := `CREATE TABLE IF NOT EXISTS collections (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name TEXT NOT_NULL,
		plural TEXT NOT_NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	result, err := m.DB.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
	if rowsAffected, _ := result.RowsAffected(); rowsAffected > 0 {
		m.Insert("post", "posts")
		m.Insert("page", "pages")
	}
}

func (m *CollectionModel) Insert(name string, plural string) (int, error) {
	stmt := `INSERT INTO collections (name, plural)
	VALUES (?, ?, UTC_TIMESTAMP()))`
	result, err := m.DB.Exec(stmt, name, plural)
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
	stmt := `SELECT id, name, plural, createdAt FROM collections WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)
	collection := &Collection{}
	err := row.Scan(&collection.ID, &collection.Name, &collection.Plural, &collection.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	return collection, nil
}

func (m *CollectionModel) FindByName(name string) (*Collection, error) {
	stmt := `SELECT id, name, plural, createdAt FROM collections WHERE name = ?`
	row := m.DB.QueryRow(stmt, name)
	collection := &Collection{}
	err := row.Scan(&collection.ID, &collection.Name, &collection.Plural, &collection.CreatedAt)
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
	stmt := `SELECT id, name, plural, createdAt FROM collections`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	collections := []*Collection{}
	for rows.Next() {
		collection := &Collection{}
		err = rows.Scan(&collection.ID, &collection.Name, &collection.Plural, &collection.CreatedAt)
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