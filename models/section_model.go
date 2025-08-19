package models

import (
	"time"

	"github.com/google/uuid"
)

// Section (Level 3)
type Section struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:id"`
	DepartmentID uuid.UUID `gorm:"type:uuid;not null;column:department_id"`
	Name         string    `gorm:"size:100;not null;column:name"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`

	Department Department
	Employees  []Employee `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:SectionID"`
}

func (Section) TableName() string { return "sections" }
