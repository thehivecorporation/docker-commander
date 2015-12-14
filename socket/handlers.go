package socket

import (
	"log"

	"github.com/sayden/docker-commander/discovery"
	"github.com/sayden/docker-commander/swarm"
)

func handlerCluster() {
	s := swarm.GetClient()

	// Cluster info
	r, err := s.GetHosts()
	if err != nil {
		log.Fatal(err)
		return
	}

	//Foreach host, get containers
	d := discovery.GetClient()
	hs, err := d.ListHosts()
  if err != nil{
    log.Fatal(err)
    return
  }

  var containers := []map[string]interface{}

  for _, h := range hs {
    
  }
}
