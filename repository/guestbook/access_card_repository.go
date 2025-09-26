package guestbook

import (
	"guestbook_backend/db"
	"guestbook_backend/models"
	"guestbook_backend/repository/redis"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AccessCardRepository interface {
	SetDB(db *gorm.DB)
	ClearTransactionDB()
	Upsert(accessCard *models.AccessCard) error
	ReleaseCard(cardID uuid.UUID) error
	GetAccessCardByCardNumber(card_number string) (*models.AccessCard, error)
	GetAll(name string, page int, pagesize int) (*[]models.AccessCard, int64, error)
	SyncCardToRedis(cardID uuid.UUID) error
}

type accessCardRepository struct {
	db              *gorm.DB
	redisRepository redis.RedisRepository
}

func NewAccessCardRepository() AccessCardRepository {
	return &accessCardRepository{
		db:              db.GetDB(),
		redisRepository: redis.NewRedisRepository(),
	}
}

// // Override db (misal untuk transaksi)
func (s *accessCardRepository) SetDB(db *gorm.DB) {
	s.db = db
}

func (s *accessCardRepository) ClearTransactionDB() {
	s.db = db.GetDB() // reset ke default non-transactional DB
}

func (s *accessCardRepository) Upsert(accessCard *models.AccessCard) error {

	return s.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "card_number"}}, // unique key
		DoUpdates: clause.AssignmentColumns([]string{"visitor_id", "policy_id", "updated_at"}),
	}).Create(&accessCard).Error

}

func (s *accessCardRepository) ReleaseCard(cardID uuid.UUID) error {
	return s.db.Model(&models.AccessCard{}).
		Where("id = ?", cardID).
		Updates(map[string]interface{}{
			"visitor_id": nil, // atau null-kan visitor_id
			"updated_at": time.Now(),
		}).Error
}

func (s *accessCardRepository) GetAccessCardByCardNumber(card_number string) (*models.AccessCard, error) {
	accessCardModel := new(models.AccessCard)

	result := s.db.Where("card_number = ?", card_number).First(accessCardModel)

	if result.Error != nil {
		return nil, result.Error
	}

	return accessCardModel, nil

}

func (s *accessCardRepository) GetAll(card_number string, page int, pagesize int) (*[]models.AccessCard, int64, error) {
	var total int64
	accessCardModel := new([]models.AccessCard)

	query := s.db.Model(accessCardModel)

	if card_number != "" {
		query = query.Where("card_number ILIKE ?", "%"+card_number+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pagesize
	result := query.Preload("Visitor").Offset(offset).Limit(pagesize).Find(accessCardModel)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, total, err
	}

	return accessCardModel, total, err

}

func (s *accessCardRepository) SyncCardToRedis(cardID uuid.UUID) error {
	var card models.AccessCard
	if err := s.db.
		Preload("Policy.Devices").
		First(&card, "id = ?", cardID).Error; err != nil {
		return err
	}

	var deviceIDs []string

	for _, d := range card.Policy.Devices {
		deviceIDs = append(deviceIDs, d.DeviceID.String())
	}

	return s.redisRepository.SetCardDevices(card.CardNumber, deviceIDs)
}
