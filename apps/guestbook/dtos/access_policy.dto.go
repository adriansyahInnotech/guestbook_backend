package dtos

import "github.com/google/uuid"

type AccessPolicy struct {
	ID        string      `json:"id" `
	Name      string      `json:"name" validate:"required"`
	DeviceIDs []uuid.UUID `json:"device_ids" validate:"required,dive,uuid"`
}
