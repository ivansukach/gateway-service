package services

import (
	"github.com/ivansukach/gateway-service/internal/repositories"
	log "github.com/sirupsen/logrus"
)

func UpdateEndpoints(rps repositories.Repository) ([]*repositories.Endpoints, error) {
	endpoints, err := rps.GetNewGateways()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	newEndpoints := make([]*repositories.Endpoints, 0)
	for i := range endpoints {
		latestId := rps.GetLatestId()
		if latestId < endpoints[i].Id {
			newEndpoints = append(newEndpoints, endpoints[i])
			latestId = endpoints[i].Id
		}
		rps.SetLatestId(latestId)
	}
	return newEndpoints, nil
}
