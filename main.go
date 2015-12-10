package main

import (
    "github.com/gin-gonic/gin"
    "github.com/sayden/docker-commander/routes/app"
    "github.com/sayden/docker-commander/routes/api"
)

func main(){
  ginApp := gin.Default()

  //Init Front routes
  app.Init(ginApp)

  //Init api routes
  api.Init(ginApp)

  ginApp.Run(":8000")
}
