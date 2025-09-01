package controllers

import (
	"guestbook_backend/apps/guestbook/services"
	"guestbook_backend/helper"
	"guestbook_backend/repository"
)

type Controllers struct {
	Visitor      *VisitorController
	Device       *DeviceController
	AccessPolicy *AccessPolicyController
	Company      *CompanyController
	Division     *DivisionController
	Departement  *DepartementController
	Section      *SectionController
}

func NewControllers(helper *helper.Helper) *Controllers {

	guestbookRepository := repository.NewGuestbookRepository()
	allService := services.NewServices(helper, guestbookRepository)

	return &Controllers{
		Visitor:      NewVisitorController(helper, allService.Visitor),
		Device:       NewDeviceController(helper, allService.Device),
		AccessPolicy: NewAccessPolicyController(helper, allService.AccessPolicy),
		Company:      NewCompanyController(helper, allService.Company),
		Division:     NewDivisionController(helper, allService.Division),
		Departement:  NewDepartementController(helper, allService.Department),
		Section:      NewSectionController(helper, allService.Section),
	}
}
