package domain

import (
	"fmt"
	"github.com/smapig/go-ddd-sample/core/infrastructure/log"
	"github.com/smapig/go-ddd-sample/core/infrastructure/orm"
	"gorm.io/gorm"
)

type genericRepository struct {
	logger    log.Logger
	dbContext orm.DbContext
}

func (g genericRepository) DbContextWithPreloads(preloads []string) *gorm.DB {
	conn := g.DbContext()

	for _, preload := range preloads {
		conn = conn.Preload(preload)
	}

	return conn
}

func (g genericRepository) GetAll(target interface{}, limit, offset int, preloads ...string) error {
	g.logger.Debugf("Executing GetAll on %T", target)
	queryBuilder := g.DbContextWithPreloads(preloads).Unscoped()

	if limit > 0 {
		queryBuilder = queryBuilder.Limit(limit)
	}

	if offset >= 0 {
		queryBuilder = queryBuilder.Offset(offset)
	}

	res := queryBuilder.Find(target)
	return g.HandleError(res)
}

func (g genericRepository) GetBy(target interface{}, filters map[string]interface{}, limit, offset int,
preloads ...string) (interface{}, error) {
	g.logger.Debugf("Executing GetBy on %T", target)
	queryBuilder := g.DbContextWithPreloads(preloads)

	for field, value := range filters {
		queryBuilder = queryBuilder.Where(fmt.Sprintf("%v = ?", field), value)
	}

	if limit > 0 {
		queryBuilder = queryBuilder.Limit(limit)
	}

	if offset >= 0 {
		queryBuilder = queryBuilder.Offset(offset)
	}

	res := queryBuilder.Find(&target)
	return target, g.HandleError(res)
}

func (g genericRepository) GetOne(target interface{}, filters map[string]interface{}, preloads ...string) error {
	g.logger.Debugf("Executing GetOne on %T", target)
	queryBuilder := g.DbContextWithPreloads(preloads)

	for field, value := range filters {
		queryBuilder = queryBuilder.Where(fmt.Sprintf("%v = ?", field), value)
	}

	res := queryBuilder.First(target)
	return g.HandleOneError(res)
}

func (g genericRepository) Create(target interface{}) error {
	g.logger.Debugf("Executing Create on %T", target)
	res := g.DbContext().Create(target)
	return g.HandleError(res)

}

func (g genericRepository) Update(target interface{}) error {
	g.logger.Debugf("Executing Update on %T", target)
	res := g.DbContext().Save(target)
	return g.HandleError(res)}

func (g genericRepository) Delete(target interface{}) error {
	g.logger.Debugf("Executing Update on %T", target)
	res := g.DbContext().Delete(target)
	return g.HandleError(res)
}

func (g genericRepository) CreateUoW(target interface{}, tx *gorm.DB) error {
	g.logger.Debugf("Executing CreateUoW on %T", target)
	res := tx.Create(target)
	return g.HandleError(res)
}

func (g genericRepository) UpdateUoW(target interface{}, tx *gorm.DB) error {
	g.logger.Debugf("Executing UpdateUoW on %T", target)
	res := tx.Save(target)
	return g.HandleError(res)
}

func (g genericRepository) DeleteUoW(target interface{}, tx *gorm.DB) error {
	g.logger.Debugf("Executing DeleteUoW on %T", target)
	res := tx.Delete(target)
	return g.HandleError(res)
}

func (g genericRepository) DbContext() *gorm.DB {
	return g.dbContext.DB()
}

func (g genericRepository) HandleError(res *gorm.DB) error {
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		err := fmt.Errorf("Error: %w", res.Error)
		g.logger.Error(err)
		return err
	}

	return nil
}

func (g genericRepository) HandleOneError(res *gorm.DB) error {
	if err := g.HandleError(res); err != nil {
		return err
	}

	if res.RowsAffected != 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func NewGenericRepository(dbContext orm.DbContext, logger log.Logger) orm.UnitOfWorkRepository {
	return &genericRepository{
		logger, dbContext,
	}
}
