package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/prathusingh/gallery/gateway/config"
	"github.com/prathusingh/gallery/pkg/logger"
)

type server struct {
	log  logger.Logger
	cfg  *config.Config
	echo *echo.Echo
}

func NewServer(log logger.Logger, cfg *config.Config) *server {
	return &server{log: log, cfg: cfg, echo: echo.New()}
}

func (s *server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	s.log.Infof("API Gateway is listening on PORT: %s", s.cfg.Http.Port)

	if err := s.echo.Server.Shutdown(ctx); err != nil {
		s.log.WarnMsg("echo.Server.Shutdown", err)
	}

	return nil
}
