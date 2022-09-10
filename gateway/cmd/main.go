package main

import (
	"flag"
	"log"

	"github.com/prathusingh/gallery/gateway/config"
	"github.com/prathusingh/gallery/pkg/logger"
	"github.com/prathusingh/gallery/gateway/internal"
)

func main() {
	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	appLogger := logger.NewAppLogger(cfg.Logger)

	s := server.NewServer(appLogger, cfg)
	s.Run()
}
