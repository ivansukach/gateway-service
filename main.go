package main

import (
	"fmt"
	"github.com/ivansukach/book-service/protocol"
	"github.com/ivansukach/gateway-service/config"
	"github.com/ivansukach/gateway-service/handlers/authentication-service"
	"github.com/ivansukach/gateway-service/handlers/book-service"
	"github.com/ivansukach/gateway-service/handlers/ping-pong"
	"github.com/ivansukach/gateway-service/handlers/profile-service"
	"github.com/ivansukach/gateway-service/middlewares"
	protocolPP "github.com/ivansukach/grpc-server/protocol"
	protocolPS "github.com/ivansukach/profile-service/protocol"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	protocolAS "github.com/leshachaplin/grpc-server/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"strconv"
	"time"
)

func isItTimeToUpdateConfig(e *echo.Echo) {
	for {
		cfg := config.Load()
		version := cfg.Version
		if time.Now().Unix() >= cfg.TTL {
			cfg := config.Load()
			if cfg.Version != version {
				go NewUrls(e)
				return
			}

		} else {
			diff := cfg.TTL - time.Now().Unix()
			time.Sleep(time.Duration(diff) * time.Second)
		}
	}
}
func NewConnectionInterfaces(cfg config.Config, opts grpc.DialOption) (pp *ping_pong.PingPong,
	ps *profile_service.ProfileAccounting, bs *book_service.BookAccounting,
	as *authentication_service.Auth, jwt *middlewares.JWT) {

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
	jwt = middlewares.NewJWT(clientAS)
	as = authentication_service.NewHandlerAuth(clientAS)
	return
}
func NewUrls(e *echo.Echo) {
	cfg := config.Load()
	version := cfg.Version
	opts := grpc.WithInsecure()

	pp, ps, bs, as, jwt := NewConnectionInterfaces(cfg, opts)
	e.POST("v"+strconv.Itoa(int(version))+"book/create", bs.Create)
	e.POST("v"+strconv.Itoa(int(version))+"book/update", bs.Update)
	e.POST("v"+strconv.Itoa(int(version))+"book/delete", bs.Delete)
	e.POST("v"+strconv.Itoa(int(version))+"book/getById", bs.GetById)
	e.POST("v"+strconv.Itoa(int(version))+"book/listing", bs.Listing)

	e.POST("v"+strconv.Itoa(int(version))+"games/pingPong", pp.PingPong)

	e.POST("v"+strconv.Itoa(int(version))+"profile/create", ps.Create)
	e.POST("v"+strconv.Itoa(int(version))+"profile/update", ps.Update)
	e.POST("v"+strconv.Itoa(int(version))+"profile/delete", ps.Delete)
	e.POST("v"+strconv.Itoa(int(version))+"profile/getByLogin", ps.GetByLogin)
	e.POST("v"+strconv.Itoa(int(version))+"profile/listing", ps.Listing)

	e.POST("v"+strconv.Itoa(int(version))+"auth/signIn", as.SignIn)
	e.POST("v"+strconv.Itoa(int(version))+"auth/signUp", as.SignUp)
	e.POST("v"+strconv.Itoa(int(version))+"auth/deleteUser", as.DeleteUser, jwt.Middleware)
	e.POST("v"+strconv.Itoa(int(version))+"auth/deleteClaims", as.DeleteClaims, jwt.Middleware)
	e.POST("v"+strconv.Itoa(int(version))+"auth/addClaims", as.AddClaims, jwt.Middleware)
	go isItTimeToUpdateConfig(e)
}

func main() {
	log.Println("Client started")
	cfg := config.Load()
	version := cfg.Version
	opts := grpc.WithInsecure()
	pp, ps, bs, as, jwt := NewConnectionInterfaces(cfg, opts)
	e := echo.New()
	e.Static("/", "static")
	e.Use(middleware.Gzip())
	e.POST("v"+strconv.Itoa(int(version))+"/book/create", bs.Create)
	e.POST("v"+strconv.Itoa(int(version))+"/book/update", bs.Update)
	e.POST("v"+strconv.Itoa(int(version))+"/book/delete", bs.Delete)
	e.POST("v"+strconv.Itoa(int(version))+"/book/getById", bs.GetById)
	e.POST("v"+strconv.Itoa(int(version))+"/book/listing", bs.Listing)

	e.POST("v"+strconv.Itoa(int(version))+"/games/pingPong", pp.PingPong)

	e.POST("v"+strconv.Itoa(int(version))+"/profile/create", ps.Create)
	e.POST("v"+strconv.Itoa(int(version))+"/profile/update", ps.Update)
	e.POST("v"+strconv.Itoa(int(version))+"/profile/delete", ps.Delete)
	e.POST("v"+strconv.Itoa(int(version))+"/profile/getByLogin", ps.GetByLogin)
	e.POST("v"+strconv.Itoa(int(version))+"/profile/listing", ps.Listing)

	e.POST("v"+strconv.Itoa(int(version))+"/auth/signIn", as.SignIn)
	e.POST("v"+strconv.Itoa(int(version))+"/auth/signUp", as.SignUp)
	e.POST("v"+strconv.Itoa(int(version))+"/auth/deleteUser", as.DeleteUser, jwt.Middleware)
	e.POST("v"+strconv.Itoa(int(version))+"/auth/deleteClaims", as.DeleteClaims, jwt.Middleware)
	e.POST("v"+strconv.Itoa(int(version))+"/auth/addClaims", as.AddClaims, jwt.Middleware)
	go isItTimeToUpdateConfig(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Port)))
}
