package guestbook

import (
	"guestbook_backend/apps/guestbook/dtos"
	"guestbook_backend/db"

	"gorm.io/gorm"
)

type DashboardRepository interface {
	SetDB(db *gorm.DB)
	ClearTransactionDB()
	GetAll() (*dtos.DashboardStats, []dtos.PeakHour, error)
}

type dashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository() DashboardRepository {
	return &dashboardRepository{
		db: db.GetDB(),
	}
}

// // Override db (misal untuk transaksi)
func (s *dashboardRepository) SetDB(db *gorm.DB) {
	s.db = db
}

func (s *dashboardRepository) ClearTransactionDB() {
	s.db = db.GetDB() // reset ke default non-transactional DB
}

func (s *dashboardRepository) GetAll() (*dtos.DashboardStats, []dtos.PeakHour, error) {

	dashboardStats := new(dtos.DashboardStats)

	result := s.db.Raw(`
SELECT
  (SELECT COUNT(DISTINCT visitor_id) FROM visits) AS total_unique_visitors,
  (SELECT COUNT(DISTINCT visitor_id) FROM visits WHERE DATE(check_in)=CURRENT_DATE AND check_out IS NULL) AS total_active_today,
  (SELECT COUNT(DISTINCT visitor_id) FROM visits WHERE DATE(check_in)=CURRENT_DATE) AS total_visitors_today
`).Scan(&dashboardStats)

	var peakHours []dtos.PeakHour
	result2 := s.db.Raw(`
SELECT 
    EXTRACT(HOUR FROM check_in)::INT AS hour,
    COUNT(*) AS total_visits
FROM visits
WHERE DATE(check_in) = CURRENT_DATE
GROUP BY hour
ORDER BY hour
`).Scan(&peakHours)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, nil, result.Error
	}

	if result2.Error != nil && result2.Error != gorm.ErrRecordNotFound {
		return nil, nil, result2.Error
	}

	return dashboardStats, peakHours, nil

}
