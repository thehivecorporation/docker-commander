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
	i, err := s.ListInfo()
	if err != nil {
		log.Println(err)
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

func addContainersForEachAgent(s swarm.Swarm, ag *[]parsers.DockerClientNode, p parsers.ContainerParser) error {
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

//
// func addImagesForEachAgent(s swarm.Swarm, ag *[]parsers.DockerClientNode, p parsers.ImageParser) error {
// 	if len(*ag) == 0 {
// 		log.Println("ERROR: There are no agents in ag parameter")
// 	}
//
// 	agents := *ag
//
// 	//Foreach host, get its containers
// 	for i := range agents {
// 		h := &agents[i]
// 		isb, err := s.ListImages()
// 		if err != nil {
// 			return err
// 		}
//
// 		is, err := p.ParseImages(&isb)
// 		if err != nil {
// 			return err
// 		}
// 		h.Images = is
// 	}
//
// 	return nil
// }
//
// // GetFullInfo joins all available info of the cluster in a single response
// func GetFullInfo() int {
// 	s := swarm.GetClient(swarm.TYPE_ENV)
// 	i := discovery.GetClient(swarm.TYPE_ENV)
//
// 	cluster, err := getClusterInfo(s)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	agentsNodes, err := getAgentsList(i)
// 	if err != nil {
// 		log.Println(err)
// 	}
//
// 	parsers := parsers.DockerClientParser{}
// 	addContainersForEachAgent(s, &agentsNodes, parsers)
//
// 	addImagesForEachAgent(s, &agentsNodes, parsers)
//
// 	info := map[string]interface{}{}
// 	info["clusterInfo"] = cluster
// 	info["agents"] = agents
//
// 	return info
// }
