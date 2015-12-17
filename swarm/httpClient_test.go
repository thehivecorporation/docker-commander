package swarm

import (
	"encoding/json"
	"testing"
)

func TestIsErrorInURL(t *testing.T) {
	err := isErrorInURL("asdfasdf")
	if err == nil {
		t.Fatal("Host is incorrect but error is nil. Err must contain an error")
	}
}

func TestGetHosts(t *testing.T) {
	mock := HTTPClientMock{Host: "http://some_url"}
	c, err := mock.ListInfo()

	if err != nil {
		t.Fatal(err)
	}

	var dat map[string]interface{}

	if err := json.Unmarshal(c, &dat); err != nil {
		t.Fatal(err)
	}

	if dat["ID"] == nil {
		t.Fail()
	}

	if dat["Containers"] == nil {
		t.Fail()
	}

}

func TestGetContainers(t *testing.T) {
	mock := HTTPClientMock{Host: "http://some_url"}
	c, err := mock.ListContainers()

	if err != nil {
		t.Fatal(err)
	}

	var dat []map[string]interface{}
	if err := json.Unmarshal(c, &dat); err != nil {
		t.Fatal(err)
	}

	if len(dat) == 0 {
		t.Log("No containers found")
	} else {
		c := dat[0]

		if c["Id"] == nil {
			t.Fail()
		}

		if c["Image"] == nil {
			t.Fail()
		}
	}

}

func TestGetImages(t *testing.T) {
	mock := HTTPClientMock{Host: "http://some_url"}
	c, err := mock.ListImages()

	if err != nil {
		t.Fatal(err)
	}

	var dat []map[string]interface{}
	if err := json.Unmarshal(c, &dat); err != nil {
		t.Fatal(err)
	}

	if len(dat) == 0 {
		t.Log("No Images found")
	} else {
		c := dat[0]

		if c["Id"] == nil {
			t.Fatal("'Id' field not found")
		}

		if c["RepoTags"] == nil {
			t.Fatal("'RepoTags' field not found")
		}
	}
}
