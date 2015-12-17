package discovery

import "errors"

//MockDiscoveryOk is a mock to produce expected results
type MockDiscoveryOk struct {
	Host string
}

//MockDiscoveryError is a mock to produce errors
type MockDiscoveryError struct {
	Host string
}

// ListHosts mocked
func (c *MockDiscoveryOk) ListHosts() ([]Node, error) {
	ns := []Node{Node{IP: "ip1"}, Node{IP: "ip2"}}

	return ns, nil
}
func (c *MockDiscoveryOk) WatchHosts() {}

//ListHosts Error
func (c *MockDiscoveryError) ListHosts() ([]Node, error) {
	e := errors.New("An error")

	return nil, e
}
func (c *MockDiscoveryError) WatchHosts() {}
