package objectMappers

import (
	"telerad-core-module/internals/entities"
	imagingResultSheetTemplateControllerResponses "telerad-core-module/internals/responses/imaging-result-sheet-template-controller_responses"
)

func ToStaffGetAImagingResultSheetTemplateResponse(template entities.ImagingResultSheetTemplateEntity) imagingResultSheetTemplateControllerResponses.StaffGetAImagingResultSheetTemplateResponse {
	response := imagingResultSheetTemplateControllerResponses.StaffGetAImagingResultSheetTemplateResponse{
		Uuid:               template.Uuid,
		CreatedAt:          template.CreatedAt,
		CreatedBy:          template.CreatedBy,
		UpdatedAt:          template.UpdatedAt,
		UpdatedBy:          template.UpdatedBy,
		TeleradPartnerUuid: template.TeleradPartnerUuid,
		HtmlContent:        template.HtmlContent,
		IsActive:           template.IsActive,
	}

	if template.TeleradPartner != nil {
		response.TeleradPartnerCode = &template.TeleradPartner.Code
		response.TeleradPartnerName = &template.TeleradPartner.Name
	}

	return response
}

func ToStaffGetListImagingResultSheetTemplateSlice(templates []entities.ImagingResultSheetTemplateEntity) []imagingResultSheetTemplateControllerResponses.StaffGetListImagingResultSheetTemplateResponse {
	result := make([]imagingResultSheetTemplateControllerResponses.StaffGetListImagingResultSheetTemplateResponse, 0, len(templates))

	for _, template := range templates {
		element := imagingResultSheetTemplateControllerResponses.StaffGetListImagingResultSheetTemplateResponse{
			Uuid:               template.Uuid,
			TeleradPartnerUuid: template.TeleradPartnerUuid,
			IsActive:           template.IsActive,
			CreatedAt:          template.CreatedAt,
		}

		if template.TeleradPartner != nil {
			element.TeleradPartnerCode = &template.TeleradPartner.Code
			element.TeleradPartnerName = &template.TeleradPartner.Name
		}

		result = append(result, element)
	}

	return result
}
