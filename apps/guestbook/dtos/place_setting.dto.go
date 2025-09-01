package dtos

type PlaceSetting struct {
	ID       string `json:"from_id"`
	ToID     string `json:"to_id"`
	PolicyID string `json:"policy_id"`
	Name     string `json:"name" validate:"required"`
	Code     string `json:"code" validate:"required"`
	Address  string `json:"address" validate:"required"`
}
