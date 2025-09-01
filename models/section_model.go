package models

import (
	"time"

	"github.com/google/uuid"
)

// Section (Level 3)
type Section struct {
	ID           uuid.UUID     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:id"`
	DepartmentID *uuid.UUID    `gorm:"type:uuid;column:department_id"`
	Name         string        `gorm:"size:100;not null;column:name"`
	Code         string        `gorm:"column:code"`
	PolicyID     *uuid.UUID    `gorm:"column:policy_id"`
	Policy       *AccessPolicy `gorm:"foreignKey:PolicyID"`
	CreatedAt    time.Time     `gorm:"column:created_at"`
	UpdatedAt    time.Time     `gorm:"column:updated_at"`

	Department *Department
	Employees  []Employee `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:SectionID"`
}

func (Section) TableName() string { return "sections" }
