package models

import (
	"time"

	"github.com/google/uuid"
)

// Department (Level 2)
type Department struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:id"`
	DivisionID uuid.UUID `gorm:"type:uuid;not null;column:division_id"`
	Name       string    `gorm:"size:100;not null;column:name"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`

	Division Division
	Sections []Section `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:DepartmentID"`
}

func (Department) TableName() string { return "departments" }
