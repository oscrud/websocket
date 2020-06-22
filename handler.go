package websocket

import (
	"errors"
	"net/http"
)

// socket :
func (t *Transport) socketHandler(res http.ResponseWriter, req *http.Request) {
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
	delete(t.sessions, uid)
}

// redirect :
func (t *Transport) redirectHandler(res http.ResponseWriter, req *http.Request) {
	if t.notFoundHandler != nil {
		t.notFoundHandler(res, req)
		return
	}

	if t.redirect {
		http.Redirect(res, req, t.path, 302)
		return
	}

	err := errors.New("websocket: not found")
	t.errorHandler(res, req, err)
}
