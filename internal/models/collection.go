package models

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"storality.com/storality/internal/helpers/exceptions"
	"storality.com/storality/internal/helpers/shout"
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

func (m *CollectionModel) Init() error {
	var tableExists bool
	query := `SELECT EXISTS (SELECT 1 FROM sqlite_master WHERE type='table' AND name='collections')`
	err := m.DB.QueryRow(query).Scan(&tableExists)
	if err != nil {
		shout.Error.Fatal(err)
	}
	if !tableExists {
		stmt := `CREATE TABLE collections (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			plural TEXT NOT NULL,
			createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`
		_, err = m.DB.Exec(stmt)
		if err != nil {
			shout.Error.Fatal(err)
		}

		_, err := m.Insert("post", "posts")
		if err != nil {
			shout.Error.Fatal(err)
		}
		_, err = m.Insert("page", "pages")
		if err != nil {
			shout.Error.Fatal()
		}
	}
	return nil
}

func (m *CollectionModel) Insert(name string, plural string) (int, error) {
	createdAt := time.Now()
	stmt, err := m.DB.Prepare("INSERT INTO collections (name, plural, createdAt, updatedAt) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(name, plural, createdAt, createdAt)
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
	stmt, err := m.DB.Prepare("SELECT * FROM collections WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	collection := &Collection{}
	err = row.Scan(
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
	stmt, err := m.DB.Prepare("SELECT * FROM collections WHERE name = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(name)
	collection := &Collection{}
	err = row.Scan(
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

func (m *CollectionModel) FindByPlural(plural string) (*Collection, error) {
	trim := strings.Trim(plural, "/")
	stmt, err := m.DB.Prepare("SELECT * FROM collections WHERE plural = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(trim)
	collection := &Collection{}
	err = row.Scan(
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

func (m *CollectionModel) FindMany() ([]*Collection, error) {
	stmt := "SELECT * FROM collections"
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