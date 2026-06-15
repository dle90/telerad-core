package staffAccountControllerResponses

import (
	"time"

	"github.com/google/uuid"
)

type StaffGetListStaffAccountResponse struct {
	Uuid       uuid.UUID `json:"uuid"`
	Code       string    `json:"code"`
	FullName   string    `json:"fullName"`
	Gender     string    `json:"gender"`
	Phone      *string   `json:"phone"`
	Email      *string   `json:"email"`
	Username   *string   `json:"username"`
	IsActive   bool      `json:"isActive"`
	Modalities []string  `json:"modalities"`
	Roles      []string  `json:"roles"`
	CreatedAt  time.Time `json:"createdAt"`
	//
	HasAccount bool `json:"hasAccount"`
}
