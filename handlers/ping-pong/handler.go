package ping_pong

import (
	protocolPP "github.com/ivansukach/grpc-server/protocol"
)

type PingPong struct {
	client protocolPP.GetResponseClient
}

func NewHandlerPP(client protocolPP.GetResponseClient) *PingPong {
	return &PingPong{client: client}
}
