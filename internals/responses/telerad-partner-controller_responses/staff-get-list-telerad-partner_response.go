package teleradPartnerControllerResponses

import (
	"time"

	"github.com/google/uuid"
)

type StaffGetListTeleradPartnerResponse struct {
	Uuid       uuid.UUID `json:"uuid"`
	Code       string    `json:"code"`
	Name       string    `json:"name"`
	IsActive   bool      `json:"isActive"`
	Username   string    `json:"username"`
	Contact    *string   `json:"contact"`
	Callback   bool      `json:"callback"`
	Modalities []string  `json:"modalities"`
	CreatedAt  time.Time `json:"createdAt"`
}
