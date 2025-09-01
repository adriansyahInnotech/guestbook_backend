package models

import (
	"time"

	"github.com/google/uuid"
)

type Visitor struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:id"`
	FullName     string    `gorm:"size:100;not null;column:full_name"`
	Company      string    `gorm:"size:150;column:company"`
	Phone        string    `gorm:"size:20;column:phone"`
	IDCardType   string    `gorm:"type:id_card_type_enum;not null;column:id_card_type"`
	IDCardNumber string    `gorm:"size:50;not null;uniqueIndex;column:id_card_number"`
	DataCard     string    `gorm:"column:data_card"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`

	AccessCards []AccessCard `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:VisitorID"`
	Visits      []Visit      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:VisitorID"`
}

func (Visitor) TableName() string { return "visitors" }
