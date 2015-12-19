package routes

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sayden/docker-commander/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/sayden/docker-commander/Godeps/_workspace/src/github.com/gorilla/websocket"
	"github.com/sayden/docker-commander/discovery"
	"github.com/sayden/docker-commander/logger"
	"github.com/sayden/docker-commander/socket"
	"github.com/sayden/docker-commander/swarm"
)

var log = logger.WithField("routes")

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
		log.Errorf("Failed to set websocket upgrade: %+v", err)
		return
	}
	log.Info("Websocket started successfully")

	go messageHandler(conn, s, i)
}

func messageHandler(conn *websocket.Conn, s swarm.Swarm, i discovery.InfoService) {
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Error(err)
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
				log.Error("Error trying to get cluster info", err)
				sendMessage(err, conn, msgType)
			} else {
				sendMessage(info, conn, msgType)
			}
		case "agent:containers":
			//TODO Gets containers of some swarm agent
			ipI := jsonMsg["ip"]
			if ip, ok := ipI.(string); ok {
				s = swarm.GetClientWithIP(ip)
				info, err := socket.GetContainers(s, ip)
				if err != nil {
					log.Error("Error trying to get containers list", err)
					sendMessage(err, conn, msgType)
				} else {
					sendMessage(info, conn, msgType)
				}
			} else {
				err = errors.New("Error trying to parse json with containers")
				log.Error(err)
				sendMessage(err, conn, msgType)
			}

		case "agent:images":
			//Gets images of some swarm agent
			ipI := jsonMsg["ip"]
			if ip, ok := ipI.(string); ok {
				s = swarm.GetClientWithIP(ip)
				info, err := socket.GetImages(s, ip)
				if err != nil {
					log.Error("Error trying to get images list", err)
					sendMessage(err, conn, msgType)
				} else {
					sendMessage(info, conn, msgType)
				}
			} else {
				err = errors.New("Error trying to parse json")
				log.Error(err)
				sendMessage(err, conn, msgType)
			}
		default:
			fmt.Println("Unknown message", msg)
		}

	}
}

func sendMessage(data interface{}, conn *websocket.Conn, msgType int) {
	//Ask cluster state
	var res []byte
	json, err := json.Marshal(data)
	if err != nil {
		log.Error("Error trying to parse interface")
		res = []byte(err.Error())
	} else {
		res = json
	}
	conn.WriteMessage(msgType, res)

}
