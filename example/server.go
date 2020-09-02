package main

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/oscrud/oscrud"
	ws "github.com/oscrud/websocket"
)

// SendMessage :
func SendMessage(ctx *oscrud.Context) *oscrud.Context {

	var i struct {
		Message string `qm:"message"`
	}

	if err := ctx.BindAll(&i); err != nil {
		return ctx.JSON(500, map[string]interface{}{
			"ok":      false,
			"message": err,
		})
	}

	transport, ok := ctx.GetTransport(ws.TransportID)
	if ok {
		socket := transport.(*ws.Transport)
		session := socket.GetSession("debug")
		if session != nil {
			if err := session.SendMessage([]byte(i.Message)); err != nil {
				return ctx.JSON(500, map[string]interface{}{
					"ok":      false,
					"message": err,
				})
			}
			return ctx.JSON(200, map[string]interface{}{
				"ok": true,
			})
		}
		return ctx.JSON(404, map[string]interface{}{
			"ok":      false,
			"message": "session not found",
		})
	}
	return ctx.JSON(404, map[string]interface{}{
		"ok":      false,
		"message": "transport not found",
	})
}

func main() {
	server := oscrud.NewOscrud()
	server.RegisterTransport(
		ws.New().UsePort(3000).UseHandler(
			func(messageType int, message []byte, session *ws.Session) {
				bytes := []byte("hello world reply from server")
				session.SendJSON(map[string]interface{}{
					"ok":     true,
					"status": 200,
				})
				session.SendMessage(bytes)
			},
		).UseAuthHandler(
			func(http.ResponseWriter, *http.Request) (string, *ws.Session, error) {
				return "debug", nil, nil
			},
		).UseUpgrader(websocket.Upgrader{
			HandshakeTimeout: time.Duration(30 * time.Second),
			CheckOrigin: func(request *http.Request) bool {
				return true
			},
		}),
	)

	server.RegisterEndpoint("GET", "/websocket/:message", SendMessage)
	server.Start()
}
