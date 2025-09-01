package models

import (
	"time"

	"github.com/google/uuid"
)

type Division struct {
	ID        uuid.UUID     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:id"`
	CompanyID *uuid.UUID    `gorm:"type:uuid;column:company_id"`
	Name      string        `gorm:"size:100;not null;column:name"`
	Code      string        `gorm:"column:code"`
	PolicyID  *uuid.UUID    `gorm:"column:policy_id"`
	Policy    *AccessPolicy `gorm:"foreignKey:PolicyID"`
	CreatedAt time.Time     `gorm:"column:created_at"`
	UpdatedAt time.Time     `gorm:"column:updated_at"`

	Company     *Company
	Departments []Department `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:DivisionID"`
}

func (Division) TableName() string { return "divisions" }
