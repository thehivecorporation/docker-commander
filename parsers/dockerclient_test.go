package parsers

import "testing"

func TestParseContainer(t *testing.T) {
	str := `[{"Id":"6514d918bc01b0df8098053086c72cf9e4e3f0b166987b5b5db78d1f6b783ea8","Names":["/swarm-agent"],"Image":"swarm","Command":"/swarm manage etcd://192.168.1.35:2379","Created":1450018068,"Ports":[{"IP":"0.0.0.0","PrivatePort":2375,"PublicPort":8081,"Type":"tcp"}],"Labels":null,"Status":"Up 2 hours","HostConfig":{"NetworkMode":"default"}},{"Id":"a149e76d6a95f45c5ee546886fcb43e68ae5323121f84f2302007f3b28aa45ab","Names":["/swarm"],"Image":"swarm","Command":"/swarm join --addr=192.168.1.35:2375 etcd://192.168.1.35:2379","Created":1450017970,"Ports":[{"PrivatePort":2375,"Type":"tcp"}],"Labels":null,"Status":"Up 2 hours","HostConfig":{"NetworkMode":"default"}}]`
	byt := []byte(str)

	dp := DockerClientParser{}
	containers, err := dp.parseContainer(&byt)

	if err != nil {
		t.Fatal("Container parse incorrect")
	}

	if containers[0].Id != "6514d918bc01b0df8098053086c72cf9e4e3f0b166987b5b5db78d1f6b783ea8" {
		t.Fatal("Container parse failed")
	}

	//Pass an string that its going to produce an error
	anErr := []byte("{\"asfasdf\":1")
	containers, err = dp.parseContainer(&anErr)

	if err == nil {
		t.Fatal("Container parse must be incorrect")
	}

	if containers != nil {
		t.Fatal("Container res should be nil")
	}
}

func TestParseImages(t *testing.T) {
	str := `[{"Id":"e9ff33e7e5b9a683ab735bbe99450c50bd0b64c4e414d12c94ff93b345d3bb18","ParentId":"22218d75fdd7dfba85feb60f8165eecb5ba13b9364b07ec2b8cd76e887bd6d78","RepoTags":["docker.io/swarm:latest"],"RepoDigests":[],"Created":1449709381,"Size":0,"VirtualSize":17146714,"Labels":null},{"Id":"088f61431e99e47b795a8e0026753bb701b607a6bc356123ac3e058b71376753","ParentId":"8ece2cf6713e1370451fbefbeb5e8985cc27c11ab1b21f188c8a04ac6db92aa9","RepoTags":["sayden/centos-simplest-node:latest"],"RepoDigests":[],"Created":1449525441,"Size":0,"VirtualSize":544843301,"Labels":{"License":"GPLv2","Vendor":"CentOS"}}]`
	byt := []byte(str)

	dp := DockerClientParser{}
	images, err := dp.parseImages(&byt)

	if err != nil {
		t.Fatal("Image Parse incorrect")
	}

	if images[0].Id != "e9ff33e7e5b9a683ab735bbe99450c50bd0b64c4e414d12c94ff93b345d3bb18" {
		t.Fatal("Image parse failed")
	}

	//Pass an string that its going to produce an error
	anErr := []byte("{\"asfasdf\":1")
	images, err = dp.parseImages(&anErr)

	if err == nil {
		t.Fatal("Container parse must be incorrect")
	}

	if images != nil {
		t.Fatal("Container res should be nil")
	}
}

func TestParseInfo(t *testing.T) {
	str := `{"ID":"","Containers":102,"Driver":"","DriverStatus":[["\bRole","primary"],["\bStrategy","spread"],["\bFilters","health, port, dependency, affinity, constraint"],["\bNodes","1"],["osboxes","192.168.1.39:2375"],[" â”” Status","Healthy"],[" â”” Containers","102"],[" â”” Reserved CPUs","0 / 1"],[" â”” Reserved Memory","0 B / 1.888 GiB"],[" â”” Labels","executiondriver=native-0.2, kernelversion=3.10.0-123.el7.x86_64, operatingsystem=CentOS Linux 7 (Core), storagedriver=devicemapper"]],"ExecutionDriver":"","Images":5,"KernelVersion":"","OperatingSystem":"","NCPU":1,"MemTotal":2027035852,"Name":"6514d918bc01","Labels":null,"Debug":false,"NFd":0,"NGoroutines":0,"SystemTime":"2015-12-13T17:25:40.540782247Z","NEventsListener":0,"InitPath":"","InitSha1":"","IndexServerAddress":"","MemoryLimit":true,"SwapLimit":true,"IPv4Forwarding":true,"BridgeNfIptables":true,"BridgeNfIp6tables":true,"DockerRootDir":"","HttpProxy":"","HttpsProxy":"","NoProxy":""}`
	byt := []byte(str)

	dp := DockerClientParser{}
	info, err := dp.parseInfo(&byt)

	if err != nil {
		t.Fatal("Info Parse incorrect")
	}

	if info.Containers != 102 {
		t.Fatal("Info parse failed")
	}

	//Pass an string that its going to produce an error
	anErr := []byte("{\"asfasdf\":1")
	info, err = dp.parseInfo(&anErr)

	if err == nil {
		t.Fatal("Info parse must be incorrect")
	}
}
