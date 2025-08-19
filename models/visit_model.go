package models

import (
	"time"

	"github.com/google/uuid"
)

type Visit struct {
	ID           uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:id"`
	VisitorID    uuid.UUID  `gorm:"type:uuid;not null;column:visitor_id"`
	CompanyID    *uuid.UUID `gorm:"type:uuid;column:company_id"`
	DivisionID   *uuid.UUID `gorm:"type:uuid;column:division_id"`
	DepartmentID *uuid.UUID `gorm:"type:uuid;column:department_id"`
	SectionID    *uuid.UUID `gorm:"type:uuid;column:section_id"`
	EmployeeID   *uuid.UUID `gorm:"type:uuid;column:employee_id"`
	FrontDeskID  *uuid.UUID `gorm:"type:uuid;column:front_desk_id"`
	AccessCardID *uuid.UUID `gorm:"type:uuid;column:access_card_id"`
	DeviceID     *uuid.UUID `gorm:"type:uuid;column:device_id"`
	CheckIn      time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP;column:check_in"`
	CheckOut     *time.Time `gorm:"column:check_out"`
	Notes        string     `gorm:"type:text;column:notes"`
	CreatedAt    time.Time  `gorm:"column:created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at"`

	Visitor    Visitor
	Company    *Company
	Division   *Division
	Department *Department
	Section    *Section
	Employee   *Employee
	FrontDesk  *FrontDeskStaff
	AccessCard *AccessCard
	Device     *Device
}

func (Visit) TableName() string { return "visits" }
