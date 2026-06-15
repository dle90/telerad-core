package staffAccountControllerRequests

import "github.com/google/uuid"

// StaffAssignReadingPermissionRequest — phân quyền đọc phim cho nhân viên: các modality
// được phép đọc + danh sách đối tác telerad mà nhân viên được đọc phim.
type StaffAssignReadingPermissionRequest struct {
	Modalities          []string    `json:"modalities" validate:"omitempty,dive,required,oneof=CT MR US CR MG"`
	TeleradPartnerUuids []uuid.UUID `json:"teleradPartnerUuids" validate:"omitempty,dive,required"`
}
