package models

import (
	"time"

	"github.com/google/uuid"
)

type AccessCard struct {
	ID         uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:id"`
	VisitorID  uuid.UUID  `gorm:"type:uuid;not null;column:visitor_id"`
	CardNumber string     `gorm:"size:50;not null;uniqueIndex;column:card_number"`
	PolicyID   *uuid.UUID `gorm:"type:uuid;column:policy_id"` // optional, bisa null → fallback policy
	CreatedAt  time.Time  `gorm:"column:created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at"`

	Visitor Visitor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:VisitorID"`
	Policy  *AccessPolicy
	Visits  []Visit `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:AccessCardID"`
}

func (AccessCard) TableName() string { return "access_cards" }
