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
	S      *swarm.Swarm
	I      *discovery.InfoService
	Msgr   communications.FrontMessenger
	Action string
	Payload string
}

//Cluster respond to socket message "cluster" to return the entire cluster
//information
func Cluster(rp *ReceiverPayload) {
	info, err := socket.GetFullInfo(*rp.S, *rp.I)
	if err != nil {
		log.Error("Error trying to get cluster info", err)
	} else {
		sr := communications.ConnectionResponse{
			Response: &info,
			Action:   config.CONNECTION_ACTION_CLUSTER,
		}
		rp.Msgr.FrontMessage(&sr)
	}
}

func SetSwarmIP(rp *ReceiverPayload) {
	if config.CURRENT_ENV == config.DEVELOPMENT {
		config.SWARM_MANAGER_DEVELOPMENT = rp.Payload
	} else {
		config.SWARM_MANAGER_PRODUCTION = rp.Payload
	}
}
