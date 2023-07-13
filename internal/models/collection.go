package models

import (
	"database/sql"
	"errors"
	"log"
	"strings"
	"time"

	"storality.com/storality/internal/helpers/exceptions"
)

type Collection struct {
	ID 				int
	Name 			string
	Plural 		string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CollectionModel struct {
	DB *sql.DB
}

func (m *CollectionModel) CreateTable() {

	var err error
	err = m.DB.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='collections'").Scan()
	if err != nil {
		if err == sql.ErrNoRows {
			stmt := `CREATE TABLE collections (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				name TEXT NOT NULL UNIQUE,
				plural TEXT NOT NULL UNIQUE,
				createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);`
			_, err = m.DB.Exec(stmt)
			if err != nil {
				log.Fatal(err)
			}
			_, err := m.Insert("post", "posts")
			if err != nil {
				log.Fatal(err)
			}
			_, err = m.Insert("page", "pages")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func (m *CollectionModel) Insert(name string, plural string) (int, error) {
	createdAt := time.Now()
	stmt := "INSERT INTO collections (name, plural, createdAt, updatedAt) VALUES (?, ?, ?, ?)"
	result, err := m.DB.Exec(stmt, name, plural, createdAt, createdAt)
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
	stmt := `SELECT id, name, plural, createdAt, updatedAt FROM collections WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)
	collection := &Collection{}
	err := row.Scan(
		&collection.ID,
		&collection.Name,
		&collection.Plural,
		&collection.CreatedAt,
		&collection.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, exceptions.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return collection, nil
}

func (m *CollectionModel) FindByName(name string) (*Collection, error) {
	stmt := `SELECT id, name, plural, createdAt, updatedAt FROM collections WHERE name = ?`
	row := m.DB.QueryRow(stmt, name)
	collection := &Collection{}
	err := row.Scan(
		&collection.ID,
		&collection.Name,
		&collection.Plural,
		&collection.CreatedAt,
		&collection.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, exceptions.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return collection, nil
}

func (m *CollectionModel) FindBySlug(slug string) (*Collection, error) {
	plural := strings.Trim(slug, "/")
	stmt := `SELECT id, name, plural, createdAt, updatedAt FROM collections WHERE plural = ?`
	row := m.DB.QueryRow(stmt, plural)
	collection := &Collection{}
	err := row.Scan(
		&collection.ID,
		&collection.Name,
		&collection.Plural,
		&collection.CreatedAt,
		&collection.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, exceptions.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return collection, nil
}

func (m *CollectionModel) FindAll() ([]*Collection, error) {
	stmt := `SELECT id, name, plural, createdAt, updatedAt FROM collections`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	collections := []*Collection{}
	for rows.Next() {
		collection := &Collection{}
		err = rows.Scan(
			&collection.ID,
			&collection.Name,
			&collection.Plural,
			&collection.CreatedAt,
			&collection.UpdatedAt,
		)
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