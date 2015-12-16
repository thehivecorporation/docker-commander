package socket

import (
	"encoding/json"
	"log"

	"github.com/sayden/docker-commander/discovery"
	"github.com/sayden/docker-commander/entities"
	"github.com/sayden/docker-commander/swarm"
)

func getClusterInfo() {
	s := swarm.GetClient()

	// Cluster info
	r, err := s.GetHosts()
	if err != nil {
		log.Fatal(err)
		return
	}

	var cData map[string]interface{}

	if err := json.Unmarshal(r, &cData); err != nil {
		log.Fatal(err)
	}

	c := cluster.Cluster{Info: &cData}

	//Get every host
	d := discovery.GetClient()
	hs, err := d.ListHosts()
	if err != nil {
		log.Fatal(err)
		return
	}

	//Foreach host, get its containers
	for _, h := range hs {

	}
}
