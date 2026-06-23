package teleradReadingOrderControllerRequests

import (
	"time"

	"telerad-core-module/internals/types"
)

// PartnerCreateReadingOrderRequest — payload đối tác đẩy 1 ca đọc (reading order) vào telerad.
// telerad_partner_uuid lấy từ JWT của đối tác (không nhận từ body). status do hệ thống set.
type PartnerCreateReadingOrderRequest struct {
	// định danh ca từ phía đối tác
	OrderId       string `json:"orderId" validate:"required"`
	OrderCode     string `json:"orderCode" validate:"required"`
	OrderItemId   string `json:"orderItemId" validate:"required"`
	OrderItemCode string `json:"orderItemCode" validate:"required"`

	StudyInstanceUid string `json:"studyInstanceUid" validate:"required"`

	// thông tin bệnh nhân
	PatientCode *string    `json:"patientCode"`
	FullName    string     `json:"fullName" validate:"required"`
	DateOfBirth types.Date `json:"dateOfBirth" validate:"required"` // yearsOld/monthsOld/daysOld do BE tự tính
	Gender      *string    `json:"gender"`
	Phone       *string    `json:"phone"`
	Email       *string    `json:"email"`
	FullAddress *string    `json:"fullAddress"`

	// dịch vụ / thiết bị
	ServiceId       *string `json:"serviceId"`
	ServiceCode     *string `json:"serviceCode"`
	ServiceName     string  `json:"serviceName" validate:"required"`
	Modality        string  `json:"modality" validate:"required"`
	ModalityAeTitle *string `json:"modalityAeTitle"`
	ModalityCode    *string `json:"modalityCode"`
	ModalityName    *string `json:"modalityName"`

	// lâm sàng
	Note              *string   `json:"note"`
	PerformEndedAt    time.Time `json:"performEndedAt" validate:"required"`
	ClinicalDiagnosis *string   `json:"clinicalDiagnosis"`
	Icd               []string  `json:"icd" validate:"omitempty,dive,required"`

	// bộ phận chụp (PACS_BODY_PART); HIS hiện chưa gửi -> có thể null
	BodyParts []string `json:"bodyParts" validate:"omitempty,dive,required"`
}
