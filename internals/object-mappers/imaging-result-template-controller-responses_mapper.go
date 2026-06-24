package objectMappers

import (
	"telerad-core-module/internals/entities"
	imagingResultTemplateControllerResponses "telerad-core-module/internals/responses/imaging-result-template-controller_responses"
)

func ToStaffGetAImagingResultTemplateResponse(template entities.ImagingResultTemplateEntity) imagingResultTemplateControllerResponses.StaffGetAImagingResultTemplateResponse {
	return imagingResultTemplateControllerResponses.StaffGetAImagingResultTemplateResponse{
		Uuid:         template.Uuid,
		CreatedAt:    template.CreatedAt,
		CreatedBy:    template.CreatedBy,
		UpdatedAt:    template.UpdatedAt,
		UpdatedBy:    template.UpdatedBy,
		Modality:     template.Modality,
		Name:         template.Name,
		BodyParts:    emptyIfNil(template.BodyParts),
		HtmlContent:  template.HtmlContent,
		DisplayOrder: template.DisplayOrder,
		IsActive:     template.IsActive,
	}
}

func ToStaffGetListImagingResultTemplateSlice(templates []entities.ImagingResultTemplateEntity) []imagingResultTemplateControllerResponses.StaffGetListImagingResultTemplateResponse {
	result := make([]imagingResultTemplateControllerResponses.StaffGetListImagingResultTemplateResponse, 0, len(templates))

	for _, template := range templates {
		result = append(result, imagingResultTemplateControllerResponses.StaffGetListImagingResultTemplateResponse{
			Uuid:         template.Uuid,
			Modality:     template.Modality,
			Name:         template.Name,
			BodyParts:    emptyIfNil(template.BodyParts),
			DisplayOrder: template.DisplayOrder,
			IsActive:     template.IsActive,
			CreatedAt:    template.CreatedAt,
		})
	}

	return result
}
