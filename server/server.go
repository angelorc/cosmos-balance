package server

import (
	"context"
	"github.com/angelorc/cosmos-tracker/client/chain"
	_ "github.com/angelorc/cosmos-tracker/swagger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"time"
)

type Server struct {
	*echo.Echo
	chains *chain.Chains
	logger *zap.Logger
}

// @title Cosmos Tracker Server API
// @version 1.0
// @description The cosmos tracker rest server.

func NewServer(chains *chain.Chains, logger *zap.Logger) *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Debug = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	s := &Server{
		Echo:   e,
		chains: chains,
		logger: logger,
	}
	s.registerRoutes()

	return s
}

func (s *Server) ShutdownWithTimeout(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return s.Shutdown(ctx)
}
