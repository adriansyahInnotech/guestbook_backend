package controllers

import (
	"guestbook_backend/apps/guestbook/services"
	"guestbook_backend/helper"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AccessCardController struct {
	helper     *helper.Helper
	service    services.AccessCard
	validation *validator.Validate
}

func NewAccessCardController(helper *helper.Helper, service services.AccessCard) *AccessCardController {
	return &AccessCardController{
		helper:     helper,
		service:    service,
		validation: validator.New(),
	}
}

func (s *AccessCardController) GetAll(c *fiber.Ctx) error {
	tracerCtx := c.UserContext()

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook.access_card_controllers", "get_all")
	defer span.End()

	card_number := c.Query("card_number")
	page, _ := strconv.Atoi(c.Query("page"))

	data := s.service.GetAll(tracerCtx, card_number, page)
	return c.Status(data.StatusCode).JSON(data)

}
