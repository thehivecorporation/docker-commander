package swarm

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// HTTPClient implementation to access swarm info
type HTTPClient struct {
	Host string
}

// GetInfo Rest implementation
func (c *HTTPClient) GetInfo() ([]byte, error) {
	return makeHTTPGetRequest(c, "/info")
}

// GetContainers returns the Containers of a specific host. . Replicates
// GET [docker-host]:2375/containers/json
func (c *HTTPClient) GetContainers() ([]byte, error) {
	return makeHTTPGetRequest(c, "/containers/json")
}

// GetImages returns the Images of a specific host. Replicates
// GET [docker-host]:2375/images/json
func (c *HTTPClient) GetImages() ([]byte, error) {
	return makeHTTPGetRequest(c, "/images/json")
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

func makeHTTPGetRequest(c *HTTPClient, trailURL string) ([]byte, error) {
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
