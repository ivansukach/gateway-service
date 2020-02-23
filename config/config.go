package config

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/ivansukach/gateway-service/repositories"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"time"
)

type Config struct {
	Port                 int    `env:"PORT" envDefault:"8081"`
	Version              int32  `env:"VERSION"`
	AuthGRPCEndpoint     string `env:"AuthGRPCEndpointProfile"`
	BookGRPCEndpoint     string `env:"AuthGRPCEndpointBook"`
	PingPongGRPCEndpoint string `env:"AuthGRPCEndpointPingPong"`
	ProfileGRPCEndpoint  string `env:"AuthGRPCEndpointProfile"`
}

func Load() (cfg Config) {
	db, err := sqlx.Connect("postgres", "user=su password=su dbname=gateways sslmode=disable")
	if err != nil {
		log.Error(err)
		return
	}
	rps := repositories.New(db)
	endpoints, err := rps.GetNewGateways()
	latest := 0
	for i := range endpoints {
		fmt.Println("Record №", i, endpoints[i])
		if i > latest {
			latest = i
		}
	}
	cfg.Version = endpoints[latest].Version
	cfg.AuthGRPCEndpoint = endpoints[latest].AuthService
	cfg.BookGRPCEndpoint = endpoints[latest].BookService
	cfg.PingPongGRPCEndpoint = endpoints[latest].PingPong
	cfg.ProfileGRPCEndpoint = endpoints[latest].ProfileService
	if err := env.Parse(&cfg); err != nil {
		log.Printf("%+v\n", err)
	}
	return
}

func UpdateConfig(ttl int64) (cfg Config, newTtl int64) {
	for {
		if time.Now().Unix() >= ttl {
			cfg = Load()
			newTtl = time.Now().Add(time.Minute).Unix() //*30
		}
	}
	return
}
