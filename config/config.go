package config

import (
	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Port                     int    `env:"PORT" envDefault:"8081"`
	AuthGRPCEndpointBook     string `env:"AuthGRPCEndpointBook" envDefault:"localhost:1323"`
	AuthGRPCEndpointPingPong string `env:"AuthGRPCEndpointPingPong" envDefault:"localhost:1324"`
	AuthGRPCEndpointProfile  string `env:"AuthGRPCEndpointProfile" envDefault:"localhost:1325"`
}

func Load() (cfg Config) {
	if err := env.Parse(&cfg); err != nil {
		log.Printf("%+v\n", err)
	}
	return
}
