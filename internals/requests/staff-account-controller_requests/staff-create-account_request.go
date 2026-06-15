package staffAccountControllerRequests

// StaffCreateAccountRequest — cấp tài khoản đăng nhập cho nhân viên chưa có username.
// Mật khẩu được hệ thống tự sinh (trả về 1 lần), client chỉ cung cấp username.
type StaffCreateAccountRequest struct {
	Username string `json:"username" validate:"required"`
}
