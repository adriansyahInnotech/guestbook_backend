package customresponse

import (
	"time"

	"github.com/google/uuid"
)

type AccessPolicy struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Devices []Device `json:"device"`

	// Devices []Device `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:PolicyID"`
}

type Device struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Location string    `json:"location"`
	IsActive bool      `json:"is_active"`
}
