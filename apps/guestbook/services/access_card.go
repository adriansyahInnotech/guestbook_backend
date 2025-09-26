package services

import (
	"context"
	"guestbook_backend/helper"
	"guestbook_backend/helper/response/dto"
	"guestbook_backend/repository"

	"github.com/gofiber/fiber/v2"
)

type AccessCard interface {
	GetAll(tracerCtx context.Context, card_number string, page int) *dto.Response
}

type accessCard struct {
	helper              *helper.Helper
	repositoryGuestbook *repository.GuestbookRepository
}

func NewAccessCardServices(helper *helper.Helper, repositoryGuestbook *repository.GuestbookRepository) AccessCard {
	return &accessCard{
		helper:              helper,
		repositoryGuestbook: repositoryGuestbook,
	}
}

func (s *accessCard) GetAll(tracerCtx context.Context, card_number string, page int) *dto.Response {

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_access_card_services", "get_all")
	defer span.End()

	pagesize := 5

	accessCardModel, total, err := s.repositoryGuestbook.AccessCardRepository.GetAll(card_number, page, pagesize)
	if err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed to get all access_card")
	}

	totalPages := (total + int64(pagesize) - 1) / int64(pagesize)

	return s.helper.Response.JSONResponseSuccess(accessCardModel, int64(page), totalPages, "berhasil")

}
