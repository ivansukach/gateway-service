package main

import (
	"fmt"
	"github.com/ivansukach/book-service-http/handlers"
	"github.com/ivansukach/book-service/protocol"
	"github.com/ivansukach/gateway-service/config"
	"github.com/ivansukach/gateway-service/handlers"
	"github.com/ivansukach/gateway-service/handlers/authentication-service"
	"github.com/ivansukach/gateway-service/handlers/ping-pong"
	"github.com/ivansukach/gateway-service/handlers/profile-service"
	protocolPP "github.com/ivansukach/grpc-server/protocol"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Client started")
	cfg := config.Load()
	opts := grpc.WithInsecure() //WithInsecure returns a DialOption which disables transport security for this ClientConn.
	// Note that transport security is required unless WithInsecure is set.

	clientConnPingPongInterface, err := grpc.Dial(cfg.AuthGRPCEndpointPingPong, opts)
	if err != nil {
		log.Fatal(err)
	}

	defer clientConnPingPongInterface.Close()

	clientPP := protocolPP.NewGetResponseClient(clientConnPingPongInterface)
	pp := ping_pong.NewHandlerPP(clientPP)

	clientConnProfileInterface, err := grpc.Dial(cfg.AuthGRPCEndpointProfile, opts)
	if err != nil {
		log.Fatal(err)
	}

	defer clientConnProfileInterface.Close()

	clientPS := protocolPP.NewGetResponseClient(clientConnProfileInterface)
	ps := profile_service.NewHandlerPS(clientPS)

	clientConnBookInterface, err := grpc.Dial(cfg.AuthGRPCEndpointBook, opts)
	if err != nil {
		log.Fatal(err)
	}
	defer clientConnBookInterface.Close()
	client := protocol.NewBookServiceClient(clientConnPingPongInterface)
	bs := handlers.NewHandler(client)
	e := echo.New()
	e.POST("book/create", bs.Create)
	e.POST("book/update", bs.Update)
	e.POST("book/delete", bs.Delete)
	e.POST("book/getById", bs.GetById)
	e.POST("book/listing", bs.Listing)
	e.POST("games/pingPong", pp.PingPong)
	e.POST("profile/create", ps.Create)
	e.POST("profile/update", ps.Update)
	e.POST("profile/delete", ps.Delete)
	e.POST("profile/getById", ps.GetById)
	e.POST("profile/listing", ps.Listing)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Port)))
}
