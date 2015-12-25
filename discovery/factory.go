package discovery

import (
	"github.com/sayden/docker-commander/discovery/Godeps/_workspace/src/github.com/sayden/docker-commander/config"
)

// GetClient returns a configured client to talk with discovery service
func GetClient() InfoService {
	if config.CURRENT_ENV == config.DEVELOPMENT {
		switch config.DISCOVERY_SERVICE {
		case config.ETCD:
			s := EtcdClient{Host: config.ETCD_HOST_DEVELOPMENT}
			return &s
		case config.CONSUL:
			log.Fatal("Consul not yet implemented")
			return nil
		case config.ZOOKEEPER:
			log.Fatal("Zookeeper not yet implemented")
			return nil
		default:
			s := EtcdClient{Host: config.ETCD_HOST_DEVELOPMENT}
			return &s
		}

	} else if config.CURRENT_ENV == config.PRODUCTION {
		switch config.DISCOVERY_SERVICE {
		case config.ETCD:
			s := EtcdClient{Host: config.ETCD_HOST_PRODUCTION}
			return &s
		case config.CONSUL:
			log.Fatal("Not yet implemented")
			return nil
		case config.ZOOKEEPER:
			log.Fatal("Not yet implemented")
			return nil
		default:
			s := EtcdClient{Host: config.ETCD_HOST_PRODUCTION}
			return &s
		}

	} else {
		return &MockDiscoveryOk{"http://a_host"}
	}
}
