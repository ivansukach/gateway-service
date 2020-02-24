package ping_pong

import (
	"context"
	"github.com/ivansukach/grpc-server/protocol"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func (pp *PingPong) PingPong(c echo.Context) error {
	log.Info("PingPong")
	request := &protocol.GRRequest{Req: "Ping"}
	response, err := pp.client.GiveResponse(context.Background(), request)
	log.Println("Response:", response.GetRes())
	return err
}
