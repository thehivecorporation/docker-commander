package discovery

import (
	"log"

	"github.com/sayden/docker-commander/config"
)

// GetClient returns a configured client to talk with discovery service
func GetClient() InfoService {
	switch config.CURRENT_ENV {
	case config.DEVELOPMENT:
		s := EtcdClient{Host: config.ETCD}
		return &s
	case config.PRODUCTION:
		//TODO Production mode in swarm factory
		log.Fatal("Production mode not configured yet for swarm factory")
		return nil
	default:
		s := EtcdClient{Host: config.ETCD}
		return &s
	}
}
