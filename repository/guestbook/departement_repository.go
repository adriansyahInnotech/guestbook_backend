package guestbook

import (
	"guestbook_backend/db"
	"guestbook_backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DepartmentRepository interface {
	SetDB(db *gorm.DB)
	ClearTransactionDB()
	Upsert(division *models.Department) error
	GetAll(name string, page int, pagesize int) (*[]models.Department, int64, error)
	Delete(id string) error
	GetByDivisionID(id string) (*[]models.Department, error)
	GetByID(id uuid.UUID) (*models.Department, error)
}

type departmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository() DepartmentRepository {
	return &departmentRepository{
		db: db.GetDB(),
	}
}

// // Override db (misal untuk transaksi)
func (s *departmentRepository) SetDB(db *gorm.DB) {
	s.db = db
}

func (s *departmentRepository) ClearTransactionDB() {
	s.db = db.GetDB() // reset ke default non-transactional DB
}

func (s *departmentRepository) Upsert(division *models.Department) error {

	return s.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "id"},
		},
		DoUpdates: clause.AssignmentColumns([]string{
			"name",
			"code",
			"division_id",
			"policy_id",
			"updated_at",
		}),
	}).Create(division).Error

}

func (s *departmentRepository) GetAll(name string, page int, pagesize int) (*[]models.Department, int64, error) {
	var total int64
	departmentModel := new([]models.Department)

	query := s.db.Model(departmentModel)

	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pagesize
	result := query.Preload("Sections").Preload("Division").Preload("Policy").Offset(offset).Limit(pagesize).Find(departmentModel)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, total, err
	}

	return departmentModel, total, err

}

func (s *departmentRepository) GetByDivisionID(id string) (*[]models.Department, error) {
	departmentModel := new([]models.Department)

	result := s.db.Where("division_id = ?", id).Find(departmentModel)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}

	return departmentModel, nil
}

func (s *departmentRepository) GetByID(id uuid.UUID) (*models.Department, error) {

	departementModel := new(models.Department)

	result := s.db.Where("id = ?", id).First(departementModel)

	if result.Error != nil {
		return nil, result.Error
	}

	return departementModel, nil

}

func (s *departmentRepository) Delete(id string) error {
	result := s.db.Where("id = ?", id).Delete(&models.Department{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
