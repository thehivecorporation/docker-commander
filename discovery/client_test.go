package discovery

import "testing"

type mockDiscovery struct {
	Host string
}

// ListHosts mocked
func (c *mockDiscovery) ListHosts() ([]Node, error) {
	ns := []Node{Node{"ip1"}, Node{"ip2"}}

	return ns, nil
}

func (c *mockDiscovery) WatchHosts() {

}

func TestClientMock(t *testing.T) {
	i := mockDiscovery{Host: "mockIp"}
	r, err := ListHosts(&i)

	if err != nil {
		t.Fail()
	}

	for _, n := range r {
		t.Log(n.IP)
	}
}
