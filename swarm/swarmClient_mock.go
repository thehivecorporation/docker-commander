package swarm

import (
	"encoding/json"

	"github.com/sayden/docker-commander/Godeps/_workspace/src/github.com/samalba/dockerclient"
)

// HTTPClientMock is a struct to mock http.Client
type HTTPClientMock struct {
	Host string
}

// ListInfo is the mocked Get implementation of http.Get
func (h *HTTPClientMock) ListInfo() (dockerclient.Info, error) {
	str := `{"ID":"","Containers":102,"Driver":"","DriverStatus":[["\bRole","primary"],["\bStrategy","spread"],["\bFilters","health, port, dependency, affinity, constraint"],["\bNodes","1"],["osboxes","192.168.1.39:2375"],[" â”” Status","Healthy"],[" â”” Containers","102"],[" â”” Reserved CPUs","0 / 1"],[" â”” Reserved Memory","0 B / 1.888 GiB"],[" â”” Labels","executiondriver=native-0.2, kernelversion=3.10.0-123.el7.x86_64, operatingsystem=CentOS Linux 7 (Core), storagedriver=devicemapper"]],"ExecutionDriver":"","Images":5,"KernelVersion":"","OperatingSystem":"","NCPU":1,"MemTotal":2027035852,"Name":"6514d918bc01","Labels":null,"Debug":false,"NFd":0,"NGoroutines":0,"SystemTime":"2015-12-13T17:25:40.540782247Z","NEventsListener":0,"InitPath":"","InitSha1":"","IndexServerAddress":"","MemoryLimit":true,"SwapLimit":true,"IPv4Forwarding":true,"BridgeNfIptables":true,"BridgeNfIp6tables":true,"DockerRootDir":"","HttpProxy":"","HttpsProxy":"","NoProxy":""}`
	byt := []byte(str)

	jsonCs := dockerclient.Info{}
	err := json.Unmarshal(byt, &jsonCs)

	if err != nil {
		return jsonCs, err
	}

	return jsonCs, nil
}

//ListContainers returns a byte array with containers data
func (h *HTTPClientMock) ListContainers() ([]dockerclient.Container, error) {
	str := `[{"Id":"6514d918bc01b0df8098053086c72cf9e4e3f0b166987b5b5db78d1f6b783ea8","Names":["/swarm-agent"],"Image":"swarm","Command":"/swarm manage etcd://192.168.1.35:2379","Created":1450018068,"Ports":[{"IP":"0.0.0.0","PrivatePort":2375,"PublicPort":8081,"Type":"tcp"}],"Labels":null,"Status":"Up 2 hours","HostConfig":{"NetworkMode":"default"}},{"Id":"a149e76d6a95f45c5ee546886fcb43e68ae5323121f84f2302007f3b28aa45ab","Names":["/swarm"],"Image":"swarm","Command":"/swarm join --addr=192.168.1.35:2375 etcd://192.168.1.35:2379","Created":1450017970,"Ports":[{"PrivatePort":2375,"Type":"tcp"}],"Labels":null,"Status":"Up 2 hours","HostConfig":{"NetworkMode":"default"}}]`
	byt := []byte(str)

	jsonCs := []dockerclient.Container{}
	err := json.Unmarshal(byt, &jsonCs)

	if err != nil {
		return jsonCs, err
	}

	return jsonCs, nil
}

//ListImages returns a byte array with images data
func (h *HTTPClientMock) ListImages() ([]dockerclient.Image, error) {
	str := `[{"Id":"e9ff33e7e5b9a683ab735bbe99450c50bd0b64c4e414d12c94ff93b345d3bb18","ParentId":"22218d75fdd7dfba85feb60f8165eecb5ba13b9364b07ec2b8cd76e887bd6d78","RepoTags":["docker.io/swarm:latest"],"RepoDigests":[],"Created":1449709381,"Size":0,"VirtualSize":17146714,"Labels":null},{"Id":"088f61431e99e47b795a8e0026753bb701b607a6bc356123ac3e058b71376753","ParentId":"8ece2cf6713e1370451fbefbeb5e8985cc27c11ab1b21f188c8a04ac6db92aa9","RepoTags":["sayden/centos-simplest-node:latest"],"RepoDigests":[],"Created":1449525441,"Size":0,"VirtualSize":544843301,"Labels":{"License":"GPLv2","Vendor":"CentOS"}}]`
	byt := []byte(str)

	jsonCs := []dockerclient.Image{}
	err := json.Unmarshal(byt, &jsonCs)

	if err != nil {
		return jsonCs, err
	}

	return jsonCs, nil
}

// HTTPClientMockError is a struct to mock http.Client
type HTTPClientMockError struct {
	Host string
}

// ListInfo is the mocked Get implementation of http.Get with errors
func (h *HTTPClientMockError) ListInfo() (dockerclient.Info, error) {
	str := `{"ID":"","Containers":102,"Driver":"","DriverStatus":`
	byt := []byte(str)

	jsonCs := dockerclient.Info{}
	err := json.Unmarshal(byt, &jsonCs)

	if err != nil {
		return jsonCs, err
	}

	return jsonCs, nil
}

//ListContainers returns a byte array with containers data with errors
func (h *HTTPClientMockError) ListContainers() ([]dockerclient.Container, error) {
	str := `[{"Id":"6514d918bc01b0df8098053086c72cf9e4e3f0b166987b5b5db78d1f6b783ea8`
	byt := []byte(str)

	jsonCs := []dockerclient.Container{}
	err := json.Unmarshal(byt, &jsonCs)

	if err != nil {
		return jsonCs, err
	}

	return jsonCs, nil
}

//ListImages returns a byte array with images data with errors
func (h *HTTPClientMockError) ListImages() ([]dockerclient.Image, error) {
	str := `[{"Id":"e9ff33e7e5b9a683ab735bbe99450c50bd0b64c4e414d12c94ff93b345d3bb18`
	byt := []byte(str)

	jsonCs := []dockerclient.Image{}
	err := json.Unmarshal(byt, &jsonCs)

	if err != nil {
		return jsonCs, err
	}

	return jsonCs, nil
}
