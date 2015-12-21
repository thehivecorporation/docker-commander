package receivers

import (
	"github.com/sayden/docker-commander/communications"
	"github.com/sayden/docker-commander/config"
	"github.com/sayden/docker-commander/discovery"
	"github.com/sayden/docker-commander/logger"
	"github.com/sayden/docker-commander/socket"
	"github.com/sayden/docker-commander/swarm"
)

var log = logger.WithField("socker:handler")

type ReceiverPayload struct {
	s *swarm.Swarm
	i *discovery.InfoService
	msgr communications.FrontMessenger
}

//Cluster respond to socket message "cluster" to return the entire cluster
//information
func Cluster(s *swarm.Swarm, i *discovery.InfoService, msgr communications.FrontMessenger) {
	info, err := socket.GetFullInfo(*s, *i)
	if err != nil {
		log.Error("Error trying to get cluster info", err)
	} else {
		sr := communications.ConnectionResponse{
			Response: &info,
			Action:   config.CONNECTION_ACTION_CLUSTER,
		}
		msgr.FrontMessage(&sr)
	}
}
