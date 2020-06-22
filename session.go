package websocket

import (
	"github.com/gorilla/websocket"
)

// Session :
type Session struct {
	Socket *websocket.Conn
	Custom map[string]interface{}
}

// SendJSON :
func (s *Session) SendJSON(message interface{}) error {
	return s.Socket.WriteJSON(message)
}

// SendMessage :
func (s *Session) SendMessage(message []byte) error {
	return s.Socket.WriteMessage(websocket.TextMessage, message)
}
