package staffAccountControllerRequests

// UserChangePasswordRequest — user tự đổi mật khẩu (cần mật khẩu hiện tại).
type UserChangePasswordRequest struct {
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}
