package socket

import (
	"testing"

	"github.com/sayden/docker-commander/discovery"
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
	var s swarm.Swarm = &swarm.HTTPClientMock{"http://a_host"}
	dOk := discovery.MockDiscoveryOk{"http://a_host"}
	hs, err := getHostList(s, &dOk)
	if err != nil {
		t.Fatal("Error trying to get hosts list")
	}

	s = &swarm.HTTPClientMockError{"http://a_host"}
	dError := discovery.MockDiscoveryError{"http://a_host"}
	hs, err = getHostList(s, &dError)
	if err == nil {
		t.Fatal("No error after passing a malformed json")
	}

	if hs != nil {
		t.Fatal("Hosts list must not contain any data")
	}
}
