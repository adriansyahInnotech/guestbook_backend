package services

import (
	"context"
	"encoding/json"
	"guestbook_backend/apps/guestbook/dtos"
	"guestbook_backend/helper"
	"guestbook_backend/helper/response/dto"
	"guestbook_backend/models"
	"guestbook_backend/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Visitor interface {
	Upsert(tracerCtx context.Context, data *dtos.Visitor) *dto.Response
}

type visitor struct {
	helper              *helper.Helper
	repositoryGuestbook *repository.GuestbookRepository
}

func NewVisitorServices(helper *helper.Helper, repositoryGuestbook *repository.GuestbookRepository) Visitor {
	return &visitor{
		helper:              helper,
		repositoryGuestbook: repositoryGuestbook,
	}
}

func (s *visitor) Upsert(tracerCtx context.Context, data *dtos.Visitor) *dto.Response {
	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_visitor_services", "add")
	defer span.End()

	visitorModel := &models.Visitor{
		FullName:     data.FullName,
		Company:      data.Company,
		Phone:        data.Phone,
		IDCardType:   data.IDCardType,
		IDCardNumber: data.IDCardNumber,
		DataCard:     data.DataCard,
	}

	if err := s.repositoryGuestbook.VisitorRepository.Upsert(visitorModel); err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed")
	}

	visitor, err := s.repositoryGuestbook.VisitorRepository.GetByIDCardNumber(data.IDCardNumber)
	if err != nil && err != gorm.ErrRecordNotFound {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed")
	}

	byteVisitor, err := json.Marshal(visitor)
	if err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed")
	}

	if data.DeviceID != "" {

		if err := s.helper.Utils.Nats.Publish(data.DeviceID, string(byteVisitor)); err != nil {
			s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
			return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "gagal scan kartu coba lagi")
		}
	}

	return s.helper.Response.JSONResponseSuccess(visitor, 0, 0, "berhasil")

}
