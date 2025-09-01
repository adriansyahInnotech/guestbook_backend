package models

import (
	"time"

	"github.com/google/uuid"
)

// Company (Level 0)
type Company struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:id"`
	Name      string    `gorm:"size:150;not null;unique;column:name"`
	Code      string    `gorm:"column:code"`
	Address   string    `gorm:"size:250;column:address"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	Divisions []Division `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CompanyID"`
}

func (Company) TableName() string { return "companies" }
