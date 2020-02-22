package authentication_service

import (
	protocolAuth "github.com/leshachaplin/grpc-server/protocol"
)

type Auth struct {
	client protocolAuth.AuthServiceClient
}

func NewHandlerAuth(client protocolAuth.AuthServiceClient) *Auth {
	return &Auth{client: client}
}
