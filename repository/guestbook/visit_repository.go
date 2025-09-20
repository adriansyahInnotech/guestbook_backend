package guestbook

import (
	"guestbook_backend/db"
	"guestbook_backend/models"

	"gorm.io/gorm"
)

type VisitRepository interface {
	SetDB(db *gorm.DB)
	ClearTransactionDB()
	Add(visit *models.Visit) error
	GetVisitActiveByCard(cardID string) (*models.Visit, error)
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
func (s *visitRepository) Add(visit *models.Visit) error {

	result := s.db.Create(visit)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *visitRepository) GetVisitActiveByCard(cardID string) (*models.Visit, error) {

	visit := new(models.Visit)

	result := s.db.
		Joins("JOIN access_cards ON access_cards.id = visits.access_card_id ").
		Where("access_cards.card_number = ? AND visits.check_out IS NULL", cardID).
		Order("visits.check_in DESC").
		First(&visit)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}

	return visit, nil
}
