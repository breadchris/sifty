package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate will set a UUID.
func (b *Base) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	b.ID = uuid
	return nil
}
