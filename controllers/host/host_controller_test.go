package host_controller

import (
  "testing"
  "github.com/sayden/docker-commander/entities/host"
)

func TestIsHostAliveFunc(t *testing.T){
  h := host.Host {
    Ip:"127.0.0.1",
    Title:"My host",
  }

  a := isHostAlive(&h)

  if (!a){
    t.Fatal("Not alive")
  } 
}
