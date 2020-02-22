package profile_service

import (
	protocolPS "github.com/ivansukach/profile-service/protocol"
)

type ProfileAccounting struct {
	client protocolPS.ProfileServiceClient
}

func NewHandlerPS(client protocolPS.ProfileServiceClient) *ProfileAccounting {
	return &ProfileAccounting{client: client}
}
