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

type Department interface {
	Upsert(tracerCtx context.Context, data *dtos.PlaceSetting) *dto.Response
	GetAll(tracerCtx context.Context, name string, page int) *dto.Response
	GetByDivisionID(tracerCtx context.Context, id string) *dto.Response
	Delete(tracerCtx context.Context, id string) *dto.Response
}

type department struct {
	helper              *helper.Helper
	repositoryGuestbook *repository.GuestbookRepository
}

func NewDepartmentServices(helper *helper.Helper, repositoryGuestbook *repository.GuestbookRepository) Department {
	return &department{
		helper:              helper,
		repositoryGuestbook: repositoryGuestbook,
	}
}

func (s *department) Upsert(tracerCtx context.Context, data *dtos.PlaceSetting) *dto.Response {
	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_Department_services", "upsert")
	defer span.End()

	DepartmentModel := new(models.Department)

	if data.ID != "" {
		parsed, err := uuid.Parse(data.ID)
		if err != nil {
			s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
			return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed parsed id")
		}

		DepartmentModel.ID = parsed

	}

	if data.ToID != "" {
		parsed, err := uuid.Parse(data.ToID)
		if err != nil {
			s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
			return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed parsed to id")
		}

		DepartmentModel.DivisionID = &parsed
	}

	if data.PolicyID != "" {
		parsed, err := uuid.Parse(data.PolicyID)
		if err != nil {
			s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
			return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed parsed policy id")
		}

		DepartmentModel.PolicyID = &parsed
	}

	DepartmentModel.Name = data.Name
	DepartmentModel.Code = data.Code

	if err := s.repositoryGuestbook.DepartmentRepository.Upsert(DepartmentModel); err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed upsert Department")
	}

	return s.helper.Response.JSONResponseSuccess("", 0, 0, "berhasil")
}

func (s *department) GetAll(tracerCtx context.Context, name string, page int) *dto.Response {

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_Department_services", "get_all")
	defer span.End()

	pagesize := 5

	// fmt.Println("\n\n\n\n\n masuk ke Department")

	DepartmentModel, total, err := s.repositoryGuestbook.DepartmentRepository.GetAll(name, page, pagesize)
	if err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed to get all Department")
	}

	totalPages := (total + int64(pagesize) - 1) / int64(pagesize)

	return s.helper.Response.JSONResponseSuccess(DepartmentModel, int64(page), totalPages, "berhasil")

}

func (s *department) GetByDivisionID(tracerCtx context.Context, id string) *dto.Response {

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_division_services", "get_by_division_id")
	defer span.End()

	departementModel, err := s.repositoryGuestbook.DepartmentRepository.GetByDivisionID(id)
	if err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed to get_by_division_id")
	}

	return s.helper.Response.JSONResponseSuccess(departementModel, 0, 0, "berhasil")

}

func (s *department) Delete(tracerCtx context.Context, id string) *dto.Response {

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_Department_services", "delete")
	defer span.End()

	if err := s.repositoryGuestbook.DepartmentRepository.Delete(id); err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed")
	}

	return s.helper.Response.JSONResponseSuccess("", 0, 0, "berhasil")

}
