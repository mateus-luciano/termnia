package core

type Session struct {
	ID       int
	Terminal Terminal
	Title    string
}

type SessionManager struct {
	sessions map[int]*Session
	nextID   int
}

func NewSessionManager() *SessionManager {
	return &SessionManager{sessions: make(map[int]*Session)}
}

func (sm *SessionManager) NewSession(title string, t Terminal) *Session {
	sm.nextID++

	s := &Session{
		ID:       sm.nextID,
		Terminal: t,
		Title:    title,
	}
	sm.sessions[sm.nextID] = s

	return s
}

func (sm *SessionManager) CloseSession(id int) {
	if s, ok := sm.sessions[id]; ok {
		s.Terminal.Kill()
		delete(sm.sessions, id)
	}
}
