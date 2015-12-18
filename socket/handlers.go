package socket

import (
	"log"

	"github.com/sayden/docker-commander/Godeps/_workspace/src/github.com/samalba/dockerclient"
	"github.com/sayden/docker-commander/discovery"
	"github.com/sayden/docker-commander/entities"
	"github.com/sayden/docker-commander/swarm"
)

func getClusterInfo(s swarm.Swarm) (*dockerclient.Info, error) {
	// Cluster info
	i, err := s.ListInfo()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &i, nil
}

func getAgentsList(i discovery.InfoService) ([]entities.Agent, error) {
	//Get every host
	hs, err := i.ListHosts()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//Map byte Nodes to DockerClientNodes
	var agents []entities.Agent
	for _, h := range hs {
		a := entities.Agent{
			IP: h.IP,
		}

		agents = append(agents, a)
	}

	return agents, nil
}

func addContainersForEachAgent(s swarm.Swarm, ag *[]entities.Agent) error {
	if len(*ag) == 0 {
		log.Println("ERROR: There are no agents in ag parameter")
	}

	agents := *ag

	//Foreach host, get its containers
	for i := range agents {
		h := &agents[i]
		cs, err := s.ListContainers()
		if err != nil {
			return err
		}

		h.Containers = cs
	}

	return nil
}

func addImagesForEachAgent(s swarm.Swarm, ag *[]entities.Agent) error {
	if len(*ag) == 0 {
		log.Println("ERROR: There are no agents in ag parameter")
	}

	agents := *ag

	//Foreach host, get its containers
	for i := range agents {
		h := &agents[i]
		is, err := s.ListImages()
		if err != nil {
			return err
		}

		h.Images = is
	}

	return nil
}

// GetFullInfo joins all available info of the cluster in a single response
func GetFullInfo(s swarm.Swarm, i discovery.InfoService) (entities.Overall, error) {
	cluster, err := getClusterInfo(s)
	if err != nil {
		log.Println(err)
		return entities.Overall{}, err
	}
	agentsNodes, err := getAgentsList(i)
	if err != nil {
		log.Println(err)
		return entities.Overall{}, err
	}

	err = addContainersForEachAgent(s, &agentsNodes)
	if err != nil {
		return entities.Overall{}, err
	}

	err = addImagesForEachAgent(s, &agentsNodes)
	if err != nil {
		return entities.Overall{}, err
	}

	return entities.Overall{
		Cluster: *cluster,
		Agents:  agentsNodes,
	}, nil
}
