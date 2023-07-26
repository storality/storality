package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"storality.com/storality/internal/helpers/constants"
	"storality.com/storality/internal/helpers/exceptions"
	"storality.com/storality/internal/helpers/shout"
)

type Record struct {
	ID 					int
	Title 			string
	Slug 				string
	Content 		string
	Collection 	*Collection
	Status			constants.Status
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

type Filter struct {
	Collection Collection
}

type RecordModel struct {
	DB *sql.DB
}

func (m *RecordModel) Init() {
	var tableExists bool
	query := `SELECT EXISTS (SELECT 1 FROM sqlite_master WHERE type='table' AND name='records')`
	err := m.DB.QueryRow(query).Scan(&tableExists)
	if err != nil {
		shout.Error.Fatal(err)
	}
	if !tableExists {
		stmt := `CREATE TABLE records (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			slug TEXT NOT NULL,
			content TEXT NOT NULL,
			collection INTEGER NOT NULL,
			status TEXT NOT NULL,
			createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (collection) REFERENCES collections(id)
		);`
		_, err = m.DB.Exec(stmt)
		if err != nil {
			shout.Error.Fatal(err)
		}
	}

}

func (m *RecordModel) Insert(title string, slug string, content string, collection Collection) (int, error) {
	createdAt := time.Now()
	stmt, err := m.DB.Prepare(`INSERT INTO records (
		title,
		slug,
		content,
		collection,
		status,
		createdAt,
		updatedAt
	) VALUES (?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		title,
		slug,
		content,
		collection.ID,
		constants.StatusDraft,
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

func (m *RecordModel) Update(id int, title string, content string) error {
	updatedAt := time.Now()
	stmt, err := m.DB.Prepare("UPDATE records SET title = ?, content = ?, updatedAt = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(title, content, updatedAt, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *RecordModel) UpdateStatus(id int, status constants.Status) error {
	updatedAt := time.Now()
	stmt, err := m.DB.Prepare("UPDATE records SET status = ?, updatedAt = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(status, updatedAt, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *RecordModel) FindById(id int) (*Record, error) {
	stmt, err := m.DB.Prepare("SELECT * FROM records WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	record := &Record{
		Collection: &Collection{},
	}
	err = row.Scan(
		&record.ID,
		&record.Title,
		&record.Slug,
		&record.Content,
		&record.Collection.ID,
		&record.Status,
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
	stmt, err := m.DB.Prepare("SELECT * FROM records WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(title)
	record := &Record{
		Collection: &Collection{},
	}
	err = row.Scan(
		&record.ID,
		&record.Title,
		&record.Slug,
		&record.Content,
		&record.Collection.ID,
		&record.Status,
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
	stmt, err := m.DB.Prepare("SELECT * FROM records WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(slug)
	record := &Record{
		Collection: &Collection{},
	}
	err = row.Scan(
		&record.ID,
		&record.Title,
		&record.Slug,
		&record.Content,
		&record.Collection.ID,
		&record.Status,
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

func (m *RecordModel) FindMany(filter *Filter) ([]*Record, error) {
	stmt := "SELECT * FROM records"
	if filter.Collection != (Collection{}) {
		stmt += fmt.Sprintf(" WHERE collection = %d", filter.Collection.ID)
	}
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	records := []*Record{}
	for rows.Next() {
		record := &Record{
			Collection: &Collection{},
		}
		err = rows.Scan(
			&record.ID,
			&record.Title,
			&record.Slug,
			&record.Content,
			&record.Collection.ID,
			&record.Status,
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

func (m *RecordModel) Delete(id int) error {
	stmt, err := m.DB.Prepare("DELETE FROM records WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}