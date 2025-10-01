package services

import (
	"context"
	"fmt"
	"guestbook_backend/apps/guestbook/dtos"
	"guestbook_backend/helper"
	"guestbook_backend/helper/response/dto"
	"guestbook_backend/models"
	"guestbook_backend/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Division interface {
	Upsert(tracerCtx context.Context, data *dtos.PlaceSetting) *dto.Response
	GetAll(tracerCtx context.Context, name string, page int, all string) *dto.Response
	Delete(tracerCtx context.Context, id string) *dto.Response
}

type division struct {
	helper              *helper.Helper
	repositoryGuestbook *repository.GuestbookRepository
}

func NewDivisionServices(helper *helper.Helper, repositoryGuestbook *repository.GuestbookRepository) Division {
	return &division{
		helper:              helper,
		repositoryGuestbook: repositoryGuestbook,
	}
}

func (s *division) Upsert(tracerCtx context.Context, data *dtos.PlaceSetting) *dto.Response {
	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_division_services", "upsert")
	defer span.End()

	divisionModel := new(models.Division)

	if data.ID != "" {
		parsed, err := uuid.Parse(data.ID)
		if err != nil {
			s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
			return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed parsed id")
		}

		divisionModel.ID = parsed

	}

	if data.ToID != "" {
		parsed, err := uuid.Parse(data.ToID)
		if err != nil {
			s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
			return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed parsed to id")
		}

		divisionModel.CompanyID = &parsed
	}

	if data.PolicyID != "" {
		parsed, err := uuid.Parse(data.PolicyID)
		if err != nil {
			s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
			return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed parsed policy id")
		}

		divisionModel.PolicyID = &parsed
	}

	divisionModel.Name = data.Name
	divisionModel.Code = data.Code

	if err := s.repositoryGuestbook.DivisionRepository.Upsert(divisionModel); err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed upsert division")
	}

	return s.helper.Response.JSONResponseSuccess("", 0, 0, "berhasil")
}

func (s *division) GetAll(tracerCtx context.Context, name string, page int, all string) *dto.Response {

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_division_services", "get_all")
	defer span.End()

	pagesize := 5

	fmt.Println("\n\n\n\n\n masuk ke division")

	divisionModel, total, err := s.repositoryGuestbook.DivisionRepository.GetAll(name, page, pagesize, all)
	if err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed to get all division")
	}

	totalPages := (total + int64(pagesize) - 1) / int64(pagesize)

	return s.helper.Response.JSONResponseSuccess(divisionModel, int64(page), totalPages, "berhasil")

}

func (s *division) Delete(tracerCtx context.Context, id string) *dto.Response {

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_division_services", "delete")
	defer span.End()

	if err := s.repositoryGuestbook.DivisionRepository.Delete(id); err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed")
	}

	return s.helper.Response.JSONResponseSuccess("", 0, 0, "berhasil")

}
