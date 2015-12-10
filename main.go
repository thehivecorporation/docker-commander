package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sayden/docker-commander/config"
	"github.com/sayden/docker-commander/routes"
)

func main() {
	ginApp := gin.Default()

	routes.Init(ginApp)

	ginApp.Run(config.APP_PORT)
}
