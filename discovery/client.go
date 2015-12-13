package discovery

// Client abstract
type Client struct{}

// Node for representing connected clients to discovery service
type Node struct {
	IP string
}

// InfoService interface definition to access discovery services
type InfoService interface {
	ListHosts(serviceURL string) ([]Node, error)
}

// ListHosts returns the hosts joined in cluster
func (c *Client) ListHosts(serviceURL string, i InfoService) ([]Node, error) {
	r, err := i.ListHosts(serviceURL)

	if err != nil {
		return nil, err
	}

	return r, nil
}
