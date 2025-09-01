package guestbook

import (
	"guestbook_backend/db"
	"guestbook_backend/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PolicyRepository interface {
	SetDB(db *gorm.DB)
	ClearTransactionDB()
	GetAll(name string, page int, pagesize int) (*[]models.AccessPolicy, int64, error)
	GetOneByName(name string) (*models.AccessPolicy, error)
	Add(policy *models.AccessPolicy) error
	Delete(id string) error
}

type policyRepository struct {
	db *gorm.DB
}

func NewPolicyRepository() PolicyRepository {
	return &policyRepository{
		db: db.GetDB(),
	}
}

// // Override db (misal untuk transaksi)
func (s *policyRepository) SetDB(db *gorm.DB) {
	s.db = db
}

func (s *policyRepository) ClearTransactionDB() {
	s.db = db.GetDB() // reset ke default non-transactional DB
}

func (s *policyRepository) Add(policy *models.AccessPolicy) error {

	result := s.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "id"},
		},
		DoUpdates: clause.AssignmentColumns([]string{
			"name",
		}),
	}).Create(policy)

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (s *policyRepository) GetAll(name string, page int, pagesize int) (*[]models.AccessPolicy, int64, error) {
	var total int64
	policyModel := new([]models.AccessPolicy)

	query := s.db.Model(policyModel)

	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pagesize
	result := query.Preload("Devices").Offset(offset).Limit(pagesize).Find(policyModel)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, total, err
	}

	return policyModel, total, err

}

func (s *policyRepository) GetOneByName(name string) (*models.AccessPolicy, error) {
	accessPolicy := new(models.AccessPolicy)
	result := s.db.Where("name = ?", name).First(accessPolicy)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}

	return accessPolicy, nil
}

func (s *policyRepository) Delete(id string) error {
	result := s.db.Where("id = ?", id).Delete(&models.AccessPolicy{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
