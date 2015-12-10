package app

import (
    "github.com/gin-gonic/gin"
)

func Init(app *gin.Engine) {
  app.GET("/", index)
}

func index(c *gin.Context) {
    content := gin.H{"Hello": "World"}
    c.JSON(200, content)
}
