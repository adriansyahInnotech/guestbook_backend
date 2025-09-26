package controllers

import (
	"guestbook_backend/apps/guestbook/services"
	"guestbook_backend/helper"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type DashboardController struct {
	helper     *helper.Helper
	service    services.Dashboard
	validation *validator.Validate
}

func NewDashboardController(helper *helper.Helper, service services.Dashboard) *DashboardController {
	return &DashboardController{
		helper:     helper,
		service:    service,
		validation: validator.New(),
	}
}

func (s *DashboardController) GetAll(c *fiber.Ctx) error {
	tracerCtx := c.UserContext()

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook.Company_controllers", "get_all")
	defer span.End()

	data := s.service.GetAllDataDashboard(tracerCtx)
	return c.Status(data.StatusCode).JSON(data)

}
