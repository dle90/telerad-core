package teleradPartnerControllerResponses

import "github.com/google/uuid"

// StaffGetAllTeleradPartnerResponse — phần tử danh sách chọn (combobox) đối tác telerad:
// hiển thị TẤT CẢ đối tác kèm trạng thái để user tick chọn khi phân quyền đọc phim.
type StaffGetAllTeleradPartnerResponse struct {
	Uuid       uuid.UUID `json:"uuid"`
	Code       string    `json:"code"`
	Name       string    `json:"name"`
	IsActive   bool      `json:"isActive"`
	Modalities []string  `json:"modalities"`
}
