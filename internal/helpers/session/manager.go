package session

import (
	"crypto/rand"
	"fmt"
	"sync"
	"time"
)

type SessionManager struct {
	sessions 	map[string]time.Time
	mutex			sync.Mutex
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]time.Time),
	}
}

func (manager *SessionManager) GenerateSessionID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func (manager *SessionManager) CreateSession() string {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()

	sessionID := manager.GenerateSessionID()
	manager.sessions[sessionID] = time.Now()
	return sessionID
}

func (manager *SessionManager) CheckSession(sessionID string) bool {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()

	timestamp, exists := manager.sessions[sessionID]
	if !exists {
		return false
	}

	return time.Since(timestamp) <= 12 * time.Hour
}