package services

import (
	"context"
	"guestbook_backend/apps/guestbook/dtos"
	"guestbook_backend/helper"
	"guestbook_backend/helper/response/dto"
	"guestbook_backend/repository"

	"github.com/gofiber/fiber/v2"
)

type Visit interface {
	GetAllWithFilter(tracerCtx context.Context, filter dtos.VisitFilter) *dto.Response
}

type visit struct {
	helper              *helper.Helper
	repositoryGuestbook *repository.GuestbookRepository
}

func NewVisitServices(helper *helper.Helper, repositoryGuestbook *repository.GuestbookRepository) Visit {
	return &visit{
		helper:              helper,
		repositoryGuestbook: repositoryGuestbook,
	}
}

func (s *visit) GetAllWithFilter(tracerCtx context.Context, filter dtos.VisitFilter) *dto.Response {

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_visit_services", "get_all")
	defer span.End()

	pagesize := 5

	visitModel, total, err := s.repositoryGuestbook.VisitRepository.GetAllWithFilter(filter, pagesize)
	if err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed to get all Section")
	}

	totalPages := (total + int64(pagesize) - 1) / int64(pagesize)

	return s.helper.Response.JSONResponseSuccess(visitModel, int64(filter.Page), totalPages, "berhasil")

}
