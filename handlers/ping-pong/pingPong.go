package ping_pong

import (
	"context"
	"github.com/ivansukach/grpc-server/protocol"
	"github.com/labstack/echo"
	"log"
)

func (pp *PingPong) PingPong(c echo.Context) error {
	request := &protocol.GRRequest{Req: "Ping"}
	response, err := pp.client.GiveResponse(context.Background(), request)
	log.Println("Ответ сервера:", response.GetRes())
	return err
}
