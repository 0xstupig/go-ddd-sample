package domain

import (
	"github.com/smapig/go-ddd-sample/core/infrastructure/log"
	"github.com/smapig/go-ddd-sample/core/infrastructure/orm"
	"gorm.io/gorm"
)

type genericRepository struct {
	logger log.Logger
	dbContext orm.DbContext
}

func (g genericRepository) GetAll(target interface{}, limit, offset int, preloads ...string) error {
	panic("implement me")
}

func (g genericRepository) GetBy(target interface{}, filters map[string]interface{}, limit, offset int, preloads ...string) error {
	panic("implement me")
}

func (g genericRepository) GetOne(target interface{}, filters map[string]interface{}, preloads ...string) error {
	panic("implement me")
}

func (g genericRepository) Create(target interface{}) error {
	panic("implement me")
}

func (g genericRepository) Update(id string, data interface{}) error {
	panic("implement me")
}

func (g genericRepository) Delete(id string) error {
	panic("implement me")
}

func (g genericRepository) CreateUoW(target interface{}, tx *gorm.DB) error {
	panic("implement me")
}

func (g genericRepository) UpdateUoW(id string, data interface{}, tx *gorm.DB) error {
	panic("implement me")
}

func (g genericRepository) DeleteUoW(id string, tx *gorm.DB) error {
	panic("implement me")
}

func (g genericRepository) DbContext() *gorm.DB {
	panic("implement me")
}

func (g genericRepository) HandleError(res *gorm.DB) error {
	panic("implement me")
}

func (g genericRepository) HandleOneError(res *gorm.DB) error {
	panic("implement me")
}

func NewGenericRepository(dbContext orm.DbContext, logger log.Logger) orm.UnitOfWorkRepository {
	return &genericRepository{
		logger, dbContext,
	}
}