package orm

import "gorm.io/gorm"

type UnitOfWorkRepository interface {
	Repository
	CreateUoW(target interface{}, tx *gorm.DB) error
	UpdateUoW(target interface{}, tx *gorm.DB) error
	DeleteUoW(target interface{}, tx *gorm.DB) error

	DbContext() *gorm.DB
	DbContextWithPreloads(preloads []string) *gorm.DB
	HandleError(res *gorm.DB) error
	HandleOneError(res *gorm.DB) error
}
