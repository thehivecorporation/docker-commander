package parsers

import (
	"encoding/json"

	"github.com/samalba/dockerclient"
)

//ContainerParser is used to parse an array of bytes to Container types
type ContainerParser interface {
	parseContainer(cs *[]byte) ([]dockerclient.Container, error)
}

//ImageParser is used to parse an array of bytes to Images types
type ImageParser interface {
	parseImages(is *[]byte) ([]dockerclient.Image, error)
}

//InfoParser is used to parse an array of bytes to Cluster types
type InfoParser interface {
	parseInfo(i *[]byte) (dockerclient.Info, error)
}

// DockerClientParser uses dockerclient project to parse objects back
type DockerClientParser struct{}

func (d *DockerClientParser) parseContainer(cs *[]byte) ([]dockerclient.Container, error) {
	jsonCs := []dockerclient.Container{}
	err := json.Unmarshal(*cs, &jsonCs)
	if err != nil {
		return nil, err
	}

	return jsonCs, nil
}

func (d *DockerClientParser) parseImages(is *[]byte) ([]dockerclient.Image, error) {
	jsonCs := []dockerclient.Image{}
	err := json.Unmarshal(*is, &jsonCs)

	if err != nil {
		return nil, err
	}

	return jsonCs, nil
}

func (d *DockerClientParser) parseInfo(i *[]byte) (dockerclient.Info, error) {
	jsonCs := dockerclient.Info{}
	err := json.Unmarshal(*i, &jsonCs)

	if err != nil {
		return jsonCs, err
	}

	return jsonCs, nil
}
