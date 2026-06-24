package databaseQueryModels

import (
	"time"

	"github.com/google/uuid"
)

// ReadingOrderListRow — 1 dòng worklist màn "Đọc ca": các cột của
// telerad_reading_order + tên đối tác (join telerad_partner) + tên bác sĩ đọc
// (left join staff_account theo assigned_to).
type ReadingOrderListRow struct {
	Uuid               uuid.UUID  `bun:"uuid"`
	TeleradPartnerUuid uuid.UUID  `bun:"telerad_partner_uuid"`
	OrderCode          string     `bun:"order_code"`
	OrderItemCode      string     `bun:"order_item_code"`
	StudyInstanceUid   string     `bun:"study_instance_uid"`
	PatientCode        *string    `bun:"patient_code"`
	FullName           string     `bun:"full_name"`
	Gender             *string    `bun:"gender"`
	YearsOld           *int16     `bun:"years_old"`
	Phone              *string    `bun:"phone"`
	ServiceName        string     `bun:"service_name"`
	Modality           *string    `bun:"modality"`
	ModalityName       *string    `bun:"modality_name"`
	PerformEndedAt     time.Time  `bun:"perform_ended_at"`
	ReadCompletedAt    *time.Time `bun:"read_completed_at"`
	AssignedTo         *uuid.UUID `bun:"assigned_to"`
	Status             string     `bun:"status"`
	ResultReturned     bool       `bun:"result_returned"`

	// Joined fields.
	PartnerName    string  `bun:"partner_name"`
	AssignedToName *string `bun:"assigned_to_name"`
}
