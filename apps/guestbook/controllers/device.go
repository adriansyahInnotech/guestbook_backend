package controllers

import (
	"fmt"
	"guestbook_backend/apps/guestbook/dtos"
	"guestbook_backend/apps/guestbook/services"
	"guestbook_backend/helper"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type DeviceController struct {
	helper     *helper.Helper
	service    services.Device
	validation *validator.Validate
}

func NewDeviceController(helper *helper.Helper, service services.Device) *DeviceController {
	return &DeviceController{
		helper:     helper,
		service:    service,
		validation: validator.New(),
	}
}

func (s *DeviceController) Add(c *fiber.Ctx) error {
	tracerCtx := c.UserContext()

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_device_controllers", "add")
	defer span.End()

	dto := new(dtos.Device)

	if err := c.BodyParser(dto); err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return c.Status(fiber.StatusBadRequest).JSON(s.helper.Response.JSONResponseError(fiber.StatusBadRequest, "gagal parsing body"))
	}

	if err := s.validation.Struct(dto); err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return c.Status(fiber.StatusBadRequest).JSON(s.helper.Response.JSONResponseError(fiber.StatusBadRequest, "masukan data dengan benar"))
	}

	data := s.service.Upsert(tracerCtx, dto)

	return c.Status(data.StatusCode).JSON(data)
}

func (s *DeviceController) GetAll(c *fiber.Ctx) error {
	tracerCtx := c.UserContext()

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_device_controllers", "get_all")
	defer span.End()

	name := c.Query("name")
	page, _ := strconv.Atoi(c.Query("page"))

	fmt.Println("name : ", name)
	fmt.Println("page :", page)
	data := s.service.GetAll(tracerCtx, name, page)
	return c.Status(data.StatusCode).JSON(data)

}

func (s *DeviceController) Delete(c *fiber.Ctx) error {
	tracerCtx := c.UserContext()

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_device_controllers", "delete")
	defer span.End()

	id := c.Params("id")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(s.helper.Response.JSONResponseError(fiber.StatusBadRequest, "masukan data dengan benar"))
	}

	data := s.service.Delete(tracerCtx, id)
	return c.Status(data.StatusCode).JSON(data)

}

func (s *DeviceController) ValidateCard(c *fiber.Ctx) error {
	tracerCtx := c.UserContext()

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_device_controllers", "validate_card")
	defer span.End()

	cardNumber := c.Params("card_number")
	deviceID := c.Params("device_id")

	if cardNumber == "" {
		return c.Status(fiber.StatusBadRequest).JSON(s.helper.Response.JSONResponseError(fiber.StatusBadRequest, "masukan data dengan benar"))
	}

	if deviceID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(s.helper.Response.JSONResponseError(fiber.StatusBadRequest, "masukan data dengan benar"))
	}

	data := s.service.ValidateCard(tracerCtx, cardNumber, deviceID)
	return c.Status(data.StatusCode).JSON(data)

}
