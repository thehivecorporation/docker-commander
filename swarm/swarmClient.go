package swarm

// Swarm is to uncouple http.Client library
type Swarm interface {
	GetHosts() ([]byte, error)
	GetContainers() ([]byte, error)
	GetImages() ([]byte, error)
}

// GetHosts returns hosts, agents, managers... in cluster
func GetHosts(s Swarm) ([]byte, error) {
	return s.GetHosts()
}

// GetContainers returns the Containers of a specific host
func GetContainers(s Swarm) ([]byte, error) {
	return s.GetContainers()
}

// GetImages returns the Images of a specific host
func GetImages(s Swarm) ([]byte, error) {
	return s.GetImages()
}
