package guestbook

import (
	"guestbook_backend/db"
	"guestbook_backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PolicyDeviceRepository interface {
	SetDB(db *gorm.DB)
	ClearTransactionDB()
	Add(policy []models.AccessPolicyDevice) error
	Delete(id uuid.UUID) error
}

type policyDeviceRepository struct {
	db *gorm.DB
}

func NewPolicyDeviceRepository() PolicyDeviceRepository {
	return &policyDeviceRepository{
		db: db.GetDB(),
	}
}

// // Override db (misal untuk transaksi)
func (s *policyDeviceRepository) SetDB(db *gorm.DB) {
	s.db = db
}

func (s *policyDeviceRepository) ClearTransactionDB() {
	s.db = db.GetDB() // reset ke default non-transactional DB
}

func (s *policyDeviceRepository) Add(policy []models.AccessPolicyDevice) error {

	result := s.db.Create(policy)

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (s *policyDeviceRepository) Delete(id uuid.UUID) error {
	result := s.db.Where("access_policy_id = ?", id).Delete(&models.AccessPolicyDevice{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
