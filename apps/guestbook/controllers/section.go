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

type SectionController struct {
	helper     *helper.Helper
	service    services.Section
	validation *validator.Validate
}

func NewSectionController(helper *helper.Helper, service services.Section) *SectionController {
	return &SectionController{
		helper:     helper,
		service:    service,
		validation: validator.New(),
	}
}

func (s *SectionController) Add(c *fiber.Ctx) error {
	tracerCtx := c.UserContext()

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook.section_controllers", "add")
	defer span.End()

	dto := new(dtos.PlaceSetting)

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

func (s *SectionController) GetAll(c *fiber.Ctx) error {
	tracerCtx := c.UserContext()

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook.section_controllers", "get_all")
	defer span.End()

	name := c.Query("name")
	page, _ := strconv.Atoi(c.Query("page"))
	all := c.Query("all")

	fmt.Println("name : ", name)
	fmt.Println("page :", page)
	data := s.service.GetAll(tracerCtx, name, page, all)
	return c.Status(data.StatusCode).JSON(data)

}

func (s *SectionController) GetByDepartementID(c *fiber.Ctx) error {
	tracerCtx := c.UserContext()

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook.section_controllers", "get_by_section_id")
	defer span.End()

	id := c.Params("id")

	data := s.service.GetByDepartementID(tracerCtx, id)
	return c.Status(data.StatusCode).JSON(data)

}

func (s *SectionController) Delete(c *fiber.Ctx) error {
	tracerCtx := c.UserContext()

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook.section_controllers", "delete")
	defer span.End()

	id := c.Params("id")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(s.helper.Response.JSONResponseError(fiber.StatusBadRequest, "masukan data dengan benar"))
	}

	data := s.service.Delete(tracerCtx, id)
	return c.Status(data.StatusCode).JSON(data)

}
