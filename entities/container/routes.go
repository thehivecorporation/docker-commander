package container

import "github.com/gin-gonic/gin"

// InitializesRoutes to access API endpoints
func InitializesRoutes(app *gin.Engine) {
	app.GET("/container", containerList)
	app.GET("/container/:ID", getContainer)

	api := app.Group("/api")
	{
		api.GET("/container", apiList)
		api.GET("/container/:ID", apiGet)
		api.POST("/container", apiAdd)
	}

}

func getContainer(c *gin.Context) {

}

func containerList(c *gin.Context) {

}

func apiList(c *gin.Context) {
	hs := []Container{}
	c.JSON(200, hs)
}

func apiGet(c *gin.Context) {
	h := Container{}
	c.JSON(200, h)
}

func apiAdd(c *gin.Context) {
	h := Container{}
	c.JSON(200, h)
}
