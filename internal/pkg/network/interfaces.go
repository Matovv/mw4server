package network

import (	
	"github.com/Matovv/mw4server/internal/app/session"
	"github.com/Matovv/mw4server/internal/pkg/types"
)

type sessionManager interface {
	GetSession(id types.SessionID) (*session.Session, error)
}
