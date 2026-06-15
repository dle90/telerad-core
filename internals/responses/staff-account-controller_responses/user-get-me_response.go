package staffAccountControllerResponses

import (
	"telerad-core-module/internals/types"

	"github.com/google/uuid"
)

// UserGetMeResponse — thông tin bản thân của user đang đăng nhập.
type UserGetMeResponse struct {
	Uuid                  uuid.UUID   `json:"uuid"`
	Code                  string      `json:"code"`
	FullName              string      `json:"fullName"`
	DateOfBirth           *types.Date `json:"dateOfBirth"`
	Gender                string      `json:"gender"`
	CitizenIdentityNumber *string     `json:"citizenIdentityNumber"`
	Phone                 *string     `json:"phone"`
	Email                 *string     `json:"email"`
	FullAddress           *string     `json:"fullAddress"`
	Username              *string     `json:"username"`
	IsActive              bool        `json:"isActive"`
	Modalities            []string    `json:"modalities"`
	Roles                 []string    `json:"roles"`
	TeleradPartnerUuids   []uuid.UUID `json:"teleradPartnerUuids"`
}
