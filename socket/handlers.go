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
	byt, err := s.ListInfo()
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

func getAgentsList(i discovery.InfoService) ([]parsers.DockerClientNode, error) {
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

	return agents, nil
}

func addContainersForEachAgent(s swarm.Swarm, ag *[]parsers.DockerClientNode) error {
	if len(*ag) == 0 {
		log.Println("ERROR: There are no agents in the request")
	}

	//Foreach host, get its containers
	for _, h := range *ag {
		csb, err := s.ListContainers()
		if err != nil {
			return err
		}

		p := parsers.DockerClientParser{}
		cs, err := p.ParseContainer(&csb)
		if err != nil {
			return err
		}
		h.Containers = cs
	}

	return nil
}
