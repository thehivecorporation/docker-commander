package discovery

import (
	"log"
	"time"

	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

// EtcdClient struct to use with ETCD clusters
type EtcdClient struct{}

// ListHosts returns connected swarm nodes
func (e *EtcdClient) ListHosts(serviceURLs string) ([]Node, error) {
	//etcd implementation here
	cfg := client.Config{
		Endpoints: []string{serviceURLs},
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
		nodesArray[i] = Node{n.Value}
	}

	return nodesArray, nil
}
