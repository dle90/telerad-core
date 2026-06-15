package teleradPartnerControllerResponses

import (
	"time"

	"github.com/google/uuid"
)

type StaffGetATeleradPartnerResponse struct {
	Uuid            uuid.UUID  `json:"uuid"`
	CreatedAt       time.Time  `json:"createdAt"`
	CreatedBy       uuid.UUID  `json:"createdBy"`
	UpdatedAt       *time.Time `json:"updatedAt"`
	UpdatedBy       *uuid.UUID `json:"updatedBy"`
	Code            string     `json:"code"`
	Name            string     `json:"name"`
	IsActive        bool       `json:"isActive"`
	Contact         *string    `json:"contact"`
	Username        string     `json:"username"`
	Callback        bool       `json:"callback"`
	CallbackUrl     *string    `json:"callbackUrl"`
	PartnerUsername *string    `json:"partnerUsername"`
	Modalities      []string   `json:"modalities"`
}
