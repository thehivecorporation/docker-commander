package socket

import (
	"encoding/json"
	"log"

	"github.com/sayden/docker-commander/discovery"
	"github.com/sayden/docker-commander/parsers"
	"github.com/sayden/docker-commander/swarm"
)

func getClusterInfo() {
	s := swarm.GetClient()

	// Cluster info
	r, err := s.GetInfo()
	if err != nil {
		log.Fatal(err)
		return
	}

	var cData map[string]interface{}

	if err := json.Unmarshal(r, &cData); err != nil {
		log.Fatal(err)
	}

	// c := cluster.Cluster{Info: &cData}

	//Get every host
	d := discovery.GetClient()
	hs, err := d.ListHosts()
	if err != nil {
		log.Fatal(err)
		return
	}

	//Map byte Nodes to DockerClientNodes
	var a []parsers.DockerClientNode
	for _, d := range hs {
		h := parsers.DockerClientNode{
			IP: d.IP,
		}

		a = append(a, h)
	}

	p := parsers.DockerClientParser{}

	//Foreach host, get its containers
	for _, h := range a {
		csb, err := s.GetContainers()
		if err != nil {
			return
		}

		cs, err := p.ParseContainer(&csb)
		if err != nil {
			return
		}
		h.Containers = cs
	}

	//Foreach host, get its images
	for _, h := range a {
		isb, err := s.GetImages()
		if err != nil {
			return
		}

		is, err := p.ParseImages(&isb)
		if err != nil {
			return
		}
		h.Images = is
	}
}
