package socket

import (
	"fmt"
	"log"
	"testing"

	"github.com/sayden/docker-commander/discovery"
	"github.com/sayden/docker-commander/entities"
	"github.com/sayden/docker-commander/swarm"
)

func TestGetClusterInfo(t *testing.T) {
	s := swarm.GetClient(swarm.TYPE_MOCK_OK)
	i, err := getClusterInfo(s)
	if err != nil {
		t.Fatal("Error trying to get cluster info")
	}

	if i.Containers != 102 {
		t.Fatal("Containers assertion failed")
	}

	s = swarm.GetClient(swarm.TYPE_MOCK_ERROR)
	i, err = getClusterInfo(s)
	if err == nil {
		t.Fatal("There is no error after passing an incorrect json")
	}
}

func TestGetAgentsList(t *testing.T) {
	i := discovery.GetClient(swarm.TYPE_MOCK_OK)
	hs, err := getAgentsList(i)
	if err != nil {
		t.Fatal("Error trying to get hosts list")
	}

	if hs[0].IP != "ip1" {
		t.Fatal("Host[0].IP is not the expected result")
	}

	i = discovery.GetClient(swarm.TYPE_MOCK_ERROR)
	hs, err = getAgentsList(i)
	if err == nil {
		t.Fatal("No error after passing a malformed json")
	}

	if hs != nil {
		t.Fatal("Hosts list must not contain any data")
	}
}

func TestAddContainersForEachAgent(t *testing.T) {
	//Prepare
	s := swarm.GetClient(swarm.TYPE_MOCK_OK)
	i := discovery.GetClient(swarm.TYPE_MOCK_OK)
	rawAgents, err := i.ListHosts()

	agents := []entities.Agent{}
	for _, a := range rawAgents {
		agent := entities.Agent{
			IP: a.IP,
		}

		agents = append(agents, agent)
	}

	//Test
	err = addContainersForEachAgent(s, &agents)
	if err != nil {
		t.Fail()
	}

	if err != nil {
		t.Fatal("Assert failed when trying to get Agent Containers")
	}

	if len(agents) == 0 {
		log.Println(agents)
		t.Fatal("No hosts has been created")
	}

	if len(agents[0].Containers) == 0 {
		log.Println(agents[0])
		t.Fatal("No Container has been added to host 0")
	}
}

func TestAddImagesForEachAgent(t *testing.T) {
	//Prepare
	// var s swarm.Swarm = &swarm.HTTPClientMock{"http://a_host"}
	s := swarm.GetClient(swarm.TYPE_MOCK_OK)
	i := discovery.GetClient(swarm.TYPE_MOCK_OK)
	rawAgents, err := i.ListHosts()

	agents := []entities.Agent{}
	for _, a := range rawAgents {
		agent := entities.Agent{
			IP: a.IP,
		}

		agents = append(agents, agent)
	}

	//Test
	err = addImagesForEachAgent(s, &agents)
	if err != nil {
		t.Fail()
	}

	if err != nil {
		t.Fatal("Assert failed when trying to get Agent Containers")
	}

	if len(agents) == 0 {
		log.Println(agents)
		t.Fatal("No hosts has been created")
	}

	if len(agents[0].Images) == 0 {
		log.Println(agents[0])
		t.Fatal("No Images has been added to host '0'")
	}
}

func TestGetFullInfo(t *testing.T) {
	s := swarm.GetClient(swarm.TYPE_MOCK_OK)
	i := discovery.GetClient(swarm.TYPE_MOCK_OK)
	info := GetFullInfo(s, i)
	fmt.Println(info)
}
