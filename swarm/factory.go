package swarm

import (
	"log"

	"github.com/sayden/docker-commander/config"
)

// GetClient decouples Swarm Client creation and creates one based on
// config file
func GetClient() Swarm {
	switch config.CURRENT_ENV {
	case config.DEVELOPMENT:
		s := HTTPClient{config.SWARM_MANAGER}
		return &s
	case config.PRODUCTION:
		//TODO Production mode in swarm factory
		log.Fatal("Production mode not configured yet for swarm factory")
		return nil
	default:
		s := HTTPClient{config.SWARM_MANAGER}
		return &s
	}
}
