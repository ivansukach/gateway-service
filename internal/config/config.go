package config

import (
	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
	"time"
)

type Config struct {
	Port             int    `env:"PORT" envDefault:"8081"`
	TTL              int64  `env:"TTL"`
	SecretKeyAuth    string `env:"SecretKeyAuth"`
	SecretKeyRefresh string `env:"SecretKeyRefresh"`
}

func Load() (cfg Config) {
	cfg.TTL = time.Now().Add(time.Minute).Unix() - time.Now().Unix()
	cfg.SecretKeyAuth = "afrgdrsgfdhdfsgds"
	cfg.SecretKeyRefresh = "hkogijjdfiouhdfguih"
	if err := env.Parse(&cfg); err != nil {
		log.Printf("%+v\n", err)
	}
	return
}
