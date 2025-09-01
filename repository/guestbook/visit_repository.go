package guestbook

import (
	"guestbook_backend/db"

	"gorm.io/gorm"
)

type VisitRepository interface {
	SetDB(db *gorm.DB)
	ClearTransactionDB()
}

type visitRepository struct {
	db *gorm.DB
}

func NewVisitRepository() VisitRepository {
	return &visitRepository{
		db: db.GetDB(),
	}
}

// // Override db (misal untuk transaksi)
func (s *visitRepository) SetDB(db *gorm.DB) {
	s.db = db
}

func (s *visitRepository) ClearTransactionDB() {
	s.db = db.GetDB() // reset ke default non-transactional DB
}
