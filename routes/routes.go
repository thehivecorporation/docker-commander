package routes

import (
	"encoding/json"
	"fmt"

	"github.com/sayden/docker-commander/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/sayden/docker-commander/Godeps/_workspace/src/github.com/gorilla/websocket"
	"github.com/sayden/docker-commander/communications"
	"github.com/sayden/docker-commander/config"
	"github.com/sayden/docker-commander/discovery"
	"github.com/sayden/docker-commander/logger"
	"github.com/sayden/docker-commander/socket/receivers"
	"github.com/sayden/docker-commander/swarm"
)

var log = logger.WithField("routes")

// Init routes for the entire application
func Init(ginApp *gin.Engine, s swarm.Swarm, i discovery.InfoService) {
	//Common routes
	ginApp.LoadHTMLFiles("public/index.html")
	ginApp.GET("/", index)

	ginApp.Static("/js", "./public/js")
	ginApp.Static("/img", "./public/img")

	ginApp.GET("/ws", func(c *gin.Context) {
		commCh := make(chan *receivers.ReceiverPayload, 10)
		initWebSocket(c, s, i, commCh)
	})

}

func index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func initWebSocket(c *gin.Context, s swarm.Swarm, i discovery.InfoService, commCh chan *receivers.ReceiverPayload) {
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Errorf("Failed to set websocket upgrade: %+v", err)
		return
	}
	log.Info("Websocket started successfully")

	go messageHandler(conn, s, i, commCh)
	go channelHandler(commCh)
}

func messageHandler(conn *websocket.Conn, s swarm.Swarm, i discovery.InfoService, commCh chan *receivers.ReceiverPayload) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Error(err)
			break
		}

		var jsonMsg map[string]interface{}
		if err := json.Unmarshal(msg, &jsonMsg); err != nil {
			log.Fatal(err)
		}

		msgr := &communications.WebsocketCommunicator{conn}

		a := jsonMsg["action"].(string)
		rp := receivers.ReceiverPayload{&s, &i, msgr, a, ""}
		commCh <- &rp
	}
}

func channelHandler(r chan *receivers.ReceiverPayload) {
	for {
		rp := <-r

		switch rp.Action {
		case config.CONNECTION_ACTION_CLUSTER:
			//Ask entire cluster state
			go receivers.Cluster(rp)
		default:
			fmt.Println("Unknown message, returning cluster state")
			receivers.Cluster(rp)
		}
	}
}
