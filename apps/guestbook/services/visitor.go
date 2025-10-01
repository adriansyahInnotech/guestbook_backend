package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"guestbook_backend/apps/guestbook/dtos"
	"guestbook_backend/db"
	"guestbook_backend/helper"
	"guestbook_backend/helper/response/dto"
	"guestbook_backend/models"
	"guestbook_backend/repository"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Visitor interface {
	GetAll(tracerCtx context.Context, name string, page int) *dto.Response
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

func (s *visitor) GetAll(tracerCtx context.Context, name string, page int) *dto.Response {

	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_visitor_services", "get_all")
	defer span.End()

	pagesize := 5

	visitorModel, total, err := s.repositoryGuestbook.VisitorRepository.GetAll(name, page, pagesize)
	if err != nil {
		s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
		return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed to get all Section")
	}

	totalPages := (total + int64(pagesize) - 1) / int64(pagesize)

	return s.helper.Response.JSONResponseSuccess(visitorModel, int64(page), totalPages, "berhasil")

}

func (s *visitor) Upsert(tracerCtx context.Context, data *dtos.Visitor) *dto.Response {
	_, span := s.helper.Utils.JaegerTracer.StartSpan(tracerCtx, "guestbook_visitor_services", "add")
	defer span.End()

	switch data.TypeInput {
	case "scan":

		if err := s.scanFromDevice(data); err != nil {
			s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
			return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, "failed scan dari perangkat")
		}

		return s.helper.Response.JSONResponseSuccess("", 0, 0, "berhasil mengirim data ke web")

	case "manual":

		if err := s.manualInput(data); err != nil {
			s.helper.Utils.JaegerTracer.RecordSpanError(span, err)
			return s.helper.Response.JSONResponseError(fiber.StatusInternalServerError, err.Error())
		}

		return s.helper.Response.JSONResponseSuccess("", 0, 0, "berhasil assign kartu")

	}

	return s.helper.Response.JSONResponseError(fiber.StatusBadRequest, "type input tidak valid")

}

func (s *visitor) scanFromDevice(data *dtos.Visitor) error {

	visitorModel := &models.Visitor{
		FullName:     data.FullName,
		Company:      data.Company,
		Phone:        data.Phone,
		IDCardType:   data.IDCardType,
		IDCardNumber: data.IDCardNumber,
		DataCard:     datatypes.JSON([]byte(data.DataCard)),
	}

	if err := s.repositoryGuestbook.VisitorRepository.Upsert(visitorModel); err != nil {
		return err
	}

	visitor, err := s.repositoryGuestbook.VisitorRepository.GetByIDCardNumber(data.IDCardNumber)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	byteVisitor, err := json.Marshal(visitor)
	if err != nil {
		return err
	}

	if data.DeviceID != "" {

		if err := s.helper.Utils.Nats.Publish(data.DeviceID, string(byteVisitor)); err != nil {
			return err
		}
	}

	return nil

}

func (s *visitor) manualInput(data *dtos.Visitor) error {

	tx := db.GetDB().Begin()
	s.repositoryGuestbook.SetDB(tx)

	if data.PolicyID == "" {
		tx.Rollback()
		s.repositoryGuestbook.ClearTransactionDB()
		return errors.New("tidak ada policy")
	}

	visitorModel := &models.Visitor{
		FullName:     data.FullName,
		Company:      data.Company,
		Phone:        data.Phone,
		IDCardType:   data.IDCardType,
		IDCardNumber: data.IDCardNumber,
		DataCard:     datatypes.JSON([]byte(data.DataCard)),
	}

	if err := s.repositoryGuestbook.VisitorRepository.Upsert(visitorModel); err != nil {
		tx.Rollback()
		s.repositoryGuestbook.ClearTransactionDB()
		return err
	}

	visitor, err := s.repositoryGuestbook.VisitorRepository.GetByIDCardNumber(data.IDCardNumber)
	if err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		s.repositoryGuestbook.ClearTransactionDB()
		return err
	}

	accessCardModel := new(models.AccessCard)
	visitModel := new(models.Visit)

	visitModel.Notes = data.Notes
	visitModel.CheckIn = time.Now()
	visitModel.VisitorID = visitor.ID

	if data.DivisionID != "" {
		parsed, err := uuid.Parse(data.DivisionID)
		if err != nil {
			fmt.Println("masuk error division_id")
			tx.Rollback()
			s.repositoryGuestbook.ClearTransactionDB()
			return err
		}
		visitModel.DivisionID = &parsed
	}

	if data.DepartementID != "" {
		parsed, err := uuid.Parse(data.DepartementID)
		if err != nil {
			tx.Rollback()
			s.repositoryGuestbook.ClearTransactionDB()
			fmt.Println("masuk error departement_id")
			return err
		}
		visitModel.DepartmentID = &parsed
	}

	if data.SectionID != "" {
		parsed, err := uuid.Parse(data.SectionID)
		if err != nil {
			tx.Rollback()
			s.repositoryGuestbook.ClearTransactionDB()
			fmt.Println("masuk error section_id")
			return err
		}
		visitModel.SectionID = &parsed
	}

	if data.DeviceID != "" {
		parsed, err := uuid.Parse(data.DeviceID)
		if err != nil {
			tx.Rollback()
			s.repositoryGuestbook.ClearTransactionDB()
			return err
		}
		visitModel.DeviceID = &parsed
	}

	parsed, err := uuid.Parse(data.PolicyID)
	if err != nil {
		tx.Rollback()
		s.repositoryGuestbook.ClearTransactionDB()
		return err
	}
	accessCardModel.PolicyID = &parsed

	//check card active
	activeVisit, err := s.repositoryGuestbook.VisitRepository.GetVisitActiveByCard(data.AccessCard)
	if err != nil {
		tx.Rollback()
		s.repositoryGuestbook.ClearTransactionDB()
		return err
	}

	if activeVisit != nil {
		tx.Rollback()
		s.repositoryGuestbook.ClearTransactionDB()
		return errors.New("kartu tidak bisa di assign user belom checkout")
	}

	accessCardModel.VisitorID = visitor.ID
	accessCardModel.CardNumber = data.AccessCard

	if err := s.repositoryGuestbook.AccessCardRepository.Upsert(accessCardModel); err != nil {
		log.Println("\n\n\n error upsert accesscard to redis : ", err.Error())
		tx.Rollback()
		s.repositoryGuestbook.ClearTransactionDB()
		return err
	}

	if err := s.repositoryGuestbook.AccessCardRepository.SyncCardToRedis(accessCardModel.ID); err != nil {
		log.Println("\n\n\n error sync card to redis : ", err.Error())
		tx.Rollback()
		s.repositoryGuestbook.ClearTransactionDB()
		return err
	}

	visitModel.AccessCardID = &accessCardModel.ID

	if err := s.repositoryGuestbook.VisitRepository.Add(visitModel); err != nil {
		tx.Rollback()
		s.repositoryGuestbook.ClearTransactionDB()
		return err

	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		s.repositoryGuestbook.ClearTransactionDB()

		return err
	}

	s.repositoryGuestbook.ClearTransactionDB()

	return nil

}
