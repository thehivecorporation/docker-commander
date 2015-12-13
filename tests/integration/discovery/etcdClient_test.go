package discovery

import "testing"

import "github.com/sayden/docker-commander/discovery"

func TestEtcdClientIntegration(t *testing.T) {
	e := discovery.EtcdClient{Host: "http://192.168.1.35:2379"}
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

func TestWatch(t *testing.T) {
	// e := discovery.EtcdClient{Host: "http://192.168.1.35:2379"}
	// e.WatchHosts()
}
