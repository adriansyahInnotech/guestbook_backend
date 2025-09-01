package dtos

type Device struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	IsActive bool   `json:"is_active"`
}
