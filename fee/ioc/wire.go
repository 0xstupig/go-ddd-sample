//go:build wireinject
// +build wireinject

package ioc

import (
	"github.com/smapig/go-ddd-sample/core/domain"
	"github.com/smapig/go-ddd-sample/core/infrastructure/config"
	"github.com/smapig/go-ddd-sample/core/infrastructure/db"
	"github.com/smapig/go-ddd-sample/core/infrastructure/log"
	"github.com/smapig/go-ddd-sample/core/infrastructure/orm"
	"github.com/smapig/go-ddd-sample/core/service/fee"
	"github.com/smapig/go-ddd-sample/fee/controller"
	db2 "github.com/smapig/go-ddd-sample/fee/db"
	"github.com/smapig/go-ddd-sample/fee/wsgi"
)
import "github.com/google/wire"

func InitializeConfig(confPath string) (config.AppConfig, error) {
	wire.Build(config.NewConfigProvider)
	return config.AppConfig{}, nil
}

func InitializeLogger(cfg config.AppConfig) (log.Logger, error) {
	wire.Build(log.NewLogger)
	return nil, nil
}

func InitializeDbContext(logger log.Logger, conf config.AppConfig) (orm.DbContext, error) {
	wire.Build(orm.NewDBContext)
	return nil, nil
}

func InitializeGenericRepository(dbContext orm.DbContext, logger log.Logger) (orm.UnitOfWorkRepository, error) {
	wire.Build(domain.NewGenericRepository)
	return nil, nil
}

func InitializeFeeService(conf config.AppConfig, logger log.Logger, repo orm.UnitOfWorkRepository) (fee.FeeService, error) {
	wire.Build(fee.NewFeeService)
	return nil, nil
}

func InitializeController(conf config.AppConfig, logger log.Logger, feeService fee.FeeService) (controller.Controller, error) {
	wire.Build(controller.NewController)
	return nil, nil
}

func InitializeServer(confPath string) (wsgi.Server, error) {
	wire.Build(wsgi.NewServer,
		InitializeConfig, InitializeLogger, InitializeController,
		InitializeDbContext, InitializeGenericRepository, InitializeFeeService)
	return nil, nil
}

func InitializeSqlMigrator(confPath string) (db.SqlMigrator, error) {
	wire.Build(db2.NewFeeSqlMigrator, InitializeConfig, InitializeLogger, InitializeDbContext)
	return nil, nil
}
