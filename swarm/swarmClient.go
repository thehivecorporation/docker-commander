package swarm

import "github.com/samalba/dockerclient"

// Swarm is to uncouple http.Client library
type Swarm interface {
	ListInfo() (dockerclient.Info, error)
	ListContainers() ([]dockerclient.Container, error)
	ListImages() ([]dockerclient.Image, error)
}
