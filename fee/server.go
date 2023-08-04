package main

import (
	"github.com/smapig/go-ddd-sample/core/infrastructure/config"
	"github.com/smapig/go-ddd-sample/core/infrastructure/log"
	"github.com/smapig/go-ddd-sample/fee/controller"
	"github.com/smapig/go-ddd-sample/fee/http"
)

type Server interface {
	StartHTTP() error
}

type serverImpl struct {
	logger         log.Logger
	configProvider config.AppConfig
	endpoint       controller.Controller
	httpServer     http.Server
}

func NewServer(
	logger log.Logger,
	configProvider config.AppConfig,
	endpoint controller.Controller,
) Server {
	return &serverImpl{
		httpServer:     http.NewServer(endpoint, configProvider, logger),
		logger:         logger,
		configProvider: configProvider,
		endpoint:       endpoint,
	}
}

func (s *serverImpl) StartHTTP() error {
	return s.httpServer.Serve()
}
