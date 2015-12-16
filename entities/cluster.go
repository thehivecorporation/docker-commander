package cluster

import "github.com/sayden/docker-commander/discovery"

//Cluster object holds information about the entire cluster in a semi
//structured way
type Cluster struct {
	Info  *map[string]interface{}
	Hosts *[]discovery.Node
}
