package dtos

import "time"

type VisitFilter struct {
	Page     int
	PageSize int

	// filter by text
	VisitorName    string
	DivisionName   string
	DepartmentName string
	SectionName    string

	// filter by tanggal check in / out
	CheckInFrom *time.Time
	CheckInTo   *time.Time
}
