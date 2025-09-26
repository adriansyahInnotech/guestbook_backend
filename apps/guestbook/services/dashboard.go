package services

import (
	"context"
	"guestbook_backend/apps/guestbook/dtos"
	"guestbook_backend/helper"
	"guestbook_backend/helper/response/dto"
	"guestbook_backend/repository"

	"github.com/gofiber/fiber/v2"
)

type Dashboard interface {
	GetAllDataDashboard(tracerCtx context.Context) *dto.Response
}

type dashboard struct {
	helper              *helper.Helper
	repositoryGuestbook *repository.GuestbookRepository
}

func NewDashboardService(helper *helper.Helper, repositoryGuestbook *repository.GuestbookRepository) Dashboard {
	return &dashboard{
		helper:              helper,
		repositoryGuestbook: repositoryGuestbook,
	}
}

func (s *dashboard) GetAllDataDashboard(tracerCtx context.Context) *dto.Response {

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_Department_services", "upsert")
	defer span.End()

	dashboardstat, peakhour, err := s.repositoryGuestbook.DashboardRepository.GetAll()
	if err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed parsed id")
	}

	response := dtos.DashboardResponse{
		DashboardStatus: *dashboardstat,
		PeakHour:        peakhour,
	}

	return s.helper.Response.JSONResponseSuccess(response, 0, 0, "berhasil")
}
