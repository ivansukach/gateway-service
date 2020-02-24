package main

import (
	"fmt"
	"github.com/ivansukach/gateway-service/internal/config"
	"github.com/ivansukach/gateway-service/internal/handlers/authentication-service"
	"github.com/ivansukach/gateway-service/internal/handlers/book-service"
	"github.com/ivansukach/gateway-service/internal/handlers/ping-pong"
	"github.com/ivansukach/gateway-service/internal/handlers/profile-service"
	"github.com/ivansukach/gateway-service/internal/repositories"
	"github.com/ivansukach/gateway-service/internal/services"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"time"
)

func isItTimeToUpdateConfig(e *echo.Echo, TTL int64, rps repositories.Repository, opts grpc.DialOption, cfg *config.Config) {
	for {
		endpoints, err := services.UpdateEndpoints(rps)
		if err != nil {
			log.Fatal(err)
		}
		if time.Now().Unix() >= TTL {
			for i := range endpoints {
				switch endpoints[i].Name {
				case "authentication-service":
					authentication_service.NewHandlerAuth(endpoints[i].Url, endpoints[i].Version, e, opts)
				case "book-service":
					book_service.NewHandlerBS(endpoints[i].Url, endpoints[i].Version, e, opts)
				case "ping-pong":
					ping_pong.NewHandlerPP(endpoints[i].Url, endpoints[i].Version, e, opts)
				case "profile-service":
					profile_service.NewHandlerPS(endpoints[i].Url, endpoints[i].Version, e, opts)
				}
			}
			TTL = time.Now().Unix() + cfg.TTL
		} else {
			diff := TTL - time.Now().Unix()
			time.Sleep(time.Duration(diff) * time.Second)
		}
	}
}

//func NewUrls(e *echo.Echo) {
//	cfg := config.Load()
//	version := cfg.Version
//	opts := grpc.WithInsecure()
//	v:="v"+version
//	pp, ps, bs, as, jwt := NewConnectionInterfaces(cfg, opts)
//
//
//
//
//	go isItTimeToUpdateConfig(e)
//}

func main() {
	log.Println("Client started")
	db, err := sqlx.Connect("postgres", "user=su password=su dbname=gateways sslmode=disable")
	if err != nil {
		log.Error(err)
		return
	}
	cfg := config.Load()
	rps := repositories.New(db)
	ttl := time.Now().Unix()
	opts := grpc.WithInsecure()
	e := echo.New()
	e.Static("/", "static")
	e.Use(middleware.Gzip())
	go isItTimeToUpdateConfig(e, ttl, rps, opts, &cfg)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Port)))
}
