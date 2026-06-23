package teleradReadingOrderControllerResponses

import (
	"time"

	"telerad-core-module/internals/types"

	"github.com/google/uuid"
)

// StaffGetReadingOrderDetailResponse — chi tiết 1 ca đọc cho tab chi tiết màn "Đọc ca".
type StaffGetReadingOrderDetailResponse struct {
	Uuid               uuid.UUID `json:"uuid"`
	TeleradPartnerUuid uuid.UUID `json:"teleradPartnerUuid"`
	PartnerName        string    `json:"partnerName"`

	OrderCode        string `json:"orderCode"`
	OrderItemCode    string `json:"orderItemCode"`
	StudyInstanceUid string `json:"studyInstanceUid"`

	// bệnh nhân
	PatientCode *string     `json:"patientCode"`
	FullName    string      `json:"fullName"`
	DateOfBirth *types.Date `json:"dateOfBirth"`
	Gender      *string     `json:"gender"`
	YearsOld    *int16      `json:"yearsOld"`
	MonthsOld   *int16      `json:"monthsOld"`
	DaysOld     *int        `json:"daysOld"`
	Phone       *string     `json:"phone"`
	Email       *string     `json:"email"`
	FullAddress *string     `json:"fullAddress"`

	// dịch vụ / thiết bị
	ServiceName  string  `json:"serviceName"`
	Modality     *string `json:"modality"`
	ModalityName *string `json:"modalityName"`

	// lâm sàng
	Note              *string   `json:"note"`
	PerformEndedAt    time.Time `json:"performEndedAt"`
	ClinicalDiagnosis *string   `json:"clinicalDiagnosis"`
	Icd               []string  `json:"icd"`
	BodyParts         []string  `json:"bodyParts"`

	// đọc / duyệt
	AssignedTo      *uuid.UUID `json:"assignedTo"`
	AssignedToName  *string    `json:"assignedToName"`
	AssignedToMe    bool       `json:"assignedToMe"` // ca đang được CHÍNH user gọi API đọc
	ReadCompletedAt *time.Time `json:"readCompletedAt"`
	Status          string     `json:"status"`
	ResultReturned  bool       `json:"resultReturned"`
	ResultInHtml    *string    `json:"resultInHtml"` // nội dung kết quả đã nhập (html)
}
