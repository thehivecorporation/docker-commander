package restClient

import (
	"testing"

	"github.com/sayden/docker-commander/config"
	"github.com/sayden/docker-commander/swarm"
)

func TestGetHosts(t *testing.T) {
	h := swarm.HTTPClient{Host: config.SWARM_MANAGER_DEVELOPMENT}
	r, err := h.ListInfo()

	if err != nil {
		t.Fail()
	}

	if r.DriverStatus == nil {
		t.Fail()
	}
}

func TestGetContainers(t *testing.T) {
	h := swarm.HTTPClient{Host: config.SWARM_MANAGER_DEVELOPMENT}
	r, err := h.ListContainers()

	if err != nil {
		t.Fatal("Error trying to get containers")
	}

	if len(r) == 0 {
		t.Log("No containers found")
	} else {
		c := r[0]
		if c.Id == "" {
			t.Fail()
		}

		if c.Image == "" {
			t.Fail()
		}
	}
}

func TestGetImages(t *testing.T) {
	h := swarm.HTTPClient{Host: config.SWARM_MANAGER_DEVELOPMENT}
	r, err := h.ListImages()

	if err != nil {
		t.Fatal("Error trying to get containers")
	}

	if len(r) == 0 {
		t.Log("No Images found")
	} else {
		c := r[0]
		if c.Id == "" {
			t.Fatal("'Id' not found")
		}

		if len(c.RepoTags) == 0 {
			t.Fatal("'RepoTags' size is zero")
		}
	}
}
