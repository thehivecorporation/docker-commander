package communications

import (
	"encoding/json"

	"github.com/sayden/docker-commander/Godeps/_workspace/src/github.com/gorilla/websocket"
	"github.com/sayden/docker-commander/logger"
)

var log = logger.WithField("communications")

// FrontMessenger represents an object that can send messages to front...
// abstracting how this communication is done (socket, ajax, etc).
type FrontMessenger interface {
	FrontMessage(resp *ConnectionResponse)
}

//WebsocketCommunicator satisfies FrontMessenger interface.
type WebsocketCommunicator struct {
	Conn *websocket.Conn
}

//FrontMessage will use socket communication to send messages to frontend
func (w *WebsocketCommunicator) FrontMessage(resp *ConnectionResponse) {
	var res []byte
	json, err := json.Marshal(resp)
	if err != nil {
		log.Error("Error trying to parse interface")
		res = []byte(err.Error())
	} else {
		res = json
	}
	w.Conn.WriteMessage(1, res)
}
