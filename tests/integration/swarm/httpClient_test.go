package restClient

import (
	"encoding/json"
	"testing"

	"github.com/sayden/docker-commander/swarm"
)

func TestGetHosts(t *testing.T) {
	h := swarm.HTTPClient{Host: "http://192.168.1.35:8081"}
	r, err := swarm.GetHosts(&h)

	if err != nil {
		t.Fail()
	}

	var dat map[string]interface{}

	if err := json.Unmarshal(r, &dat); err != nil {
		t.Fatal(err)
	}

	if dat["DriverStatus"] == nil {
		t.Fail()
	}
}

func TestGetContainers(t *testing.T) {
	h := swarm.HTTPClient{Host: "http://192.168.1.35:2375"}
	r, err := swarm.GetContainers(&h)

	if err != nil {
		t.Fatal("Error trying to get containers")
	}

	var dat []map[string]interface{}

	if err := json.Unmarshal(r, &dat); err != nil {
		t.Fatal(err)
	}

	if len(dat) == 0 {
		t.Log("No containers found")
	} else {
		c := dat[0]
		if c["Id"] == nil {
			t.Fail()
		}

		if (c["Image"]) == nil {
			t.Fail()
		}
	}
}

func TestGetImages(t *testing.T) {
	h := swarm.HTTPClient{Host: "http://192.168.1.35:2375"}
	r, err := swarm.GetImages(&h)

	if err != nil {
		t.Fatal("Error trying to get containers")
	}

	var dat []map[string]interface{}

	if err := json.Unmarshal(r, &dat); err != nil {
		t.Fatal(err)
	}

	if len(dat) == 0 {
		t.Log("No Images found")
	} else {
		c := dat[0]
		if c["Id"] == nil {
			t.Fail()
		}

		if (c["Repo"]) == nil {
			t.Fail()
		}
	}
}
