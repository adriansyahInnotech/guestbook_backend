package models

import (
	"time"

	"github.com/google/uuid"
)

type Device struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;uniqueIndex:uniq_device_id_name;column:id"`
	Name     string    `gorm:"size:100;not null;uniqueIndex;uniqueIndex:uniq_device_id_name;column:name"`
	Location string    `gorm:"size:150;column:location"`
	ApiKey   string    `gorm:"size:255;not null;unique;column:api_key"`
	IsActive bool      `gorm:"default:true;column:is_active"`
	// PolicyID  *uuid.UUID    `gorm:"column:policy_id"`
	// Policy    *AccessPolicy `gorm:"foreignKey:PolicyID"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	AccessPolicies []AccessPolicyDevice `gorm:"foreignKey:DeviceID"`
	Visits         []Visit              `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:DeviceID"`
}

func (Device) TableName() string { return "devices" }
