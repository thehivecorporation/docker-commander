package socket

import (
	"log"

	"github.com/samalba/dockerclient"
	"github.com/sayden/docker-commander/discovery"
	"github.com/sayden/docker-commander/parsers"
	"github.com/sayden/docker-commander/swarm"
)

func getClusterInfo(s swarm.Swarm) (*dockerclient.Info, error) {
	// Cluster info
	byt, err := s.GetInfo()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	p := parsers.DockerClientParser{}
	i, err := p.ParseInfo(&byt)
	if err != nil {
		log.Println("ERROR:", err)
		return nil, err
	}

	return &i, nil
}

func getHostList(s swarm.Swarm, i discovery.InfoService) (*[]parsers.DockerClientNode, error) {
	//Get every host
	hs, err := i.ListHosts()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//Map byte Nodes to DockerClientNodes
	var agents []parsers.DockerClientNode
	for _, d := range hs {
		h := parsers.DockerClientNode{
			IP: d.IP,
		}

		agents = append(agents, h)
	}

	return &agents, nil
}

func addAgentContainers(s swarm.Swarm, ag *[]parsers.DockerClientNode) {
	//Foreach host, get its containers
	for _, h := range *ag {
		csb, err := s.GetContainers()
		if err != nil {
			return
		}

		p := parsers.DockerClientParser{}
		cs, err := p.ParseContainer(&csb)
		if err != nil {
			return
		}
		h.Containers = cs
	}
}

func addAgentImages(s swarm.Swarm, ag *[]parsers.DockerClientNode) {
	//Foreach host, get its images
	for _, h := range *ag {
		isb, err := s.GetImages()
		if err != nil {
			return
		}

		p := parsers.DockerClientParser{}
		is, err := p.ParseImages(&isb)
		if err != nil {
			return
		}
		h.Images = is
	}
}

// GetFullInfo joins all available info of the cluster in a single response
// func GetFullInfo() {
// 	s := swarm.GetClient()
// }
