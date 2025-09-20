package guestbook

import (
	"guestbook_backend/db"
	"guestbook_backend/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DeviceRepository interface {
	SetDB(db *gorm.DB)
	ClearTransactionDB()
	GetAll(name string, page int, pagesize int) (*[]models.Device, int64, error)
	Upsert(device *models.Device) error
	GetApiKey(name string) (*models.Device, error)
	BatchUpdatePolicyDevices(deviceIds []uuid.UUID, policyId uuid.UUID) error
	Delete(id string) error
}

type deviceRepository struct {
	db *gorm.DB
}

func NewDeviceRepository() DeviceRepository {
	return &deviceRepository{
		db: db.GetDB(),
	}
}

// // Override db (misal untuk transaksi)
func (s *deviceRepository) SetDB(db *gorm.DB) {
	s.db = db
}

func (s *deviceRepository) ClearTransactionDB() {
	s.db = db.GetDB() // reset ke default non-transactional DB
}

func (s *deviceRepository) GetAll(name string, page int, pagesize int) (*[]models.Device, int64, error) {
	var total int64
	deviceModel := new([]models.Device)

	query := s.db.Model(deviceModel)

	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pagesize
	result := query.Preload("AccessPolicies").Offset(offset).Limit(pagesize).Find(deviceModel)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, total, err
	}

	return deviceModel, total, err

}

func (s *deviceRepository) GetApiKey(name string) (*models.Device, error) {
	deviceModel := new(models.Device)

	result := s.db.Where("name = ? ", name).Find(deviceModel)

	if result.Error != nil {
		return nil, result.Error
	}

	return deviceModel, nil

}

func (s *deviceRepository) Upsert(device *models.Device) error {

	return s.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "id"},
		},
		DoUpdates: clause.AssignmentColumns([]string{
			"name",
			"api_key",
			"location",
			"updated_at",
		}),
	}).Create(device).Error

}

func (s *deviceRepository) BatchUpdatePolicyDevices(deviceIds []uuid.UUID, policyId uuid.UUID) error {

	result := s.db.Model(&models.Device{}).Where("id IN ?", deviceIds).Update("policy_id", policyId)

	if result.Error != nil {
		return result.Error
	}

	if err := s.db.Model(&models.Device{}).
		Where("policy_id = ?", policyId).
		Where("id NOT IN ?", deviceIds).
		Update("policy_id", nil).Error; err != nil {
		return err
	}

	return nil

}

func (s *deviceRepository) Delete(id string) error {
	result := s.db.Where("id = ?", id).Delete(&models.Device{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
