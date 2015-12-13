package commander

import "github.com/sayden/docker-commander/entities/host"

var hosts []host.Host

// GetHosts gives you every host that is registered
func GetHosts() []host.Host {
	return []host.Host{}
}
