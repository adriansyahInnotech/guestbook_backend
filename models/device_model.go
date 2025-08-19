package models

import (
	"time"

	"github.com/google/uuid"
)

type Device struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:id"`
	Name      string    `gorm:"size:100;not null;column:name"`
	Location  string    `gorm:"size:150;column:location"`
	ApiKey    string    `gorm:"size:255;not null;unique;column:api_key"`
	IsActive  bool      `gorm:"default:true;column:is_active"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	Visits []Visit `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:DeviceID"`
}

func (Device) TableName() string { return "devices" }
