package swarm

import "github.com/sayden/docker-commander/config"

//TYPE_ENV must be used in dev and production (no mock)
const TYPE_ENV = 0

//TYPE_MOCK_OK is used in testing to return a mock correct client
const TYPE_MOCK_OK int = 1

//TYPE_MOCK_ERROR is used in testing to return a mock incorrect client
const TYPE_MOCK_ERROR int = 2

// GetClient decouples Swarm Client creation and creates one based on
// config file
func GetClient(cType int) Swarm {
	switch cType {
	case TYPE_ENV:
		switch config.CURRENT_ENV {
		case config.DEVELOPMENT:
			s := HTTPClient{config.SWARM_MANAGER}
			return &s
		case config.PRODUCTION:
			// Production mode in swarm factory
			s := HTTPClient{"http://ip_not_set"}
			return &s
		default:
			s := HTTPClient{config.SWARM_MANAGER}
			return &s
		}
	case TYPE_MOCK_OK:
		return &HTTPClientMock{"http://a_host"}
	case TYPE_MOCK_ERROR:
		return &HTTPClientMockError{"http://a_host"}
	default:
		return &HTTPClientMockError{"http://a_host"}
	}
}

// GetClientWithIP returns a configured client with the specified ip
func GetClientWithIP(ip string) Swarm {
	if config.CURRENT_ENV == config.DEVELOPMENT {
		return &HTTPClientMock{"http://a_host"}
	}
	return &HTTPClient{Host: ip}
}
