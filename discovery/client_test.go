package discovery

import (
	"errors"
	"testing"
)

type mockDiscoveryOk struct {
	Host string
}

type mockDiscoveryError struct {
	Host string
}

// ListHosts mocked
func (c *mockDiscoveryOk) ListHosts() ([]Node, error) {
	ns := []Node{Node{IP: "ip1"}, Node{IP: "ip2"}}

	return ns, nil
}
func (c *mockDiscoveryOk) WatchHosts() {}

//ListHosts Error
func (c *mockDiscoveryError) ListHosts() ([]Node, error) {
	e := errors.New("An error")

	return nil, e
}
func (c *mockDiscoveryError) WatchHosts() {}

func TestClientMockOk(t *testing.T) {
	i := mockDiscoveryOk{Host: "mockIp"}
	r, err := ListHosts(&i)

	if err != nil {
		t.Fail()
	}

	for _, n := range r {
		t.Log(n.IP)
	}
}

func TestClientMockError(t *testing.T) {
	i := mockDiscoveryError{Host: "mockIp"}
	_, err := ListHosts(&i)

	if err == nil {
		t.Fail()
	}
}
