package websocket

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Transport :
type Transport struct {
	port         int
	path         string
	redirect     bool
	handler      Handler
	errorHandler ErrorHandler
	authHandler  AuthHandler
	upgrader     websocket.Upgrader
	sessions     map[string]*Session
}

// Handler :
type Handler func(messageType int, message []byte, fromSession *Session)

// AuthHandler :
type AuthHandler func(http.ResponseWriter, *http.Request) (string, *Session, error)

// ErrorHandler :
type ErrorHandler func(http.ResponseWriter, *http.Request, error)

// NewWebsocket :
func NewWebsocket() *Transport {
	return &Transport{
		port:     3000,
		upgrader: websocket.Upgrader{},
		sessions: make(map[string]*Session),
		authHandler: func(res http.ResponseWriter, req *http.Request) (string, *Session, error) {
			return uuid.New().String(), nil, nil
		},
		errorHandler: func(res http.ResponseWriter, req *http.Request, err error) {
			data := []byte("error: " + err.Error())
			res.Write(data)
		},
	}
}

// UsePort :
func (t *Transport) UsePort(port int) *Transport {
	t.port = port
	return t
}

// UseUpgrader :
func (t *Transport) UseUpgrader(upgrader websocket.Upgrader) *Transport {
	t.upgrader = upgrader
	return t
}

// UseRedirect :
func (t *Transport) UseRedirect(redirect bool) *Transport {
	t.redirect = redirect
	return t
}

// UseErrorHandler :
func (t *Transport) UseErrorHandler(handler ErrorHandler) *Transport {
	t.errorHandler = handler
	return t
}

// UseHandler :
func (t *Transport) UseHandler(handler Handler) *Transport {
	t.handler = handler
	return t
}

// UseAuthHandler :
func (t *Transport) UseAuthHandler(handler AuthHandler) *Transport {
	t.authHandler = handler
	return t
}

// UpdateSession :
func (t *Transport) UpdateSession(id string, session *Session) {
	t.sessions[id] = session
}

// GetSession :
func (t *Transport) GetSession(id string) *Session {
	if session, ok := t.sessions[id]; ok {
		return session
	}
	return nil
}
