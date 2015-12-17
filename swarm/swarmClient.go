package swarm

// Swarm is to uncouple http.Client library
type Swarm interface {
	ListInfo() ([]byte, error)
	ListContainers() ([]byte, error)
	ListImages() ([]byte, error)
}
