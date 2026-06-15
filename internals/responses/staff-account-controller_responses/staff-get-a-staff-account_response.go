package staffAccountControllerResponses

import (
	"time"

	"telerad-core-module/internals/types"

	"github.com/google/uuid"
)

type StaffGetAStaffAccountResponse struct {
	Uuid                  uuid.UUID   `json:"uuid"`
	CreatedAt             time.Time   `json:"createdAt"`
	CreatedBy             *uuid.UUID  `json:"createdBy"`
	UpdatedAt             *time.Time  `json:"updatedAt"`
	UpdatedBy             *uuid.UUID  `json:"updatedBy"`
	Code                  string      `json:"code"`
	FullName              string      `json:"fullName"`
	DateOfBirth           *types.Date `json:"dateOfBirth"`
	Gender                string      `json:"gender"`
	CitizenIdentityNumber *string     `json:"citizenIdentityNumber"`
	Phone                 *string     `json:"phone"`
	Email                 *string     `json:"email"`
	FullAddress           *string     `json:"fullAddress"`
	IsActive              bool        `json:"isActive"`
	Username              *string     `json:"username"`
	Modalities            []string    `json:"modalities"`
	Roles                 []string    `json:"roles"`
	TeleradPartnerUuids   []uuid.UUID `json:"teleradPartnerUuids"`
	//
	HasAccount bool `json:"hasAccount"`
}
