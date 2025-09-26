package guestbook

import (
	"guestbook_backend/db"
	"guestbook_backend/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type VisitorRepository interface {
	SetDB(db *gorm.DB)
	ClearTransactionDB()

	Upsert(visitor *models.Visitor) error
	GetByIDCardNumber(idcard string) (*models.Visitor, error)
	GetAll(name string, page int, pagesize int) (*[]models.Visitor, int64, error)
}

type visitorRepository struct {
	db *gorm.DB
}

func NewVisitorRepository() VisitorRepository {
	return &visitorRepository{
		db: db.GetDB(),
	}
}

// // Override db (misal untuk transaksi)
func (s *visitorRepository) SetDB(db *gorm.DB) {
	s.db = db
}

func (s *visitorRepository) ClearTransactionDB() {
	s.db = db.GetDB() // reset ke default non-transactional DB
}

func (s *visitorRepository) Upsert(visitor *models.Visitor) error {

	return s.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id_card_number"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"full_name", "company", "phone", "id_card_type", "data_card", "updated_at",
		}),
	}).Create(visitor).Error

}

func (s *visitorRepository) GetByIDCardNumber(idcard string) (*models.Visitor, error) {

	visitorModel := new(models.Visitor)

	result := s.db.Where("id_card_number = ?", idcard).First(visitorModel)

	if result.Error != nil {
		return nil, result.Error
	}

	return visitorModel, nil
}

func (s *visitorRepository) GetAll(name string, page int, pagesize int) (*[]models.Visitor, int64, error) {
	var total int64
	visitorModel := new([]models.Visitor)

	query := s.db.Model(visitorModel)

	if name != "" {
		query = query.Where("full_name ILIKE ?", "%"+name+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pagesize
	result := query.Offset(offset).Limit(pagesize).Find(visitorModel)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, total, err
	}

	return visitorModel, total, err

}
