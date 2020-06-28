package websocket

import (
	"net/http"
)

// ServerHTTP :
func (t *Transport) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	uid, session, err := t.authHandler(res, req)
	if err != nil {
		t.errorHandler(res, req, err)
		return
	}

	socket, err := t.upgrader.Upgrade(res, req, nil)
	if err != nil {
		t.errorHandler(res, req, err)
		return
	}
	defer socket.Close()
	if session == nil {
		session = new(Session)
	}
	session.Socket = socket
	t.sessions[uid] = session

	for {
		messageType, message, err := socket.ReadMessage()
		if err != nil {
			t.errorHandler(res, req, err)
			break
		}
		t.handler(messageType, message, session)
	}

	go t.closeHandler(session)
	delete(t.sessions, uid)
}
