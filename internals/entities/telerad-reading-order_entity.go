package entities

import (
	"time"

	"telerad-core-module/internals/types"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Trạng thái ca đọc (cột status).
const (
	TeleradReadingOrderStatusPending   = "PENDING"   // mới nhận từ partner, chờ phân công bác sĩ
	TeleradReadingOrderStatusAssigned  = "ASSIGNED"  // đã gán bác sĩ đọc
	TeleradReadingOrderStatusCompleted = "COMPLETED" // đã đọc xong
)

type TeleradReadingOrderEntity struct {
	bun.BaseModel `bun:"table:telerad.telerad_reading_order"`

	Uuid               uuid.UUID   `json:"uuid" bun:"uuid,pk,nullzero" db:"uuid" gorm:"column:uuid;primaryKey"`
	CreatedAt          time.Time   `json:"createdAt" bun:"created_at,notnull,nullzero" db:"created_at" gorm:"column:created_at;not null"`
	CreatedBy          uuid.UUID   `json:"createdBy" bun:"created_by,notnull" db:"created_by" gorm:"column:created_by;not null"`
	UpdatedAt          *time.Time  `json:"updatedAt" bun:"updated_at" db:"updated_at" gorm:"column:updated_at"`
	UpdatedBy          *uuid.UUID  `json:"updatedBy" bun:"updated_by" db:"updated_by" gorm:"column:updated_by"`
	TeleradPartnerUuid uuid.UUID   `json:"teleradPartnerUuid" bun:"telerad_partner_uuid,notnull" db:"telerad_partner_uuid" gorm:"column:telerad_partner_uuid;not null"`
	OrderId            string      `json:"orderId" bun:"order_id,notnull" db:"order_id" gorm:"column:order_id;not null"`
	OrderCode          string      `json:"orderCode" bun:"order_code,notnull" db:"order_code" gorm:"column:order_code;not null"`
	OrderItemId        string      `json:"orderItemId" bun:"order_item_id,notnull" db:"order_item_id" gorm:"column:order_item_id;not null"`
	OrderItemCode      string      `json:"orderItemCode" bun:"order_item_code,notnull" db:"order_item_code" gorm:"column:order_item_code;not null"`
	StudyInstanceUid   string      `json:"studyInstanceUid" bun:"study_instance_uid,notnull" db:"study_instance_uid" gorm:"column:study_instance_uid;not null"`
	PatientCode        *string     `json:"patientCode" bun:"patient_code" db:"patient_code" gorm:"column:patient_code"`
	FullName           string      `json:"fullName" bun:"full_name,notnull" db:"full_name" gorm:"column:full_name;not null"`
	DateOfBirth        *types.Date `json:"dateOfBirth" bun:"date_of_birth" db:"date_of_birth" gorm:"column:date_of_birth"`
	Gender             *string     `json:"gender" bun:"gender" db:"gender" gorm:"column:gender"`
	Phone              *string     `json:"phone" bun:"phone" db:"phone" gorm:"column:phone"`
	Email              *string     `json:"email" bun:"email" db:"email" gorm:"column:email"`
	FullAddress        *string     `json:"fullAddress" bun:"full_address" db:"full_address" gorm:"column:full_address"`
	YearsOld           *int16      `json:"yearsOld" bun:"years_old" db:"years_old" gorm:"column:years_old"`
	MonthsOld          *int16      `json:"monthsOld" bun:"months_old" db:"months_old" gorm:"column:months_old"`
	DaysOld            *int        `json:"daysOld" bun:"days_old" db:"days_old" gorm:"column:days_old"`
	ServiceId          *string     `json:"serviceId" bun:"service_id" db:"service_id" gorm:"column:service_id"`
	ServiceCode        *string     `json:"serviceCode" bun:"service_code" db:"service_code" gorm:"column:service_code"`
	ServiceName        string      `json:"serviceName" bun:"service_name,notnull" db:"service_name" gorm:"column:service_name;not null"`
	Modality           *string     `json:"modality" bun:"modality" db:"modality" gorm:"column:modality"`
	ModalityAeTitle    *string     `json:"modalityAeTitle" bun:"modality_ae_title" db:"modality_ae_title" gorm:"column:modality_ae_title"`
	ModalityCode       *string     `json:"modalityCode" bun:"modality_code" db:"modality_code" gorm:"column:modality_code"`
	ModalityName       *string     `json:"modalityName" bun:"modality_name" db:"modality_name" gorm:"column:modality_name"`
	Note               *string     `json:"note" bun:"note" db:"note" gorm:"column:note"`
	PerformEndedAt     time.Time   `json:"performEndedAt" bun:"perform_ended_at,notnull" db:"perform_ended_at" gorm:"column:perform_ended_at;not null"`
	ClinicalDiagnosis  *string     `json:"clinicalDiagnosis" bun:"clinical_diagnosis" db:"clinical_diagnosis" gorm:"column:clinical_diagnosis"`
	Icd                []string    `json:"icd" bun:"icd,array" db:"icd" gorm:"column:icd"`
	AssignedAt         *time.Time  `json:"assignedAt" bun:"assigned_at" db:"assigned_at" gorm:"column:assigned_at"`
	AssignedTo         *uuid.UUID  `json:"assignedTo" bun:"assigned_to" db:"assigned_to" gorm:"column:assigned_to"`
	ReadCompletedAt    *time.Time  `json:"readCompletedAt" bun:"read_completed_at" db:"read_completed_at" gorm:"column:read_completed_at"`
	Status             string      `json:"status" bun:"status,notnull" db:"status" gorm:"column:status;not null"`
}

func (TeleradReadingOrderEntity) TableName() string { return "telerad.telerad_reading_order" }
