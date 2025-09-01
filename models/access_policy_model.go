package models

import (
	"time"

	"github.com/google/uuid"
)

type AccessPolicy struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:id"`
	Name      string    `gorm:"size:100;not null;unique;column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	Devices []Device `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:PolicyID"`
}
