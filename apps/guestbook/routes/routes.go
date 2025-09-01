package routes

import (
	"fmt"
	"guestbook_backend/apps/guestbook/controllers"
	"guestbook_backend/helper"
	"guestbook_backend/middleware"
	"os"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	controllers *controllers.Controllers
	middleware  *middleware.Middleware
}

func NewRoutes() *Routes {
	helper := helper.NewHelper()

	return &Routes{
		controllers: controllers.NewControllers(helper),
		middleware:  middleware.NewMiddlware(helper),
	}
}

func (s *Routes) Routes(app *fiber.App) {

	v1 := app.Group(fmt.Sprintf("%s/v1", os.Getenv("PREFIX_API")))
	v1.Use(otelfiber.Middleware(otelfiber.WithServerName(os.Getenv("SERVICE_NAME"))))

	//devices
	v1.Get("/device/all", s.controllers.Device.GetAll)
	v1.Post("/device/add", s.controllers.Device.Add)
	v1.Delete("/device/delete/:id", s.controllers.Device.Delete)

	//policy
	v1.Post("/policy/add", s.controllers.AccessPolicy.Upsert)
	v1.Get("/policy/all", s.controllers.AccessPolicy.GetAll)
	v1.Delete("/policy/delete/:id", s.controllers.AccessPolicy.Delete)

	//place setting

	v1.Get("/company/all", s.controllers.Company.GetAll)
	v1.Post("/company/add", s.controllers.Company.Add)
	v1.Delete("/company/delete/:id", s.controllers.Company.Delete)

	v1.Get("/division/all", s.controllers.Division.GetAll)
	v1.Post("/division/add", s.controllers.Division.Add)
	v1.Delete("/division/delete/:id", s.controllers.Division.Delete)

	v1.Get("/departement/all", s.controllers.Departement.GetAll)
	v1.Post("/departement/add", s.controllers.Departement.Add)
	v1.Delete("/departement/delete/:id", s.controllers.Departement.Delete)

	v1.Get("/section/all", s.controllers.Section.GetAll)
	v1.Post("/section/add", s.controllers.Section.Add)
	v1.Delete("/section/delete/:id", s.controllers.Section.Delete)

	//devices
	v1.Use(s.middleware.ApiKey.CheckAuthMiddleware())

	//visitor
	v1.Post("/visitor/add", s.controllers.Visitor.Add)

}
