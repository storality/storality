package models

import (
	"database/sql"
	"time"
)

type Record struct {
	ID 					int
	Title 			string
	Slug 				string
	Content 		string
	Collection 	Collection
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

type RecordModel struct {
	DB *sql.DB
}

func (m *RecordModel) CreateTable() {

}