package db

import (
	"errors"
	"fmt"
	"github.com/smapig/go-ddd-sample/core/domain/entity"
	"github.com/smapig/go-ddd-sample/core/infrastructure/db"
	"github.com/smapig/go-ddd-sample/core/infrastructure/log"
	"github.com/smapig/go-ddd-sample/core/infrastructure/orm"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type feeSqlMigrator struct {
	dbContext orm.DbContext
	logger    log.Logger
}

func (f feeSqlMigrator) Migration() {
	conn := f.dbContext.DB()
	err := conn.AutoMigrate(entity.FiatPaymentNetwork{})
	if err != nil {
		f.logger.Errorf("Failed to migrate data tables for Fee service %w", err)
	}
}

func (f feeSqlMigrator) Seeding(scriptDir string) {
	var paths []string
	err := filepath.Walk(scriptDir, func(path string, info fs.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), ".sql") {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		panic(fmt.Sprintf("Failed to access seeding files with error: %s", err.Error()))
	}

	for _, path := range paths {
		if err := f.executeSeed(path); err != nil {
			panic(err)
		}
	}
}

func (f feeSqlMigrator) executeSeed(path string) error {
	reader, err := os.Open(path)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to open seeding file with error: %s", err.Error()))
	}

	defer func() {
		_ = reader.Close()
	}()

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to read seeding file with error: %s", err.Error()))
	}

	conn := f.dbContext.DB()
	err = conn.Exec(string(b)).Error
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to execute seeding script: \n %s \n %s", path, err.Error()))
	}

	return nil
}

func NewFeeSqlMigrator(ctx orm.DbContext) db.SqlMigrator {
	return &feeSqlMigrator{
		dbContext: ctx,
	}
}
