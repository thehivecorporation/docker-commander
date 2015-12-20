package swarm

import (
	"github.com/sayden/docker-commander/Godeps/_workspace/src/github.com/samalba/dockerclient"
	"github.com/sayden/docker-commander/logger"
)

// Swarm is to uncouple http.Client library
type Swarm interface {
	ListInfo() (dockerclient.Info, error)
	ListContainers() ([]dockerclient.Container, error)
	ListImages() ([]dockerclient.Image, error)
}

var log = logger.WithField("swarm")
