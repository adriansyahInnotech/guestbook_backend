package models

import (
	"time"

	"github.com/google/uuid"
)

type Employee struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:id"`
	SectionID uuid.UUID `gorm:"type:uuid;not null;column:section_id"`
	FullName  string    `gorm:"size:100;not null;column:full_name"`
	Email     string    `gorm:"size:150;unique;column:email"`
	Phone     string    `gorm:"size:20;column:phone"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	Section Section
}

func (Employee) TableName() string { return "employees" }
