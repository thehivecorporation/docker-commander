package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sayden/docker-commander/entities/container"
	"github.com/sayden/docker-commander/entities/host"
)

// Init routes for the entire application
func Init(ginApp *gin.Engine) {
	//Common routes
	ginApp.GET("/", index)

	//Initializes modules routes
	host.InitializesRoutes(ginApp)
	container.InitializesRoutes(ginApp)
}

func index(c *gin.Context) {
	content := gin.H{"Hello": "World"}
	c.JSON(200, content)
}
