package socket

import (
	"log"
	"testing"

	"github.com/sayden/docker-commander/discovery"
	"github.com/sayden/docker-commander/parsers"
	"github.com/sayden/docker-commander/swarm"
)

func TestGetClusterInfo(t *testing.T) {
	var s swarm.Swarm = &swarm.HTTPClientMock{"http://a_host"}
	i, err := getClusterInfo(s)
	if err != nil {
		t.Fatal("Error trying to get cluster info")
	}

	if i.Containers != 102 {
		t.Fatal("Containers assertion failed")
	}

	s = &swarm.HTTPClientMockError{"http://a_host"}
	i, err = getClusterInfo(s)
	if err == nil {
		t.Fatal("There is no error after passing an incorrect json")
	}
}

func TestGetHostsList(t *testing.T) {
	dOk := discovery.MockDiscoveryOk{"http://a_host"}
	hs, err := getAgentsList(&dOk)
	if err != nil {
		t.Fatal("Error trying to get hosts list")
	}

	if hs[0].IP != "ip1" {
		t.Fatal("Host[0].IP is not the expected result")
	}

	dError := discovery.MockDiscoveryError{"http://a_host"}
	hs, err = getAgentsList(&dError)
	if err == nil {
		t.Fatal("No error after passing a malformed json")
	}

	if hs != nil {
		t.Fatal("Hosts list must not contain any data")
	}
}

func TestAddContainersForEachAgent(t *testing.T) {
	var s swarm.Swarm = &swarm.HTTPClientMock{"http://a_host"}
	dOk := discovery.MockDiscoveryOk{"http://a_host"}
	agents, err := getAgentsList(&dOk)
	if err != nil {
		t.Fatal("Error trying to get hosts list")
	}

	err = addContainersForEachAgent(s, &agents)

	if err != nil {
		t.Fatal("Assert failed when trying to get Agent Containers")
	}

	if len(agents) == 0 {
		log.Println(agents)
		t.Fatal("No hosts has been created")
	}

	if agents[0].IP != "ip1" {
		log.Println(agents[0])
		t.Fatal("No Container has been added to host 0")
	}

	s = &swarm.HTTPClientMockError{"http://a_host"}
	agents = make([]parsers.DockerClientNode, 10)
	err = addContainersForEachAgent(s, &agents)

	if err == nil {
		t.Fatal("An error was expected")
	}
}
