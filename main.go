package websocket

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Transport :
type Transport struct {
	port            int
	path            string
	redirect        bool
	handler         func(messageType int, message []byte, session *Session)
	notFoundHandler func(http.ResponseWriter, *http.Request)
	errorHandler    func(http.ResponseWriter, *http.Request, error)
	authHandler     func(http.ResponseWriter, *http.Request) (string, *Session, error)
	upgrader        websocket.Upgrader
	sessions        map[string]*Session
}

// NewWebsocket :
func NewWebsocket() *Transport {
	return &Transport{
		port:            3000,
		path:            "/ws",
		redirect:        true,
		notFoundHandler: nil,
		handler:         nil,
		upgrader:        websocket.Upgrader{},
		sessions:        make(map[string]*Session),
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

// UseNotFoundHandler :
func (t *Transport) UseNotFoundHandler(notFoundHandler func(http.ResponseWriter, *http.Request)) *Transport {
	t.notFoundHandler = notFoundHandler
	return t
}

// UseErrorHandler :
func (t *Transport) UseErrorHandler(errorHandler func(http.ResponseWriter, *http.Request, error)) *Transport {
	t.errorHandler = errorHandler
	return t
}

// UseHandler :
func (t *Transport) UseHandler(handler func(messageType int, message []byte, fromSession *Session)) *Transport {
	t.handler = handler
	return t
}

// UsePath :
func (t *Transport) UsePath(path string) *Transport {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	t.path = path
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
