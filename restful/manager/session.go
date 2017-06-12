package manager

import (
	"github.com/songtianyi/rrframework/utils"
	"github.com/songtianyi/wechat-go/wxweb"
	"sync"
)

var (
	GlobalSessionManager = &SessionManager{
		sm: make(map[string]*wxweb.Session),
	}
)

type SessionManager struct {
	sm   map[string]*wxweb.Session
	lock sync.RWMutex
}

func (s *SessionManager) Add(session *wxweb.Session) string {
	// generate uuid
	uuid := rrutils.NewV4().String()

	s.lock.Lock()
	s.sm[uuid] = session
	s.lock.Unlock()
	return uuid
}

func (s *SessionManager) Get(uuid string) *wxweb.Session {
	s.lock.RLock()
	session := s.sm[uuid]
	s.lock.RUnlock()
	return session
}

func (s *SessionManager) Set(uuid string, session *wxweb.Session) {
	s.lock.Lock()
	s.sm[uuid] = session
	s.lock.Unlock()
	return
}
