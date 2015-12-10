package api

import (
    "github.com/gin-gonic/gin"
)

func Init(app *gin.Engine){
  app.GET("/api", index)
}

func index(c *gin.Context){
    content := gin.H{"Hello": "Api"}
    c.JSON(200, content)
}
