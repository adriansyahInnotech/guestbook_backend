package repository

import (
	guestbook "guestbook_backend/repository/guestbook"

	"gorm.io/gorm"
)

type TransactionalRepository interface {
	SetDB(tx *gorm.DB)
	ClearTransactionDB()
}

type GuestbookRepository struct {
	AccessPolicyRepository       guestbook.PolicyRepository
	VisitorRepository            guestbook.VisitorRepository
	VisitRepository              guestbook.VisitRepository
	DeviceRepository             guestbook.DeviceRepository
	CompanyRepository            guestbook.CompanyRepository
	DivisionRepository           guestbook.DivisionRepository
	DepartmentRepository         guestbook.DepartmentRepository
	SectionRepository            guestbook.SectionRepository
	AccessCardRepository         guestbook.AccessCardRepository
	AccessPolicyDeviceRepository guestbook.PolicyDeviceRepository
	DashboardRepository          guestbook.DashboardRepository

	all []TransactionalRepository
}

func NewGuestbookRepository() *GuestbookRepository {

	r := &GuestbookRepository{
		AccessPolicyRepository:       guestbook.NewPolicyRepository(),
		VisitorRepository:            guestbook.NewVisitorRepository(),
		VisitRepository:              guestbook.NewVisitRepository(),
		DeviceRepository:             guestbook.NewDeviceRepository(),
		CompanyRepository:            guestbook.NewCompanyRepository(),
		DivisionRepository:           guestbook.NewDivisionRepository(),
		DepartmentRepository:         guestbook.NewDepartmentRepository(),
		SectionRepository:            guestbook.NewSectionRepository(),
		AccessCardRepository:         guestbook.NewAccessCardRepository(),
		AccessPolicyDeviceRepository: guestbook.NewPolicyDeviceRepository(),
		DashboardRepository:          guestbook.NewDashboardRepository(),
	}

	// Kumpulkan semua repository yang support transaksi
	r.all = []TransactionalRepository{
		r.AccessPolicyRepository,
		r.VisitorRepository,
		r.VisitRepository,
		r.DeviceRepository,
		r.CompanyRepository,
		r.DivisionRepository,
		r.DepartmentRepository,
		r.SectionRepository,
		r.AccessCardRepository,
		r.AccessPolicyDeviceRepository,
		r.DashboardRepository,
	}

	return r
}

func (r *GuestbookRepository) SetDB(tx *gorm.DB) {
	for _, repo := range r.all {
		repo.SetDB(tx)
	}
}

func (r *GuestbookRepository) ClearTransactionDB() {
	for _, repo := range r.all {
		repo.ClearTransactionDB()
	}
}
