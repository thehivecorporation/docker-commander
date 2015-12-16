package discovery

import (
	"log"

	"github.com/sayden/docker-commander/config"
)

// GetClient returns a configured client to talk with discovery service
func GetClient() InfoService {
	switch config.DISCOVERY_SERVICE {
	case config.ETCD:
		s := EtcdClient{Host: config.ETCD_HOST}
		return &s
	case config.CONSUL:
		log.Fatal("Consul not yet implemented")
		return nil
	case config.ZOOKEEPER:
		log.Fatal("Zookeeper not yet implemented")
		return nil
	default:
		s := EtcdClient{Host: config.ETCD_HOST}
		return &s
	}
}
