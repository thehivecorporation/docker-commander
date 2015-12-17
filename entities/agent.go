package entities

import "github.com/sayden/docker-commander/Godeps/_workspace/src/github.com/samalba/dockerclient"

// Agent is a holder for a full structure
type Agent struct {
	IP         string
	Containers []dockerclient.Container
	Images     []dockerclient.Image
}
