package dtos

type Visitor struct {
	DeviceID     string `json:"device_id"`
	FullName     string `json:"full_name" validate:"required"`
	Company      string `json:"company"`
	Phone        string `json:"phone" validate:"required"`
	IDCardType   string `json:"id_card_type" validate:"required"`
	IDCardNumber string `json:"id_card_number" validate:"required"`
	DataCard     string `json:"data_card" validate:"required"`
}
