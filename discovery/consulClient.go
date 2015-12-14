package discovery

import (
	"errors"
	"log"
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
	//TODO
	return nil, errors.New("Not implemented yet")
}
