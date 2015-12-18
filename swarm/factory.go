package swarm

import "github.com/sayden/docker-commander/config"

// GetClient decouples Swarm Client creation and creates one based on
// config file
func GetClient() Swarm {
	if config.CURRENT_ENV == config.DEVELOPMENT {
		return &HTTPClient{config.SWARM_MANAGER_DEVELOPMENT}

	} else if config.CURRENT_ENV == config.PRODUCTION {
		return &HTTPClient{config.SWARM_MANAGER_PRODUCTION}

	} else if config.CURRENT_ENV == config.MOCK {
		return &HTTPClientMock{"http://a_host"}
	}

	return &HTTPClientMock{"http://a_host"}
}

// GetClientWithIP returns a configured client with the specified ip
func GetClientWithIP(ip string) Swarm {
	if config.CURRENT_ENV == config.DEVELOPMENT {
		return &HTTPClient{ip}

	} else if config.CURRENT_ENV == config.PRODUCTION {
		return &HTTPClient{ip}

	} else if config.CURRENT_ENV == config.MOCK {
		return &HTTPClientMock{"http://a_host"}
	}

	return &HTTPClientMock{"http://a_host"}
}
