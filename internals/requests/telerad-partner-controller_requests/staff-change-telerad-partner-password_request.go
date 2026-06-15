package teleradPartnerControllerRequests

// StaffChangeTeleradPartnerPasswordRequest — đổi mật khẩu phía telerad của đối tác.
// Mật khẩu mới do admin cung cấp (không cho đổi username).
type StaffChangeTeleradPartnerPasswordRequest struct {
	Password string `json:"password" validate:"required"`
}
