package websocket

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/oscrud/oscrud"
)

// Interface checking
var _ oscrud.Transport = &Transport{}

// Transport Definition
const (
	TransportID oscrud.TransportID = "GORILLA_WEBSOCKET"
)

// Name :
func (t *Transport) Name() oscrud.TransportID {
	return TransportID
}

// Request :
func (t *Transport) Request(request *oscrud.Request, response interface{}) error {
	return oscrud.ErrTransportNotSupport
}

// Register :
func (t *Transport) Register(method string, endpoint string, handler oscrud.TransportHandler) {

}

// Start :
func (t *Transport) Start() error {
	if t.authHandler == nil {
		return errors.New("websocket: authHandler cannot be null")
	}

	if t.errorHandler == nil {
		return errors.New("websocket: errorHandler cannot be null")
	}

	if t.handler == nil {
		return errors.New("websocket: handler cannot be null")
	}

	addr := fmt.Sprintf(":%d", t.port)
	return http.ListenAndServe(addr, t)
}
