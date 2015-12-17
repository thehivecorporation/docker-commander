package discovery

import (
	"log"

	"github.com/sayden/docker-commander/config"
)

//TYPE_ENV must be used in dev and production (no mock)
const TYPE_ENV = 0

//TYPE_MOCK_OK is used in testing to return a mock correct client
const TYPE_MOCK_OK int = 1

//TYPE_MOCK_ERROR is used in testing to return a mock incorrect client
const TYPE_MOCK_ERROR int = 2

// GetClient returns a configured client to talk with discovery service
func GetClient(cType int) InfoService {
	switch cType {
	case TYPE_ENV:
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
	case TYPE_MOCK_OK:
		return &MockDiscoveryOk{"http://a_host"}
	case TYPE_MOCK_ERROR:
		return &MockDiscoveryError{"http://a_host"}
	default:
		return &MockDiscoveryError{"http://a_host"}
	}

}
