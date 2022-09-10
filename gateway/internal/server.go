package server

import (
	"github.com/prathusingh/gallery/gateway/config"
	"github.com/prathusingh/gallery/pkg/logger"
)

type server struct {
	log logger.Logger
	cfg *config.Config
}

func (s *server) Run() error {
	s.log.Infof("API Gateway is listening on PORT: %s", s.cfg.Http.Port)

	if err := s.echo.Server.Shutdown(ctx); err != nil {
		s.log.WarnMsg("echo.Server.Shutdown", err)
	}
}
