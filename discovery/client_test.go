package discovery

import "testing"

type mockDiscovery struct{}

// ListHosts mocked
func (c *mockDiscovery) ListHosts(serviceURL string) ([]Node, error) {
	ns := []Node{Node{"ip1"}, Node{"ip2"}}

	return ns, nil
}

func TestClientMock(t *testing.T) {
	i := mockDiscovery{}
	c := Client{}
	r, err := c.ListHosts("mockIp", &i)

	if err != nil {
		t.Fail()
	}

	for _, n := range r {
		t.Log(n.IP)
	}
}
