package orm

import "gorm.io/gorm"

type UnitOfWorkRepository interface {
	Repository
	CreateUoW(target interface{}, tx *gorm.DB) error
	UpdateUoW(id string, data interface{}, tx *gorm.DB) error
	DeleteUoW(id string, tx *gorm.DB) error

	DbContext() *gorm.DB
	HandleError(res *gorm.DB) error
	HandleOneError(res *gorm.DB) error
}
