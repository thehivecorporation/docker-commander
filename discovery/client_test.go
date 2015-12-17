package discovery

import "testing"

func TestClientMockOk(t *testing.T) {
	i := MockDiscoveryOk{Host: "mockIp"}
	r, err := ListHosts(&i)

	if err != nil {
		t.Fail()
	}

	for _, n := range r {
		t.Log(n.IP)
	}
}

func TestClientMockError(t *testing.T) {
	i := MockDiscoveryError{Host: "mockIp"}
	_, err := ListHosts(&i)

	if err == nil {
		t.Fail()
	}
}
