package teleradPartnerControllerRequests

// StaffCreateTeleradPartnerRequest — tạo đối tác telerad mới. Mật khẩu phía telerad
// (Password) do client cung cấp. Nếu Callback = true thì PartnerUsername & PartnerPassword
// là bắt buộc (telerad cần credential để gọi callback sang hệ thống đối tác).
type StaffCreateTeleradPartnerRequest struct {
	Code     string  `json:"code" validate:"required"`
	Name     string  `json:"name" validate:"required"`
	Username string  `json:"username" validate:"required"` // tên đăng nhập phía telerad cấp cho đối tác
	Password string  `json:"password" validate:"required"` // mật khẩu phía telerad
	Contact  *string `json:"contact"`
	// cấu hình tài khoản phía partner
	Callback        bool     `json:"callback"`
	CallbackUrl     *string  `json:"callbackUrl"`
	PartnerUsername *string  `json:"partnerUsername" validate:"required_if=Callback true"`
	PartnerPassword *string  `json:"partnerPassword" validate:"required_if=Callback true"`
	Modalities      []string `json:"modalities" validate:"omitempty,dive,required,oneof=CT MR US CR MG"` // danh sách các modality mà partner này cung cấp dịch vụ
}
