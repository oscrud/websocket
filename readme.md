# Introduction

Websocket transport for oscrud and build using [gorilla_websocket](https://github.com/gorilla/websocket).

# Usage

| Method                                      | Usage                                              |
| ------------------------------------------- | -------------------------------------------------- |
| UsePort(port int)                           | Specify which port to receiving requests           |
| UseUpgrader(upgrader websocket.Upgrader)    | Customize websocket upgrader instead using default |
| UseErrorHandler(handler ErrorHandler)       | Customize error handler instead using default      |
| UseHandler(handler Handler)                 | Customize message handler and it's required        |
| UseAuthHandler(handler AuthHandler)         | Customize auth handler instead using default       |
| GetSession(key string)                      | Get session by id                                  |
| UpdateSession(key string, session *Session) | Update session by id                               |

# Example

```go
package main

import (
    ws "github.com/oscrud/websocket"
	"github.com/gorilla/websocket"
)

func main() {
    t := ws.NewWebsocket()
    t.UsePort(3000) // specify port
    t.UseHandler(
        func(messageType int, message []byte, session *ws.Session) {
            bytes := []byte("hello world reply from server")
            session.SendMessage(bytes)
        },
    )
    t.UseUpgrader(
        websocket.Upgrader{
			CheckOrigin: func(request *http.Request) bool {
				return true
			},
        },
    )
    t.UseAuthHandler(
        // mainly to specify which user in the conncetion
		func(res http.ResponseWriter, req *http.Request) (string, *Session, error) {
			return uuid.New().String(), nil, nil 
		},
    )
}

```