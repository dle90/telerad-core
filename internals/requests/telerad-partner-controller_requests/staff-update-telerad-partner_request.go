package teleradPartnerControllerRequests

// StaffUpdateTeleradPartnerRequest — sửa thông tin chung của đối tác.
// Không đụng tới credential phía telerad / cấu hình phía partner / status
// (các phần đó có endpoint riêng).
type StaffUpdateTeleradPartnerRequest struct {
	Code       string   `json:"code" validate:"required"`
	Name       string   `json:"name" validate:"required"`
	Contact    *string  `json:"contact"`
	Modalities []string `json:"modalities"`
}
