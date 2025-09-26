package controllers

import (
	"guestbook_backend/apps/guestbook/dtos"
	"guestbook_backend/apps/guestbook/services"
	"guestbook_backend/helper"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type VisitController struct {
	helper     *helper.Helper
	service    services.Visit
	validation *validator.Validate
}

func NewVisitController(helper *helper.Helper, service services.Visit) *VisitController {
	return &VisitController{
		helper:     helper,
		service:    service,
		validation: validator.New(),
	}
}

func (s *VisitController) GetAllWithFilter(c *fiber.Ctx) error {

	tracerCtx := c.UserContext()

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook.visit_controllers", "get_all_with_filter")
	defer span.End()

	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))
	visitorName := c.Query("visitor_name")
	divisionName := c.Query("division_name")
	sectionName := c.Query("section_name")
	checkInFromStr := c.Query("check_in_from")
	checkInToStr := c.Query("check_in_to")

	var checkInFrom, checkInTo *time.Time
	layout := "2006-01-02" // atau "2006-01-02T15:04:05" kalau mau timestamp lengkap
	if checkInFromStr != "" {
		t, err := time.Parse(layout, checkInFromStr)
		if err == nil {
			checkInFrom = &t
		}
	}
	if checkInToStr != "" {
		t, err := time.Parse(layout, checkInToStr)
		if err == nil {
			checkInTo = &t
		}
	}

	filter := dtos.VisitFilter{
		Page:         page,
		PageSize:     pageSize,
		VisitorName:  visitorName,
		DivisionName: divisionName,
		SectionName:  sectionName,
		CheckInFrom:  checkInFrom,
		CheckInTo:    checkInTo,
	}

	data := s.service.GetAllWithFilter(tracerCtx, filter)

	return c.Status(data.StatusCode).JSON(data)
}
