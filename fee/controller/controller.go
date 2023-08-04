package controller

import (
	"github.com/smapig/go-ddd-sample/core/infrastructure/config"
	"github.com/smapig/go-ddd-sample/core/infrastructure/log"
	"github.com/smapig/go-ddd-sample/core/service/fee"
)

type Controller interface {
	FeeController
}

type controllerImpl struct {
	conf       config.AppConfig
	logger     log.Logger
	feeService fee.FeeService
}

func NewController(conf config.AppConfig, logger log.Logger, feeService fee.FeeService) Controller {
	return &controllerImpl{
		conf, logger, feeService,
	}
}
