package discovery

import (
	"testing"

	"github.com/sayden/docker-commander/tests/integration/discovery/Godeps/_workspace/src/github.com/sayden/docker-commander/config"
	"github.com/sayden/docker-commander/tests/integration/discovery/Godeps/_workspace/src/github.com/sayden/docker-commander/discovery"
)

func TestConsulClientIntegration(t *testing.T) {
	e := discovery.ConsulClient{Host: config.CONSUL_HOST_DEVELOPMENT}
	r, err := e.ListHosts()

	if err != nil {
		t.Fatalf("Error listing hosts %q", err)
	}

	if len(r) == 0 {
		t.Log("Swarm nodes list is empty")
	}

	for _, n := range r {
		t.Logf("Node %q in list", n.IP)
	}
}
