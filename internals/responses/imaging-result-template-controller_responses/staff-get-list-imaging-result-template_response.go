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
	DisplayOrder *int16    `json:"displayOrder"`
	IsActive     bool      `json:"isActive"`
	CreatedAt    time.Time `json:"createdAt"`
}
