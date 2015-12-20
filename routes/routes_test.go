package routes

import (
	"fmt"
	"testing"

	"github.com/sayden/docker-commander/Godeps/_workspace/src/github.com/gorilla/websocket"
)

type conn struct{}

func (c *conn) WriteMessage(messageType int, data []byte) error {
	fmt.Println("Something")
	return nil
}

func TestSendMessage(t *testing.T) {
	mockConn := &websocket.Conn{}
	res := sendMessage(struct {
		Hello string
		World string
	}{
		Hello: "Hello",
		World: "World",
	}, mockConn, 1)

	log.Info(res)
}
