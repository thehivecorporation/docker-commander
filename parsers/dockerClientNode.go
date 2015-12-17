package parsers

import "github.com/samalba/dockerclient"

// DockerClientNode is a holder for a full structure
type DockerClientNode struct {
	IP         string
	Containers []dockerclient.Container
	Images     []dockerclient.Image
}
