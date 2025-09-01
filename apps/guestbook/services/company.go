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

type Company interface {
	Upsert(tracerCtx context.Context, data *dtos.PlaceSetting) *dto.Response
	GetAll(tracerCtx context.Context, name string, page int) *dto.Response
	Delete(tracerCtx context.Context, id string) *dto.Response
}

type company struct {
	helper              *helper.Helper
	repositoryGuestbook *repository.GuestbookRepository
}

func NewCompanyServices(helper *helper.Helper, repositoryGuestbook *repository.GuestbookRepository) Company {
	return &company{
		helper:              helper,
		repositoryGuestbook: repositoryGuestbook,
	}
}

func (s *company) Upsert(tracerCtx context.Context, data *dtos.PlaceSetting) *dto.Response {
	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_company_services", "upsert")
	defer span.End()

	CompanyModel := new(models.Company)

	if data.ID != "" {
		parsed, err := uuid.Parse(data.ID)
		if err != nil {
			s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
			return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed parsed id")
		}

		CompanyModel.ID = parsed

	}

	CompanyModel.Name = data.Name
	CompanyModel.Code = data.Code
	CompanyModel.Address = data.Address

	if err := s.repositoryGuestbook.CompanyRepository.Upsert(CompanyModel); err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed upsert company")
	}

	return s.helper.Response.JSONResponseSuccess("", 0, 0, "berhasil")
}

func (s *company) GetAll(tracerCtx context.Context, name string, page int) *dto.Response {

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_company_services", "get_all")
	defer span.End()

	pagesize := 5

	companyModel, total, err := s.repositoryGuestbook.CompanyRepository.GetAll(name, page, pagesize)
	if err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed to get all company")
	}

	totalPages := (total + int64(pagesize) - 1) / int64(pagesize)

	return s.helper.Response.JSONResponseSuccess(companyModel, int64(page), totalPages, "berhasil")

}

func (s *company) Delete(tracerCtx context.Context, id string) *dto.Response {

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_company_services", "delete")
	defer span.End()

	if err := s.repositoryGuestbook.CompanyRepository.Delete(id); err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed")
	}

	return s.helper.Response.JSONResponseSuccess("", 0, 0, "berhasil")

}
