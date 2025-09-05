package guestbook

import (
	"guestbook_backend/db"

	"gorm.io/gorm"
)

type AccessCardRepository interface {
	SetDB(db *gorm.DB)
	ClearTransactionDB()
}

type accessCardRepository struct {
	db *gorm.DB
}

func NewAccessCardRepository() AccessCardRepository {
	return &accessCardRepository{
		db: db.GetDB(),
	}
}

// // Override db (misal untuk transaksi)
func (s *accessCardRepository) SetDB(db *gorm.DB) {
	s.db = db
}

func (s *accessCardRepository) ClearTransactionDB() {
	s.db = db.GetDB() // reset ke default non-transactional DB
}
