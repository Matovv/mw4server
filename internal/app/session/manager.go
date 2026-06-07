package session

import (
	"log"
	"sync"
	"time"

	"github.com/Matovv/mw4server/internal/pkg/errors"
	"github.com/Matovv/mw4server/internal/pkg/types"
)

type SessionManager struct {
    mu sync.RWMutex	
    Sessions map[types.SessionID]*Session
	TickRate uint8
}

func NewSessionManager(tickRate uint8) *SessionManager {
	if tickRate == 0 {
		tickRate = 10
	}
	return &SessionManager{
		Sessions: make(map[types.SessionID]*Session),
		TickRate: tickRate,
	}
}

func (m *SessionManager) GetSession(id types.SessionID) (*Session, error) {
    m.mu.RLock()
    defer m.mu.RUnlock()
	if _, ok := m.Sessions[id]; ok {
		return m.Sessions[id], nil
	}
    return nil, errors.ErrSessionNotFound
}

func (m *SessionManager) AddSession(s *Session) {
	m.mu.Lock()
	defer m.mu.Unlock()
	log.Println("session added:", s.ID)	
	m.Sessions[s.ID] = s
}

func (m *SessionManager) Run() {
	log.Println("session manager online!", "tick rate:", m.TickRate, "per second")	
	ticker := time.NewTicker(time.Second / time.Duration(m.TickRate))
	defer ticker.Stop()
	for range ticker.C {
		for _, session := range m.Sessions {
			session.SessionTick(m.TickRate)
		}
	}
}
