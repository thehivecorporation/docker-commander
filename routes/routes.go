package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sayden/docker-commander/entities/container"
	"github.com/sayden/docker-commander/entities/host"
)

// Init routes for the entire application
func Init(ginApp *gin.Engine) {
	//Common routes
	ginApp.LoadHTMLFiles("public/index.html")
	ginApp.GET("/", index)

	ginApp.GET("/ws", func(c *gin.Context) {
		webSocketHandler(c.Writer, c.Request)
	})

	//Initializes modules route	s
	host.InitializesRoutes(ginApp)
	container.InitializesRoutes(ginApp)
}

func index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func webSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var jsonMsg map[string]interface{}
		if err := json.Unmarshal(msg, &jsonMsg); err != nil {
			log.Fatal(err)
		}

		conn.WriteMessage(t, msg)
	}
}

func socketMsgHandler(j map[string]interface{}) {
	a := j["action"]

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
	}
}
