package teleradPartnerControllerResponses

import "github.com/google/uuid"

// StaffTeleradPartnerPartnerConfigResponse — xem cấu hình tài khoản phía partner.
// partnerPassword là plaintext (telerad lưu để gọi sang hệ thống đối tác) nên
// được trả ra cho admin xem.
type StaffTeleradPartnerPartnerConfigResponse struct {
	Uuid            uuid.UUID `json:"uuid"`
	Callback        bool      `json:"callback"`
	CallbackUrl     *string   `json:"callbackUrl"`
	PartnerUsername *string   `json:"partnerUsername"`
	PartnerPassword *string   `json:"partnerPassword"`
}
