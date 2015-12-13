// Package host entity. A host is a remote machine that has containers. Has a 1
// to many relationship with containers
package host

import "github.com/sayden/docker-commander/entities/container"

// Host describe a remote machine hosting Docker containers
type Host struct {
	ID         int64  `form:"id" json:"id" binding:"required"`
	Title      string `form:"title" json:"title" binding:"required"`
	IP         string `form:"ip" json:"ip" binding:"required"`
	Containers []container.Container
}

// GetContainers access host to retrieve a list of containers to user
func GetContainers() []container.Container {
	return []container.Container{}
}

// GetImages gives the host's installed images in JSON format
func GetImages() []container.Container {
	return []container.Container{}
}

// GetRunningContainers gives every running container in host in JSON format
func GetRunningContainers() []container.Container {
	return []container.Container{}
}
