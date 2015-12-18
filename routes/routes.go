package routes

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/sayden/docker-commander/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/sayden/docker-commander/Godeps/_workspace/src/github.com/gorilla/websocket"
	"github.com/sayden/docker-commander/discovery"
	"github.com/sayden/docker-commander/socket"
	"github.com/sayden/docker-commander/swarm"
)

// Init routes for the entire application
func Init(ginApp *gin.Engine, s swarm.Swarm, i discovery.InfoService) {
	//Common routes
	ginApp.LoadHTMLFiles("public/index.html")
	ginApp.GET("/", index)

	ginApp.GET("/ws", func(c *gin.Context) {
		initWebSocket(c, s, i)
	})
}

func index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func initWebSocket(c *gin.Context, s swarm.Swarm, i discovery.InfoService) {
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}

	go messageHandler(conn, s, i)
}

func messageHandler(conn *websocket.Conn, s swarm.Swarm, i discovery.InfoService) {
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
			//Ask cluster state
			info, err := socket.GetFullInfo(s, i)
			if err != nil {
				conn.WriteMessage(t, []byte(err.Error()))
			} else {
				json, err := json.Marshal(info)
				var res []byte
				if err != nil {
					res = []byte(err.Error())
				} else {
					res = json
				}
				conn.WriteMessage(t, res)
			}
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
