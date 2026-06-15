package staffAccountControllerResponses

import "github.com/google/uuid"

// StaffAccountCredentialResponse — kết quả tạo tài khoản / reset mật khẩu.
// password là plaintext tự sinh, chỉ trả về 1 lần.
type StaffAccountCredentialResponse struct {
	Uuid     uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}
