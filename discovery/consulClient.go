package discovery

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

// ConsulClient uses consul as discovery service
type ConsulClient struct {
	Host string
}

// WatchHosts TODO use some goroutines here
func (c *ConsulClient) WatchHosts() {
	log.Fatal("Not implemented yet")
}

// ListHosts TODO returns connected swarm nodes
func (c *ConsulClient) ListHosts() ([]Node, error) {
	conf := api.DefaultConfig()
	conf.Address = c.Host
	client, err := api.NewClient(conf)
	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	// Lookup the pair
	pairs, _, err := kv.List("/docker/swarm/nodes/", nil)
	if err != nil {
		panic(err)
	}

	nodesArray := make([]Node, len(pairs))
	for i, pair := range pairs {
		nodesArray[i] = Node{IP: pair.Key}
	}

	return nodesArray, nil
}
