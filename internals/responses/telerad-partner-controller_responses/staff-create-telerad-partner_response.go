package teleradPartnerControllerResponses

import "github.com/google/uuid"

// StaffCreateTeleradPartnerResponse — kết quả tạo đối tác (mật khẩu do client cung
// cấp lúc tạo nên không trả lại).
type StaffCreateTeleradPartnerResponse struct {
	Uuid     uuid.UUID `json:"uuid"`
	Code     string    `json:"code"`
	Name     string    `json:"name"`
	IsActive bool      `json:"isActive"`
	Username string    `json:"username"`
}
