package guestbook

import (
	"guestbook_backend/db"
	"guestbook_backend/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SectionRepository interface {
	SetDB(db *gorm.DB)
	ClearTransactionDB()
	Upsert(division *models.Section) error
	GetAll(name string, page int, pagesize int, all string) (*[]models.Section, int64, error)
	GetByDepartmentID(id string) (*[]models.Section, error)
	Delete(id string) error
}

type sectionRepository struct {
	db *gorm.DB
}

func NewSectionRepository() SectionRepository {
	return &sectionRepository{
		db: db.GetDB(),
	}
}

// // Override db (misal untuk transaksi)
func (s *sectionRepository) SetDB(db *gorm.DB) {
	s.db = db
}

func (s *sectionRepository) ClearTransactionDB() {
	s.db = db.GetDB() // reset ke default non-transactional DB
}

func (s *sectionRepository) Upsert(division *models.Section) error {

	return s.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "id"},
		},
		DoUpdates: clause.AssignmentColumns([]string{
			"name",
			"code",
			"department_id",
			"policy_id",
			"updated_at",
		}),
	}).Create(division).Error

}

func (s *sectionRepository) GetAll(name string, page int, pagesize int, all string) (*[]models.Section, int64, error) {
	if all == "true" {
		var total int64
		sectionModel := new([]models.Section)

		query := s.db.Model(sectionModel)

		err := query.Count(&total).Error
		if err != nil {
			return nil, 0, err
		}

		result := query.Find(sectionModel)

		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			return nil, total, err
		}

		return sectionModel, total, err

	} else {
		var total int64
		sectionModel := new([]models.Section)

		query := s.db.Model(sectionModel)

		if name != "" {
			query = query.Where("name ILIKE ?", "%"+name+"%")
		}

		err := query.Count(&total).Error
		if err != nil {
			return nil, 0, err
		}

		offset := page * pagesize
		result := query.Preload("Department").Preload("Policy").Order("created_at desc").Offset(offset).Limit(pagesize).Order("created_at desc").Find(sectionModel)

		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			return nil, total, err
		}

		return sectionModel, total, err
	}

}

func (s *sectionRepository) GetByDepartmentID(id string) (*[]models.Section, error) {
	sectionModel := new([]models.Section)

	result := s.db.Where("department_id = ?", id).Find(sectionModel)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}

	return sectionModel, nil
}

func (s *sectionRepository) Delete(id string) error {
	result := s.db.Where("id = ?", id).Delete(&models.Section{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
