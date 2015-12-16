package routes

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/sayden/docker-commander/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/sayden/docker-commander/Godeps/_workspace/src/github.com/gorilla/websocket"
)

// Init routes for the entire application
func Init(ginApp *gin.Engine) {
	//Common routes
	ginApp.LoadHTMLFiles("public/index.html")
	ginApp.GET("/", index)

	ginApp.GET("/ws", initWebSocket)
}

func index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func initWebSocket(c *gin.Context) {
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}

	go messageHandler(conn)
}

func messageHandler(conn *websocket.Conn) {
	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var jsonMsg map[string]interface{}
		if err := json.Unmarshal(msg, &jsonMsg); err != nil {
			log.Fatal(err)
		}

		a := jsonMsg["action"]

		switch a {
		case "cluster":
			//TODO Ask cluster state
			fmt.Println("Not yet implemented")
		case "agent:containers":
			//TODO Gets containers of some swarm agent
			fmt.Println("Not yet implemented")
		case "agent:images":
			//TODO Gets images of some swarm agent
			fmt.Println("Not yet implemented")
		default:
			fmt.Println("Handle message")
			conn.WriteMessage(t, msg)
		}

	}
}
