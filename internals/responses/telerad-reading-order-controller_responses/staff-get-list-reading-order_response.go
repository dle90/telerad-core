package teleradReadingOrderControllerResponses

import (
	"time"

	"github.com/google/uuid"
)

// StaffGetListReadingOrderResponse — 1 dòng danh sách ca đọc ở màn "Đọc ca".
type StaffGetListReadingOrderResponse struct {
	Uuid               uuid.UUID  `json:"uuid"`
	TeleradPartnerUuid uuid.UUID  `json:"teleradPartnerUuid"`
	PartnerName        string     `json:"partnerName"`
	OrderCode          string     `json:"orderCode"`
	OrderItemCode      string     `json:"orderItemCode"`
	StudyInstanceUid   string     `json:"studyInstanceUid"`
	PatientCode        *string    `json:"patientCode"`
	FullName           string     `json:"fullName"`
	Gender             *string    `json:"gender"`
	YearsOld           *int16     `json:"yearsOld"`
	Phone              *string    `json:"phone"`
	ServiceName        string     `json:"serviceName"`
	Modality           *string    `json:"modality"`
	ModalityName       *string    `json:"modalityName"`
	PerformEndedAt     time.Time  `json:"performEndedAt"`
	ReadCompletedAt    *time.Time `json:"readCompletedAt"`
	AssignedTo         *uuid.UUID `json:"assignedTo"`
	AssignedToName     *string    `json:"assignedToName"`
	Status             string     `json:"status"`
}
