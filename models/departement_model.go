package models

import (
	"time"

	"github.com/google/uuid"
)

// Department (Level 2)
type Department struct {
	ID         uuid.UUID     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:id"`
	DivisionID *uuid.UUID    `gorm:"type:uuid;column:division_id"`
	Name       string        `gorm:"size:100;not null;column:name"`
	Code       string        `gorm:"column:code"`
	PolicyID   *uuid.UUID    `gorm:"column:policy_id"`
	Policy     *AccessPolicy `gorm:"foreignKey:PolicyID"`
	CreatedAt  time.Time     `gorm:"column:created_at"`
	UpdatedAt  time.Time     `gorm:"column:updated_at"`

	Division *Division
	Sections []Section `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:DepartmentID"`
}

func (Department) TableName() string { return "departments" }
