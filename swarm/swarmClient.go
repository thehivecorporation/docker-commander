package swarm

// Swarm is to uncouple http.Client library
type Swarm interface {
	GetInfo() ([]byte, error)
	GetContainers() ([]byte, error)
	GetImages() ([]byte, error)
}

// GetInfo returns hosts, agents, managers... in cluster
func GetInfo(s Swarm) ([]byte, error) {
	return s.GetInfo()
}

// GetContainers returns the Containers of a specific host
func GetContainers(s Swarm) ([]byte, error) {
	return s.GetContainers()
}

// GetImages returns the Images of a specific host
func GetImages(s Swarm) ([]byte, error) {
	return s.GetImages()
}
