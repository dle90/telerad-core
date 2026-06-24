package imagingResultTemplateControllerResponses

import (
	"time"

	"github.com/google/uuid"
)

type StaffGetAImagingResultTemplateResponse struct {
	Uuid         uuid.UUID  `json:"uuid"`
	CreatedAt    time.Time  `json:"createdAt"`
	CreatedBy    uuid.UUID  `json:"createdBy"`
	UpdatedAt    *time.Time `json:"updatedAt"`
	UpdatedBy    *uuid.UUID `json:"updatedBy"`
	Modality     string     `json:"modality"`
	Name         string     `json:"name"`
	BodyParts    []string   `json:"bodyParts"`
	HtmlContent  string     `json:"htmlContent"`
	DisplayOrder *int16     `json:"displayOrder"`
	IsActive     bool       `json:"isActive"`
}
