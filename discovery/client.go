package discovery

import (
	"github.com/sayden/docker-commander/logger"
)

// Node for representing connected clients to discovery service
type Node struct {
	IP         string
	Containers []map[string]interface{}
	Images     []map[string]interface{}
}

// InfoService interface definition to access discovery services
type InfoService interface {
	ListHosts() ([]Node, error)
	WatchHosts()
}

var log = logger.WithField("discovery")
