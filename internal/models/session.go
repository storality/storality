package models

import (
	"database/sql"
	"time"

	"storality.com/storality/internal/helpers/shout"
)

type Session struct {
	Token		string
	Data 		[]byte
	Expires	time.Time
}

type SessionModel struct {
	DB *sql.DB
}

func (m *SessionModel) Init() {
	var tableExists bool
	query := `SELECT EXISTS (SELECT 1 FROM sqlite_master WHERE type='table' AND name='sessions')`
	err := m.DB.QueryRow(query).Scan(&tableExists)
	if err != nil {
		shout.Error.Fatal(err)
	}
	if !tableExists {
		stmt := `CREATE TABLE sessions (
			token TEXT PRIMARY KEY NOT NULL,
			data BLOB,
			expires TIMESTAMP
		);`
		_, err = m.DB.Exec(stmt)
		if err != nil {
			shout.Error.Fatal(err)
		}
	}
	go m.flush()
}

func (m *SessionModel) Create(s *Session) error {
	stmt, err := m.DB.Prepare("INSERT INTO sessions(token, data, expires) values(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(s.Token, s.Data, s.Expires)
	return err
}

func (m *SessionModel) Get(token string) (*Session, error) {
	s := &Session{}
	err := m.DB.QueryRow("SELECT token, data, expires FROM sessions WHERE token = ?", token).Scan(&s.Token, &s.Data, &s.Expires)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (m *SessionModel) Delete(token string) error {
	_, err := m.DB.Exec("DELETE FROM sessions WHERE token = ?", token)
	return err
}

func (m *SessionModel) flush() {
	ticker := time.NewTicker(24 * time.Hour)
	for range ticker.C {
		_, err := m.DB.Exec("DELETE FROM sessions WHERE expires < datetime('now')")
		if err != nil {
			shout.Error.Println(err)
		}
	}
}