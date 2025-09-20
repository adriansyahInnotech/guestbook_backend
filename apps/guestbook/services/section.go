package services

import (
	"context"
	"guestbook_backend/apps/guestbook/dtos"
	"guestbook_backend/helper"
	"guestbook_backend/helper/response/dto"
	"guestbook_backend/models"
	"guestbook_backend/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Section interface {
	Upsert(tracerCtx context.Context, data *dtos.PlaceSetting) *dto.Response
	GetAll(tracerCtx context.Context, name string, page int) *dto.Response
	GetByDepartementID(tracerCtx context.Context, id string) *dto.Response
	Delete(tracerCtx context.Context, id string) *dto.Response
}

type section struct {
	helper              *helper.Helper
	repositoryGuestbook *repository.GuestbookRepository
}

func NewSectionServices(helper *helper.Helper, repositoryGuestbook *repository.GuestbookRepository) Section {
	return &section{
		helper:              helper,
		repositoryGuestbook: repositoryGuestbook,
	}
}

func (s *section) Upsert(tracerCtx context.Context, data *dtos.PlaceSetting) *dto.Response {
	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_Section_services", "upsert")
	defer span.End()

	SectionModel := new(models.Section)

	if data.ID != "" {
		parsed, err := uuid.Parse(data.ID)
		if err != nil {
			s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
			return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed parsed id")
		}

		SectionModel.ID = parsed

	}

	if data.ToID != "" {
		parsed, err := uuid.Parse(data.ToID)
		if err != nil {
			s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
			return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed parsed to id")
		}

		SectionModel.DepartmentID = &parsed
	}

	if data.PolicyID != "" {
		parsed, err := uuid.Parse(data.PolicyID)
		if err != nil {
			s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
			return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed parsed policy id")
		}

		SectionModel.PolicyID = &parsed
	}

	SectionModel.Name = data.Name
	SectionModel.Code = data.Code

	if err := s.repositoryGuestbook.SectionRepository.Upsert(SectionModel); err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed upsert Section")
	}

	return s.helper.Response.JSONResponseSuccess("", 0, 0, "berhasil")
}

func (s *section) GetAll(tracerCtx context.Context, name string, page int) *dto.Response {

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_Section_services", "get_all")
	defer span.End()

	pagesize := 5

	// fmt.Println("\n\n\n\n\n masuk ke Section")

	SectionModel, total, err := s.repositoryGuestbook.SectionRepository.GetAll(name, page, pagesize)
	if err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed to get all Section")
	}

	totalPages := (total + int64(pagesize) - 1) / int64(pagesize)

	return s.helper.Response.JSONResponseSuccess(SectionModel, int64(page), totalPages, "berhasil")

}

func (s *section) Delete(tracerCtx context.Context, id string) *dto.Response {

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_Section_services", "delete")
	defer span.End()

	if err := s.repositoryGuestbook.SectionRepository.Delete(id); err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed")
	}

	return s.helper.Response.JSONResponseSuccess("", 0, 0, "berhasil")

}

func (s *section) GetByDepartementID(tracerCtx context.Context, id string) *dto.Response {

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_section_services", "get_by_section_id")
	defer span.End()

	departementModel, err := s.repositoryGuestbook.SectionRepository.GetByDepartmentID(id)
	if err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed to get_by_section_id")
	}

	return s.helper.Response.JSONResponseSuccess(departementModel, 0, 0, "berhasil")

}
