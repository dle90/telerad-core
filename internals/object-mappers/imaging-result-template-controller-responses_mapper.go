package objectMappers

import (
	"telerad-core-module/internals/entities"
	fieldValues "telerad-core-module/internals/entities/field-values"
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

func ToStaffGetImagingResultTemplateFormOptionsResponse(modalities []string, bodyParts []fieldValues.ColumnValueString) imagingResultTemplateControllerResponses.StaffGetImagingResultTemplateFormOptionsResponse {
	options := make([]imagingResultTemplateControllerResponses.ImagingResultTemplateOption, 0, len(bodyParts))
	for _, bodyPart := range bodyParts {
		options = append(options, imagingResultTemplateControllerResponses.ImagingResultTemplateOption{
			Value: bodyPart.Value,
			Name:  bodyPart.Name,
		})
	}

	return imagingResultTemplateControllerResponses.StaffGetImagingResultTemplateFormOptionsResponse{
		Modalities: modalities,
		BodyParts:  options,
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
