package models

import (
	"database/sql"
	"errors"
	"time"

	"storality.com/storality/internal/helpers/exceptions"
)

type Record struct {
	ID 					int
	Title 			string
	Slug 				string
	Content 		string
	Collection 	*Collection
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

type RecordModel struct {
	DB *sql.DB
}

func (m *RecordModel) VerifyTable() error {
	var tableExists bool
	query := `SELECT EXISTS (SELECT 1 FROM sqlite_master WHERE type='table' AND name='records')`
	err := m.DB.QueryRow(query).Scan(&tableExists)
	if err != nil {
		return err
	}
	if !tableExists {
		stmt := `CREATE TABLE records (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL UNIQUE,
			slug TEXT NOT NULL UNIQUE,
			content TEXT NOT NULL UNIQUE,
			createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			collection INTEGER NOT NULL,
			FOREIGN_KEY (collection) REFERENCES collections(id)
		)`
		_, err = m.DB.Exec(stmt)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *RecordModel) Insert(title string, slug string, content string, collection Collection) (int, error) {
	createdAt := time.Now()
	stmt := `INSERT INTO collections (
		title,
		slug,
		content,
		collection,
		createdAt,
		updatedAt
	) VALUES (?, ?, ?, ?, ?, ?)`
	result, err := m.DB.Exec(
		stmt,
		title,
		slug,
		content,
		collection.ID,
		createdAt,
		createdAt,
	)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *RecordModel) FindById(id int) (*Record, error) {
	stmt := `SELECT id, title, slug, content, collection, createdAt, updatedAt FROM collections WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)
	record := &Record{}
	err := row.Scan(
		&record.ID,
		&record.Title,
		&record.Slug,
		&record.Content,
		&record.Collection.ID,
		&record.CreatedAt,
		&record.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, exceptions.ErrNoRecord
		} else {
			return nil, err
		}
	}

	collectionModel := &CollectionModel{DB: m.DB}
	collection, err := collectionModel.FindById(record.Collection.ID)
	if err != nil {
			return nil, err
	}
	record.Collection = collection

	return record, nil
}

func (m *RecordModel) FindByTitle(title string) (*Record, error) {
	stmt := `SELECT id, title, slug, content, collection, createdAt, updatedAt FROM collections WHERE title = ?`
	row := m.DB.QueryRow(stmt, title)
	record := &Record{}
	err := row.Scan(
		&record.ID,
		&record.Title,
		&record.Slug,
		&record.Content,
		&record.Collection.ID,
		&record.CreatedAt,
		&record.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, exceptions.ErrNoRecord
		} else {
			return nil, err
		}
	}

	collectionModel := &CollectionModel{DB: m.DB}
	collection, err := collectionModel.FindById(record.Collection.ID)
	if err != nil {
			return nil, err
	}
	record.Collection = collection

	return record, nil
}

func (m *RecordModel) FindBySlug(slug string) (*Record, error) {
	stmt := `SELECT id, title, slug, content, collection, createdAt, updatedAt FROM collections WHERE slug = ?`
	row := m.DB.QueryRow(stmt, slug)
	record := &Record{}
	err := row.Scan(
		&record.ID,
		&record.Title,
		&record.Slug,
		&record.Content,
		&record.Collection.ID,
		&record.CreatedAt,
		&record.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, exceptions.ErrNoRecord
		} else {
			return nil, err
		}
	}

	collectionModel := &CollectionModel{DB: m.DB}
	collection, err := collectionModel.FindById(record.Collection.ID)
	if err != nil {
			return nil, err
	}
	record.Collection = collection

	return record, nil
}

func (m *RecordModel) FindAll() ([]*Record, error) {
	stmt := `SELECT id, title, slug, content, collection, createdAt, updatedAt FROM records`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	records := []*Record{}
	for rows.Next() {
		record := &Record{}
		err = rows.Scan(
			&record.ID,
			&record.Title,
			&record.Slug,
			&record.Content,
			&record.Collection.ID,
			&record.CreatedAt,
			&record.UpdatedAt,
		)
		if err != nil {
			return nil , err
		}

		collectionModel := &CollectionModel{DB: m.DB}
		collection, err := collectionModel.FindById(record.Collection.ID)
		if err != nil {
				return nil, err
		}
		record.Collection = collection

		records = append(records, record)
	}
	if rows.Err(); err != nil {
		return nil, err
	}
	return records, nil
}