package models

import (
	"time"

	"github.com/google/uuid"
)

type AccessCard struct {
	ID         uuid.UUID     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:id"`
	CardNumber string        `gorm:"size:50;not null;unique;column:card_number"`
	VisitorID  *uuid.UUID    `gorm:"type:uuid;column:visitor_id"`
	IssuedAt   *time.Time    `gorm:"column:issued_at"`
	ReturnedAt *time.Time    `gorm:"column:returned_at"`
	PolicyID   *uuid.UUID    `gorm:"column:policy_id"`
	Policy     *AccessPolicy `gorm:"foreignKey:PolicyID"`
	CreatedAt  time.Time     `gorm:"column:created_at"`
	UpdatedAt  time.Time     `gorm:"column:updated_at"`
	Visitor    *Visitor
}

func (AccessCard) TableName() string { return "access_cards" }
