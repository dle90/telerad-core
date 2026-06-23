package imagingResultTemplateControllerResponses

import (
	"time"

	"github.com/google/uuid"
)

type StaffGetListImagingResultTemplateResponse struct {
	Uuid         uuid.UUID `json:"uuid"`
	Modality     string    `json:"modality"`
	Name         string    `json:"name"`
	BodyParts    []string  `json:"bodyParts"`
	FontSize     int       `json:"fontSize"`
	LineSpacing  float64   `json:"lineSpacing"`
	DisplayOrder *int16    `json:"displayOrder"`
	IsActive     bool      `json:"isActive"`
	CreatedAt    time.Time `json:"createdAt"`
}
