package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"storality.com/storality/internal/helpers/exceptions"
	"storality.com/storality/internal/helpers/shout"
)

type Session struct {
	ID        int
	UserID    int
	Token     string
	ExpiresAt time.Time
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
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			userID INTEGER NOT NULL,
			token TEXT NOT NULL,
			expiresAt TIMESTAMP NOT NULL
		);`
		_, err = m.DB.Exec(stmt)
		if err != nil {
			shout.Error.Fatal(err)
		}
	}
	go startSessionCleanupTask(m, time.Hour)
}

func (m *SessionModel) CreateSession(userID int, token string, expiresIn time.Duration) (int, error) {
	expiresAt := time.Now().Add(expiresIn)
	stmt, err := m.DB.Prepare(`INSERT INTO sessions (
		userID,
		token,
		expiresAt
	) VALUES (?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		userID,
		token,
		expiresAt,
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

func (m *SessionModel) GetSessionByID(sessionID int) (*Session, error) {
	stmt, err := m.DB.Prepare("SELECT * FROM sessions WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(sessionID)
	session := &Session{}
	err = row.Scan(
		&session.ID,
		&session.UserID,
		&session.Token,
		&session.ExpiresAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, exceptions.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return session, nil
}

func (m *SessionModel) DeleteSession(sessionID int) error {
	stmt, err := m.DB.Prepare("DELETE FROM sessions WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(sessionID)
	if err != nil {
		return err
	}
	return nil
}

func (m *SessionModel) DeleteExpiredSessions() error {
    _, err := m.DB.Exec("DELETE FROM sessions WHERE expiresAt < ?", time.Now())
    if err != nil {
        return fmt.Errorf("delete expired sessions error: %w", err)
    }
    return nil
}

func startSessionCleanupTask(m *SessionModel, interval time.Duration) {
    ticker := time.NewTicker(interval)
    defer ticker.Stop()

    for range ticker.C {
        if err := m.DeleteExpiredSessions(); err != nil {
            shout.Error.Printf("Error cleaning up expired sessions: %v", err)
        }
    }
}



