package host

import "github.com/gin-gonic/gin"

// InitializesRoutes to access API endpoints
func InitializesRoutes(app *gin.Engine) {
	app.GET("/host", hostList)
	app.GET("/host/:ID", getHost)

	api := app.Group("/api")
	{
		api.GET("/host", apiList)
		api.GET("/host/:ID", apiGet)
		api.POST("/host", apiAdd)
	}

}

func getHost(c *gin.Context) {

}

func hostList(c *gin.Context) {

}

func apiList(c *gin.Context) {
	hs := []Host{}
	c.JSON(200, hs)
}

// Get a host by its ID
func apiGet(c *gin.Context) {
	h := Host{}
	c.JSON(200, h)
}

// Add inserts a new host to watch in the list
func apiAdd(c *gin.Context) {
	h := Host{}
	c.JSON(200, h)
}
