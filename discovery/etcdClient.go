package discovery

import (
	"fmt"
	"log"
	"time"

	"github.com/sayden/docker-commander/Godeps/_workspace/src/github.com/coreos/etcd/client"
	"github.com/sayden/docker-commander/Godeps/_workspace/src/golang.org/x/net/context"
)

// EtcdClient struct to use with ETCD clusters
// TODO Make it work with an ETCD cluster instead of only one node
type EtcdClient struct {
	Host string
}

// WatchHosts TODO use some goroutines here
func (e *EtcdClient) WatchHosts() {
	kapi, err := e.getClient()

	if err != nil {
		panic("")
	}

	o := client.WatcherOptions{Recursive: true}
	w := kapi.Watcher("/test", &o)

	r, err := w.Next(context.Background())

	//TODO DELETE
	opGet := client.GetOptions{Recursive: true}
	respNodes, err := kapi.Get(context.Background(), "/test", &opGet)
	ns := respNodes.Node.Nodes

	// TODO Make a break action to stop recursivity
	if err == nil {
		if len(ns) == 0 {
			fmt.Println("Exiting app")
		} else {
			log.Printf("%q action triggered on %q to set it to %q\n", r.Action, r.Node.Key, r.Node.Value)
			log.Println(len(ns))
			e.WatchHosts()
		}
	} else {
		log.Fatal(err)
	}
}

// ListHosts returns connected swarm nodes
func (e *EtcdClient) ListHosts() ([]Node, error) {
	kapi, err := e.getClient()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// get "/docker/swarm/nodes" list
	o := client.GetOptions{Recursive: true}
	resp, err := kapi.Get(context.Background(), "/docker/swarm/nodes", &o)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if resp.Node.Dir == false {
		log.Fatal("Node not recognized as directory")
	}

	nodesArray := make([]Node, len(resp.Node.Nodes))
	for i, n := range resp.Node.Nodes {
		nodesArray[i] = Node{IP: n.Value}
	}

	return nodesArray, nil
}

func (e *EtcdClient) getClient() (client.KeysAPI, error) {
	cfg := client.Config{
		Endpoints: []string{e.Host},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	kapi := client.NewKeysAPI(c)
	return kapi, nil
}
