package websocket

import (
	"github.com/gorilla/websocket"
)

// Session :
type Session struct {
	Socket *websocket.Conn
	State  map[string]interface{}
}

// Get :
func (s *Session) Get(key string) interface{} {
	return s.State[key]
}

// Set :
func (s *Session) Set(key string, data interface{}) {
	s.State[key] = data
}

// SendJSON :
func (s *Session) SendJSON(message interface{}) error {
	return s.Socket.WriteJSON(message)
}

// SendMessage :
func (s *Session) SendMessage(message []byte) error {
	return s.Socket.WriteMessage(websocket.TextMessage, message)
}
