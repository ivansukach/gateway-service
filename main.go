package main

import (
	"fmt"
	"github.com/ivansukach/book-service/protocol"
	"github.com/ivansukach/gateway-service/config"
	"github.com/ivansukach/gateway-service/handlers/authentication-service"
	"github.com/ivansukach/gateway-service/handlers/book-service"
	"github.com/ivansukach/gateway-service/handlers/ping-pong"
	"github.com/ivansukach/gateway-service/handlers/profile-service"
	protocolPP "github.com/ivansukach/grpc-server/protocol"
	protocolPS "github.com/ivansukach/profile-service/protocol"
	"github.com/labstack/echo"
	protocolAS "github.com/leshachaplin/grpc-server/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"time"
)

func isItTimeToUpdateConfig(ttl int64, version int32, e *echo.Echo) {
	for {
		if time.Now().Unix() >= ttl {
			cfg := config.Load()
			if cfg.Version != version {
				go NewUrls(e)
				return
			}
			ttl = time.Now().Add(time.Minute).Unix() //*30
		} else {
			diff := ttl - time.Now().Unix()
			time.Sleep(time.Duration(diff) * time.Second)
		}
	}
}
func NewConnectionInterfaces(cfg config.Config, opts grpc.DialOption) (pp *ping_pong.PingPong,
	ps *profile_service.ProfileAccounting, bs *book_service.BookAccounting,
	as *authentication_service.Auth) {

	clientConnPingPongInterface, err := grpc.Dial(cfg.PingPongGRPCEndpoint, opts)
	if err != nil {
		log.Fatal(err)
	}

	defer clientConnPingPongInterface.Close()

	clientPP := protocolPP.NewGetResponseClient(clientConnPingPongInterface)
	pp = ping_pong.NewHandlerPP(clientPP)

	clientConnProfileInterface, err := grpc.Dial(cfg.ProfileGRPCEndpoint, opts)
	if err != nil {
		log.Fatal(err)
	}

	defer clientConnProfileInterface.Close()

	clientPS := protocolPS.NewProfileServiceClient(clientConnProfileInterface)
	ps = profile_service.NewHandlerPS(clientPS)

	clientConnBookInterface, err := grpc.Dial(cfg.BookGRPCEndpoint, opts)
	if err != nil {
		log.Fatal(err)
	}
	defer clientConnBookInterface.Close()
	client := protocol.NewBookServiceClient(clientConnPingPongInterface)
	bs = book_service.NewHandler(client)

	clientConnAuthServiceInterface, err := grpc.Dial(cfg.AuthGRPCEndpoint, opts)
	if err != nil {
		log.Fatal(err)
	}

	defer clientConnAuthServiceInterface.Close()

	clientAS := protocolAS.NewAuthServiceClient(clientConnPingPongInterface)
	as = authentication_service.NewHandlerAuth(clientAS)
	return
}
func NewUrls(e *echo.Echo) {
	cfg := config.Load()
	version := cfg.Version
	ttl := time.Now().Add(time.Minute).Unix() //*30
	opts := grpc.WithInsecure()

	pp, ps, bs, as := NewConnectionInterfaces(cfg, opts)
	e.POST("book/create", bs.Create)
	e.POST("book/update", bs.Update)
	e.POST("book/delete", bs.Delete)
	e.POST("book/getById", bs.GetById)
	e.POST("book/listing", bs.Listing)

	e.POST("games/pingPong", pp.PingPong)

	e.POST("profile/create", ps.Create)
	e.POST("profile/update", ps.Update)
	e.POST("profile/delete", ps.Delete)
	e.POST("profile/getByLogin", ps.GetByLogin)
	e.POST("profile/listing", ps.Listing)

	e.POST("auth/signIn", as.SignIn)
	e.POST("auth/signUp", as.SignUp)
	e.POST("auth/deleteUser", as.DeleteUser)
	e.POST("auth/deleteClaims", as.DeleteClaims)
	e.POST("auth/addClaims", as.AddClaims)
	go isItTimeToUpdateConfig(ttl, version, e)
}

func main() {
	log.Println("Client started")
	cfg := config.Load()
	version := cfg.Version
	ttl := time.Now().Add(time.Minute).Unix() //*30
	opts := grpc.WithInsecure()
	pp, ps, bs, as := NewConnectionInterfaces(cfg, opts)
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
	e.POST("profile/getByLogin", ps.GetByLogin)
	e.POST("profile/listing", ps.Listing)

	e.POST("auth/signIn", as.SignIn)
	e.POST("auth/signUp", as.SignUp)
	e.POST("auth/deleteUser", as.DeleteUser)
	e.POST("auth/deleteClaims", as.DeleteClaims)
	e.POST("auth/addClaims", as.AddClaims)
	go isItTimeToUpdateConfig(ttl, version, e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Port)))
}
