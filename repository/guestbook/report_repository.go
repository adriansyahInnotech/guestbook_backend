package guestbook

import (
	"guestbook_backend/db"

	"gorm.io/gorm"
)

type ReportRepository interface {
	SetDB(db *gorm.DB)
	ClearTransactionDB()
}

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository() ReportRepository {
	return &reportRepository{db: db.GetDB()}
}

func (s *reportRepository) SetDB(db *gorm.DB) {
	s.db = db
}

func (s *reportRepository) ClearTransactionDB() {
	s.db = db.GetDB() // reset ke default non-transactional DB
}
