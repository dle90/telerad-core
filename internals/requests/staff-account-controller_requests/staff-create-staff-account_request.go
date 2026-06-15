package staffAccountControllerRequests

import (
	"telerad-core-module/internals/types"
)

// StaffCreateStaffAccountRequest — tạo hồ sơ nhân viên (chưa có tài khoản đăng nhập).
// username/password được cấp sau qua endpoint "tạo tài khoản"; phân quyền đọc phim
// (modalities + đối tác telerad) cũng cấp sau qua endpoint phân quyền riêng.
type StaffCreateStaffAccountRequest struct {
	Code                  string      `json:"code" validate:"required"`
	FullName              string      `json:"fullName" validate:"required"`
	Gender                string      `json:"gender" validate:"required"`
	DateOfBirth           *types.Date `json:"dateOfBirth"`
	CitizenIdentityNumber *string     `json:"citizenIdentityNumber"`
	Phone                 *string     `json:"phone"`
	Email                 *string     `json:"email"`
	FullAddress           *string     `json:"fullAddress"`
	Roles                 []string    `json:"roles" validate:"omitempty,dive,required,oneof=DOCTOR"`
}
