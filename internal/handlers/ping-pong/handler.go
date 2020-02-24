package ping_pong

import (
	protocolPP "github.com/ivansukach/grpc-server/protocol"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type PingPong struct {
	client protocolPP.GetResponseClient
}

func NewHandlerPP(url string, version string, e *echo.Echo, opts grpc.DialOption) *PingPong {
	clientConnPingPong, err := grpc.Dial(url, opts)
	if err != nil {
		log.Error(err)
	}

	defer clientConnPingPong.Close()

	clientPP := protocolPP.NewGetResponseClient(clientConnPingPong)
	pp := &PingPong{client: clientPP}
	v := "v" + version
	e.GET(v+"games/pingPong", pp.PingPong)
	return pp
}
