package main

import (
    "github.com/gin-gonic/gin"
    "docker-commander/routes/index"
)

func main(){
  app := gin.Default()
  app.GET("/", index.Index)
  app.Run(":8000")
}
