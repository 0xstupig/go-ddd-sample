package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ID struct {
	ID int64 `gorm:"primaryKey"`
}

type AuditTime struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SoftDelete struct {
	DeletedAt gorm.DeletedAt
}

type BaseEntity struct {
	ID
	AuditTime
	SoftDelete
}

type BaseUuidEntity struct {
	ID uuid.UUID `gorm:"default:gen_random_uuid()"`

	AuditTime
	SoftDelete
}

