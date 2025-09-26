package guestbook

import (
	"guestbook_backend/apps/guestbook/dtos"
	"guestbook_backend/db"
	"guestbook_backend/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VisitRepository interface {
	SetDB(db *gorm.DB)
	ClearTransactionDB()
	Add(visit *models.Visit) error
	GetVisitActiveByCard(cardID string) (*models.Visit, error)
	GetAllWithFilter(f dtos.VisitFilter, pageSize int) (*[]models.Visit, int64, error)
	UpdateCheckoutByAccessCard(access_card_id uuid.UUID) error
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

func (s *visitRepository) UpdateCheckoutByAccessCard(access_card_id uuid.UUID) error {

	now := time.Now()

	result := s.db.Model(&models.Visit{}).Where("access_card_id = ? AND check_out IS NULL", access_card_id).
		Updates(map[string]interface{}{
			"check_out":  now,
			"updated_at": now,
		})

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

func (s *visitRepository) GetAllWithFilter(f dtos.VisitFilter, pageSize int) (*[]models.Visit, int64, error) {
	var total int64
	visits := new([]models.Visit)

	query := s.db.Model(&models.Visit{})

	// Dinamis join & where
	if f.VisitorName != "" {
		query = query.Joins("JOIN visitors ON visitors.id = visits.visitor_id").
			Where("visitors.full_name ILIKE ?", "%"+f.VisitorName+"%")
	}

	if f.DivisionName != "" {
		query = query.Joins("JOIN divisions ON divisions.id = visits.division_id").
			Where("divisions.name ILIKE ?", "%"+f.DivisionName+"%")
	}

	if f.DepartmentName != "" {
		query = query.Joins("JOIN departments ON departments.id = visits.department_id").
			Where("departments.name ILIKE ?", "%"+f.DepartmentName+"%")
	}

	if f.SectionName != "" {
		query = query.Joins("JOIN sections ON sections.id = visits.section_id").
			Where("sections.name ILIKE ?", "%"+f.SectionName+"%")
	}

	if f.CheckInFrom != nil {
		query = query.Where("visits.check_in::date >= ?", f.CheckInFrom)
	}

	if f.CheckInTo != nil {
		query = query.Where("visits.check_in::date <= ?", f.CheckInTo)
	}

	// hitung total
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// pagination
	page := f.Page
	if page < 1 {
		page = 1
	}

	offset := (page - 1) * pageSize

	// ambil data dengan preload
	err := query.
		Preload("Visitor").
		Preload("Employee").
		Preload("Company").
		Preload("Division").
		Preload("Department").
		Preload("Section").
		Preload("AccessCard").
		Order("visits.check_in DESC").
		Offset(offset).
		Limit(pageSize).
		Find(visits).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, total, err
	}

	return visits, total, nil
}
