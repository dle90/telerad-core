package staffAccountControllerRequests

import (
	"telerad-core-module/internals/types"
)

// StaffUpdateStaffAccountRequest — sửa hồ sơ nhân viên. Không đụng tới username/password/
// is_active, phân quyền đọc phim (modalities + đối tác), roles — đều có endpoint riêng.
type StaffUpdateStaffAccountRequest struct {
	Code                  string      `json:"code" validate:"required"`
	FullName              string      `json:"fullName" validate:"required"`
	Gender                string      `json:"gender" validate:"required,oneof=MALE FEMALE"`
	DateOfBirth           *types.Date `json:"dateOfBirth"`
	CitizenIdentityNumber *string     `json:"citizenIdentityNumber"`
	Phone                 *string     `json:"phone"`
	Email                 *string     `json:"email"`
	FullAddress           *string     `json:"fullAddress"`
}
