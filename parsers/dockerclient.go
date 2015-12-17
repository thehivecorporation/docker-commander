package parsers

import (
	"encoding/json"

	"github.com/samalba/dockerclient"
)

// DockerClientParser uses dockerclient project to parse objects back
type DockerClientParser struct{}

//ContainerParser is used to parse an array of bytes to Container types
type ContainerParser interface {
	ParseContainer(cs *[]byte) ([]dockerclient.Container, error)
}

//ImageParser is used to parse an array of bytes to Images types
type ImageParser interface {
	ParseImages(is *[]byte) ([]dockerclient.Image, error)
}

//InfoParser is used to parse an array of bytes to Cluster types
type InfoParser interface {
	ParseInfo(i *[]byte) (dockerclient.Info, error)
}

// ParseContainer parses a []byte to dockerclient.Container objects
func (d *DockerClientParser) ParseContainer(cs *[]byte) ([]dockerclient.Container, error) {
	jsonCs := []dockerclient.Container{}
	err := json.Unmarshal(*cs, &jsonCs)
	if err != nil {
		return nil, err
	}

	return jsonCs, nil
}

// ParseImages parses a []byte to dockerclient.Image objects
func (d *DockerClientParser) ParseImages(is *[]byte) ([]dockerclient.Image, error) {
	jsonCs := []dockerclient.Image{}
	err := json.Unmarshal(*is, &jsonCs)

	if err != nil {
		return nil, err
	}

	return jsonCs, nil
}

// ParseInfo parses a []byte to dockerclient.Info objects
func (d *DockerClientParser) ParseInfo(i *[]byte) (dockerclient.Info, error) {
	jsonCs := dockerclient.Info{}
	err := json.Unmarshal(*i, &jsonCs)

	if err != nil {
		return jsonCs, err
	}

	return jsonCs, nil
}
