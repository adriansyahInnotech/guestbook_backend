package guestbook

import (
	"guestbook_backend/db"
	"guestbook_backend/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CompanyRepository interface {
	SetDB(db *gorm.DB)
	ClearTransactionDB()
	GetAll(name string, page int, pagesize int) (*[]models.Company, int64, error)
	Upsert(company *models.Company) error
	Delete(id string) error
}

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository() CompanyRepository {
	return &companyRepository{
		db: db.GetDB(),
	}
}

// // Override db (misal untuk transaksi)
func (s *companyRepository) SetDB(db *gorm.DB) {
	s.db = db
}

func (s *companyRepository) ClearTransactionDB() {
	s.db = db.GetDB() // reset ke default non-transactional DB
}

func (s *companyRepository) Upsert(company *models.Company) error {

	return s.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "id"},
		},
		DoUpdates: clause.AssignmentColumns([]string{
			"name",
			"code",
			"address",
			"updated_at",
		}),
	}).Create(company).Error

}

func (s *companyRepository) GetAll(name string, page int, pagesize int) (*[]models.Company, int64, error) {
	var total int64
	CompanyModel := new([]models.Company)

	query := s.db.Model(CompanyModel)

	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pagesize
	result := query.Preload("Divisions").Offset(offset).Limit(pagesize).Find(CompanyModel)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, total, err
	}

	return CompanyModel, total, err

}

func (s *companyRepository) Delete(id string) error {
	result := s.db.Where("id = ?", id).Delete(&models.Company{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
