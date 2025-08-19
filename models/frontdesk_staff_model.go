package models

import (
	"time"

	"github.com/google/uuid"
)

type FrontDeskStaff struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:id"`
	FullName  string    `gorm:"size:100;not null;column:full_name"`
	Email     string    `gorm:"size:150;unique;column:email"`
	Phone     string    `gorm:"size:20;column:phone"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (FrontDeskStaff) TableName() string { return "front_desk_staffs" }
