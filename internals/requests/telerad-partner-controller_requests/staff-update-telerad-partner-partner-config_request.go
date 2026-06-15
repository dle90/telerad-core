package teleradPartnerControllerRequests

// StaffUpdateTeleradPartnerPartnerConfigRequest — cấu hình tài khoản phía partner:
// thông tin telerad dùng để callback / gọi sang hệ thống của đối tác.
type StaffUpdateTeleradPartnerPartnerConfigRequest struct {
	Callback        bool    `json:"callback"`
	CallbackUrl     *string `json:"callbackUrl"`
	PartnerUsername *string `json:"partnerUsername" validate:"required_if=Callback true"`
	PartnerPassword *string `json:"partnerPassword" validate:"required_if=Callback true"`
}
