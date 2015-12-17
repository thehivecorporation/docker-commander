package swarm

import "testing"

func TestIsErrorInURL(t *testing.T) {
	err := isErrorInURL("asdfasdf")
	if err == nil {
		t.Fatal("Host is incorrect but error is nil. Err must contain an error")
	}
}

func TestListInfo(t *testing.T) {
	mock := HTTPClientMock{Host: "http://some_url"}
	_, err := mock.ListInfo()

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetContainers(t *testing.T) {
	mock := HTTPClientMock{Host: "http://some_url"}
	dat, err := mock.ListContainers()

	if err != nil {
		t.Fatal(err)
	}

	if len(dat) == 0 {
		t.Log("No containers found")
	}

}

func TestGetImages(t *testing.T) {
	mock := HTTPClientMock{Host: "http://some_url"}
	dat, err := mock.ListImages()

	if err != nil {
		t.Fatal(err)
	}

	if len(dat) == 0 {
		t.Log("No Images found")
	}
}
