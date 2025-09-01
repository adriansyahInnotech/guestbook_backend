package services

import (
	"context"
	"guestbook_backend/apps/guestbook/dtos"
	"guestbook_backend/db"
	"guestbook_backend/helper"
	"guestbook_backend/helper/response/dto"
	"guestbook_backend/models"
	"guestbook_backend/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AccessPolicy interface {
	GetAll(tracerCtx context.Context, name string, page int) *dto.Response
	Upsert(tracerCtx context.Context, data *dtos.AccessPolicy) *dto.Response
	Delete(tracerCtx context.Context, id string) *dto.Response
}

type accessPolicy struct {
	helper              *helper.Helper
	repositoryGuestbook *repository.GuestbookRepository
}

func NewAccessPolicyServices(helper *helper.Helper, repositoryGuestbook *repository.GuestbookRepository) AccessPolicy {
	return &accessPolicy{
		helper:              helper,
		repositoryGuestbook: repositoryGuestbook,
	}
}

func (s *accessPolicy) GetAll(tracerCtx context.Context, name string, page int) *dto.Response {

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_access_policy_services", "get_all")
	defer span.End()

	pagesize := 5

	policyModel, total, err := s.repositoryGuestbook.AccessPolicyRepository.GetAll(name, page, pagesize)
	if err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed to get all policy")
	}

	totalPages := (total + int64(pagesize) - 1) / int64(pagesize)

	return s.helper.Response.JSONResponseSuccess(policyModel, int64(page), totalPages, "berhasil")

}

func (s *accessPolicy) Upsert(tracerCtx context.Context, data *dtos.AccessPolicy) *dto.Response {

	tx := db.GetDB().Begin()
	s.repositoryGuestbook.SetDB(tx)

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_policy_services", "upsert")
	defer span.End()

	newPolicyModel := new(models.AccessPolicy)

	if data.ID != "" {
		parsed, err := uuid.Parse(data.ID)
		if err != nil {
			tx.Rollback()
			s.repositoryGuestbook.ClearTransactionDB()
			s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
			return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed parsed id")
		}

		newPolicyModel.ID = parsed

	}

	newPolicyModel.Name = data.Name

	if err := s.repositoryGuestbook.AccessPolicyRepository.Add(newPolicyModel); err != nil {
		tx.Rollback()
		s.repositoryGuestbook.ClearTransactionDB()
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed add policy")
	}

	if err := s.repositoryGuestbook.DeviceRepository.BatchUpdatePolicyDevices(data.DeviceIDs, newPolicyModel.ID); err != nil {
		tx.Rollback()
		s.repositoryGuestbook.ClearTransactionDB()
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed update policy device")
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		s.repositoryGuestbook.ClearTransactionDB()
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed commit accsess policy")
	}

	s.repositoryGuestbook.ClearTransactionDB()

	return s.helper.Response.JSONResponseSuccess("", 0, 0, "berhasil")

}

func (s *accessPolicy) Delete(tracerCtx context.Context, id string) *dto.Response {

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_policy_services", "delete")
	defer span.End()

	if err := s.repositoryGuestbook.AccessPolicyRepository.Delete(id); err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed")
	}

	return s.helper.Response.JSONResponseSuccess("", 0, 0, "berhasil")

}
