package guestbook

import (
	"guestbook_backend/db"
	"guestbook_backend/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DivisionRepository interface {
	SetDB(db *gorm.DB)
	ClearTransactionDB()
	Upsert(division *models.Division) error
	GetAll(name string, page int, pagesize int) (*[]models.Division, int64, error)
	Delete(id string) error
}

type divisionRepository struct {
	db *gorm.DB
}

func NewDivisionRepository() DivisionRepository {
	return &divisionRepository{
		db: db.GetDB(),
	}
}

// // Override db (misal untuk transaksi)
func (s *divisionRepository) SetDB(db *gorm.DB) {
	s.db = db
}

func (s *divisionRepository) ClearTransactionDB() {
	s.db = db.GetDB() // reset ke default non-transactional DB
}

func (s *divisionRepository) Upsert(division *models.Division) error {

	return s.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "id"},
		},
		DoUpdates: clause.AssignmentColumns([]string{
			"name",
			"code",
			"company_id",
			"policy_id",
			"updated_at",
		}),
	}).Create(division).Error

}

func (s *divisionRepository) GetAll(name string, page int, pagesize int) (*[]models.Division, int64, error) {
	var total int64
	divisionModel := new([]models.Division)

	query := s.db.Model(divisionModel)

	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pagesize
	result := query.Preload("Company").Preload("Policy").Offset(offset).Limit(pagesize).Find(divisionModel)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, total, err
	}

	return divisionModel, total, err

}

func (s *divisionRepository) Delete(id string) error {
	result := s.db.Where("id = ?", id).Delete(&models.Division{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
