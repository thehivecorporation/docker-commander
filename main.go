package main

import (
	"github.com/sayden/docker-commander/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/sayden/docker-commander/config"
	"github.com/sayden/docker-commander/discovery"
	"github.com/sayden/docker-commander/routes"
	"github.com/sayden/docker-commander/swarm"
)

func main() {
	ginApp := gin.Default()

	s := swarm.GetClient()
	i := discovery.GetClient()

	routes.Init(ginApp, s, i)

	ginApp.Run(config.APP_PORT)
}
