package discovery

// Node for representing connected clients to discovery service
type Node struct {
	IP string
}

// InfoService interface definition to access discovery services
type InfoService interface {
	ListHosts() ([]Node, error)
	WatchHosts()
}

// ListHosts returns the hosts joined in cluster
func ListHosts(i InfoService) ([]Node, error) {
	r, err := i.ListHosts()

	if err != nil {
		return nil, err
	}

	return r, nil
}

// WatchHosts is to attach a closure to any change event in the hosts list
func WatchHosts(i InfoService) {
	i.WatchHosts()
}
