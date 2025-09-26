package services

import (
	"guestbook_backend/helper"
	"guestbook_backend/repository"
)

type Services struct {
	Visitor      Visitor
	Device       Device
	AccessPolicy AccessPolicy
	Company      Company
	Division     Division
	Department   Department
	Section      Section
	Visit        Visit
	AccessCard   AccessCard
	Dashboard    Dashboard
}

func NewServices(helper *helper.Helper, repositoryGuestbook *repository.GuestbookRepository) *Services {
	return &Services{
		Visitor:      NewVisitorServices(helper, repositoryGuestbook),
		Device:       NewDeviceServices(helper, repositoryGuestbook),
		AccessPolicy: NewAccessPolicyServices(helper, repositoryGuestbook),
		Company:      NewCompanyServices(helper, repositoryGuestbook),
		Division:     NewDivisionServices(helper, repositoryGuestbook),
		Department:   NewDepartmentServices(helper, repositoryGuestbook),
		Section:      NewSectionServices(helper, repositoryGuestbook),
		Visit:        NewVisitServices(helper, repositoryGuestbook),
		AccessCard:   NewAccessCardServices(helper, repositoryGuestbook),
		Dashboard:    NewDashboardService(helper, repositoryGuestbook),
	}
}
