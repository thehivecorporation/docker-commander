package socket

import (
	"log"
	"testing"

	"github.com/sayden/docker-commander/discovery"
	"github.com/sayden/docker-commander/entities"
	"github.com/sayden/docker-commander/swarm"
)

func TestGetClusterInfo(t *testing.T) {
	s := swarm.HTTPClientMock{"http://a_host"}
	i, err := getClusterInfo(&s)
	if err != nil {
		t.Fatal("Error trying to get cluster info")
	}

	if i.Containers != 102 {
		t.Fatal("Containers assertion failed")
	}

	sErr := swarm.HTTPClientMockError{"http://a_host"}
	i, err = getClusterInfo(&sErr)
	if err == nil {
		t.Fatal("There is no error after passing an incorrect json")
	}
}

func TestGetAgentsList(t *testing.T) {
	i := discovery.MockDiscoveryOk{"http://a_host"}
	hs, err := getAgentsList(&i)
	if err != nil {
		t.Fatal("Error trying to get hosts list")
	}

	if hs[0].IP != "ip1" {
		t.Fatal("Host[0].IP is not the expected result")
	}

	iErr := discovery.MockDiscoveryError{"http://a_host"}
	hs, err = getAgentsList(&iErr)
	if err == nil {
		t.Fatal("No error after passing a malformed json")
	}

	if hs != nil {
		t.Fatal("Hosts list must not contain any data")
	}
}

func TestAddContainersForEachAgent(t *testing.T) {
	//Prepare
	i := discovery.MockDiscoveryOk{"http://a_host"}
	rawAgents, err := i.ListHosts()

	agents := []entities.Agent{}
	for _, a := range rawAgents {
		agent := entities.Agent{
			IP: a.IP,
		}

		agents = append(agents, agent)
	}

	//Test
	s := swarm.HTTPClientMock{"http://a_host"}
	err = addContainersForEachAgent(&s, &agents)
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

func TestGetContainers(t *testing.T) {
	s := swarm.HTTPClientMock{"http://a_host"}
	ip := "http://non_important_string"

	cs, err := GetContainers(&s, ip)
	if err != nil {
		t.Fatal("Error trying to get containers")
	}

	if len(*cs) == 0 {
		t.Fatal("Containers must be over `0`")
	}
}

func TestAddImagesForEachAgent(t *testing.T) {
	//Prepare
	s := swarm.HTTPClientMock{"http://a_host"}
	i := discovery.MockDiscoveryOk{"http://a_host"}
	rawAgents, err := i.ListHosts()

	agents := []entities.Agent{}
	for _, a := range rawAgents {
		agent := entities.Agent{
			IP: a.IP,
		}

		agents = append(agents, agent)
	}

	//Test
	err = addImagesForEachAgent(&s, &agents)
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
	s := swarm.HTTPClientMock{"http://a_host"}
	i := discovery.MockDiscoveryOk{"http://a_host"}
	info, err := GetFullInfo(&s, &i)
	if err != nil {
		t.Fatal(err)
	}

	//Test some info
	if info.Cluster.Containers != 102 {
		t.Fail()
	}

	//Test there are some agents
	if len(info.Agents) == 0 {
		t.Fail()
	}

	if info.Agents[0].IP != "ip1" {
		t.Fail()
	}

	if info.Agents[0].Containers[0].Id != "6514d918bc01b0df8098053086c72cf9e4e3f0b166987b5b5db78d1f6b783ea8" {
		t.Fail()
	}

	if info.Agents[0].Images[0].Id != "e9ff33e7e5b9a683ab735bbe99450c50bd0b64c4e414d12c94ff93b345d3bb18" {
		t.Fail()
	}

}
