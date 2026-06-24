package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// ImagingResultTemplateEntity — mẫu nội dung kết quả CĐHA (findings) theo modality +
// bộ phận chụp (body_parts = mã PACS_BODY_PART). html_content là nội dung mẫu chèn vào
// vùng kết quả của phiếu.
type ImagingResultTemplateEntity struct {
	bun.BaseModel `bun:"table:telerad.imaging_result_template"`

	Uuid         uuid.UUID  `json:"uuid" bun:"uuid,pk,nullzero" db:"uuid" gorm:"column:uuid;primaryKey"`
	CreatedAt    time.Time  `json:"createdAt" bun:"created_at,notnull,nullzero" db:"created_at" gorm:"column:created_at;not null"`
	CreatedBy    uuid.UUID  `json:"createdBy" bun:"created_by,notnull" db:"created_by" gorm:"column:created_by;not null"`
	UpdatedAt    *time.Time `json:"updatedAt" bun:"updated_at" db:"updated_at" gorm:"column:updated_at"`
	UpdatedBy    *uuid.UUID `json:"updatedBy" bun:"updated_by" db:"updated_by" gorm:"column:updated_by"`
	Modality     string     `json:"modality" bun:"modality,notnull" db:"modality" gorm:"column:modality;not null"`
	Name         string     `json:"name" bun:"name,notnull" db:"name" gorm:"column:name;not null"`
	BodyParts    []string   `json:"bodyParts" bun:"body_parts,array" db:"body_parts" gorm:"column:body_parts"`
	HtmlContent  string     `json:"htmlContent" bun:"html_content,notnull" db:"html_content" gorm:"column:html_content;not null"`
	DisplayOrder *int16     `json:"displayOrder" bun:"display_order" db:"display_order" gorm:"column:display_order"`
	IsActive     bool       `json:"isActive" bun:"is_active,notnull" db:"is_active" gorm:"column:is_active;not null"`
}

func (ImagingResultTemplateEntity) TableName() string {
	return "telerad.imaging_result_template"
}
