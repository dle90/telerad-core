package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// ImagingResultSheetTemplateEntity — mẫu (bố cục) phiếu kết quả CĐHA, cấu hình theo
// từng CSYT (telerad_partner). html_content là khung phiếu chứa vùng kết quả.
type ImagingResultSheetTemplateEntity struct {
	bun.BaseModel `bun:"table:telerad.imaging_result_sheet_template"`

	Uuid               uuid.UUID  `json:"uuid" bun:"uuid,pk,nullzero" db:"uuid" gorm:"column:uuid;primaryKey"`
	CreatedAt          time.Time  `json:"createdAt" bun:"created_at,notnull,nullzero" db:"created_at" gorm:"column:created_at;not null"`
	CreatedBy          uuid.UUID  `json:"createdBy" bun:"created_by,notnull" db:"created_by" gorm:"column:created_by;not null"`
	UpdatedAt          *time.Time `json:"updatedAt" bun:"updated_at" db:"updated_at" gorm:"column:updated_at"`
	UpdatedBy          *uuid.UUID `json:"updatedBy" bun:"updated_by" db:"updated_by" gorm:"column:updated_by"`
	TeleradPartnerUuid uuid.UUID  `json:"teleradPartnerUuid" bun:"telerad_partner_uuid,notnull" db:"telerad_partner_uuid" gorm:"column:telerad_partner_uuid;not null"`
	HtmlContent        string     `json:"htmlContent" bun:"html_content,notnull" db:"html_content" gorm:"column:html_content;not null"`
	IsActive           bool       `json:"isActive" bun:"is_active,notnull" db:"is_active" gorm:"column:is_active;not null"`

	TeleradPartner *TeleradPartnerEntity `json:"-" bun:"rel:belongs-to,join:telerad_partner_uuid=uuid" gorm:"foreignKey:TeleradPartnerUuid;references:Uuid"`
}

func (ImagingResultSheetTemplateEntity) TableName() string {
	return "telerad.imaging_result_sheet_template"
}
