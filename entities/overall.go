package entities

import "github.com/samalba/dockerclient"

//Overall is an object to store overall information about the cluster
type Overall struct {
	Cluster dockerclient.Info
	Agents  []Agent
}
