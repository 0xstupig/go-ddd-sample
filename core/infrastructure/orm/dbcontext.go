package orm

import (
	"fmt"
	"github.com/avast/retry-go"
	"github.com/pkg/errors"
	"github.com/smapig/go-ddd-sample/core/infrastructure/config"
	"github.com/smapig/go-ddd-sample/core/infrastructure/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type DbContext interface {
	DB() *gorm.DB
}

type gDbContext struct {
	db *gorm.DB
}

func (c gDbContext) DB() *gorm.DB {
	return c.db
}

func NewDBContext(logger log.Logger, cfg config.AppConfig) (DbContext, error) {
	dbConf := cfg.Db
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dbConf.Host, dbConf.Username, dbConf.Password, dbConf.DbName, dbConf.Port)
	var (
		db  *gorm.DB
		err error
	)

	err = retry.Do(func() error {
		gormConfig := &gorm.Config{}
		db, err = gorm.Open(postgres.Open(dsn), gormConfig)
		return err
	},
		retry.Attempts(dbConf.RetryAttempts),
		retry.Delay(time.Second),
	)

	if err != nil {
		logger.Error(errors.Wrap(err, "failed to initialize DbContext"))
		return nil, err
	}

	logger.Info("successfully initialize DbContext")
	if cfg.Debug {
		db = db.Debug()
	}

	return gDbContext{db}, nil
}
