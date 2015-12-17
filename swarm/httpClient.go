package swarm

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/sayden/docker-commander/Godeps/_workspace/src/github.com/samalba/dockerclient"
)

// HTTPClient implementation to access swarm info
type HTTPClient struct {
	Host string
}

// ListInfo Rest implementation
func (c *HTTPClient) ListInfo() (dockerclient.Info, error) {
	i, err := c.makeHTTPGetRequest("/info")
	if err != nil {
		return dockerclient.Info{}, err
	}
	jsonCs := dockerclient.Info{}
	err = json.Unmarshal(i, &jsonCs)

	if err != nil {
		return jsonCs, err
	}

	return jsonCs, nil
}

// ListContainers returns the Containers of a specific host. . Replicates
// GET [docker-host]:2375/containers/json
func (c *HTTPClient) ListContainers() ([]dockerclient.Container, error) {
	i, err := c.makeHTTPGetRequest("/containers/json")
	if err != nil {
		return []dockerclient.Container{}, err
	}
	jsonCs := []dockerclient.Container{}
	err = json.Unmarshal(i, &jsonCs)

	if err != nil {
		return jsonCs, err
	}

	return jsonCs, nil
}

// ListImages returns the Images of a specific host. Replicates
// GET [docker-host]:2375/images/json
func (c *HTTPClient) ListImages() ([]dockerclient.Image, error) {
	i, err := c.makeHTTPGetRequest("/images/json")
	if err != nil {
		return []dockerclient.Image{}, err
	}
	jsonCs := []dockerclient.Image{}
	err = json.Unmarshal(i, &jsonCs)

	if err != nil {
		return jsonCs, err
	}

	return jsonCs, nil
}

//TODO The URL must have HTTP en port included
func isErrorInURL(url string) error {
	//Check url string has the 'http://' included
	if strings.Contains(url, "http://") != true {
		return errors.New("Url is not recognized as valid URL, maybe you" +
			"forgot to add 'http://'")
	}

	return nil
}

func (c *HTTPClient) makeHTTPGetRequest(trailURL string) ([]byte, error) {
	if err := isErrorInURL(c.Host); err != nil {
		return nil, err
	}

	url := c.Host + trailURL
	log.Println("Requesting to", url)
	r, err := http.Get(c.Host + trailURL)
	defer r.Body.Close()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	byt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return byt, nil
}
