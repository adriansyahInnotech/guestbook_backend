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

type Device interface {
	GetAll(tracerCtx context.Context, name string, page int) *dto.Response
	Upsert(tracerCtx context.Context, data *dtos.Device) *dto.Response
	Delete(tracerCtx context.Context, id string) *dto.Response
}

type device struct {
	helper              *helper.Helper
	repositoryGuestbook *repository.GuestbookRepository
}

func NewDeviceServices(helper *helper.Helper, repositoryGuestbook *repository.GuestbookRepository) Device {
	return &device{
		helper:              helper,
		repositoryGuestbook: repositoryGuestbook,
	}
}

func (s *device) GetAll(tracerCtx context.Context, name string, page int) *dto.Response {

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_device_services", "get_all")
	defer span.End()

	pagesize := 5

	devicesModel, total, err := s.repositoryGuestbook.DeviceRepository.GetAll(name, page, pagesize)
	if err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed to get all customers")
	}

	totalPages := (total + int64(pagesize) - 1) / int64(pagesize)

	return s.helper.Response.JSONResponseSuccess(devicesModel, int64(page), totalPages, "berhasil")

}

func (s *device) Upsert(tracerCtx context.Context, data *dtos.Device) *dto.Response {
	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_device_services", "upsert")
	defer span.End()

	DeviceModel := new(models.Device)

	if data.ID != "" {
		parsed, err := uuid.Parse(data.ID)
		if err != nil {
			s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
			return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed parsed id")
		}

		DeviceModel.ID = parsed

	}

	DeviceModel.Name = data.Name
	DeviceModel.Location = data.Location
	DeviceModel.IsActive = data.IsActive

	apikey, err := s.helper.Utils.ApiKey.GenerateApiKey()
	if err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed")
	}

	DeviceModel.ApiKey = apikey

	if err := s.repositoryGuestbook.DeviceRepository.Upsert(DeviceModel); err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed")
	}

	return s.helper.Response.JSONResponseSuccess(DeviceModel, 0, 0, "berhasil")
}

func (s *device) Delete(tracerCtx context.Context, id string) *dto.Response {

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_device_services", "delete")
	defer span.End()

	if err := s.repositoryGuestbook.DeviceRepository.Delete(id); err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed")
	}

	return s.helper.Response.JSONResponseSuccess("", 0, 0, "berhasil")

}
