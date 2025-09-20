package models

import (
	"time"

	"github.com/google/uuid"
)

type AccessPolicyDevice struct {
	AccessPolicyID uuid.UUID `gorm:"type:uuid;not null;column:access_policy_id"`
	DeviceID       uuid.UUID `gorm:"type:uuid;not null;column:device_id"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`

	AccessPolicy AccessPolicy `gorm:"foreignKey:AccessPolicyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Device       Device       `gorm:"foreignKey:DeviceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (AccessPolicyDevice) TableName() string { return "access_policy_devices" }
