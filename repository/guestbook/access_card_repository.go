package guestbook

import (
	"guestbook_backend/db"
	"guestbook_backend/models"
	"guestbook_backend/repository/redis"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AccessCardRepository interface {
	SetDB(db *gorm.DB)
	ClearTransactionDB()
	Upsert(accessCard *models.AccessCard) error
	GetByCardNumber(access_card string) (*models.AccessCard, error)
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

func (s *accessCardRepository) GetByCardNumber(access_card string) (*models.AccessCard, error) {

	access_cardModel := new(models.AccessCard)

	result := s.db.Where("card_number = ?", access_card).Find(access_cardModel)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}

	return access_cardModel, nil

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
