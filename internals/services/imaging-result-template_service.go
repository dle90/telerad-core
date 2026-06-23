package services

import (
	"context"
	"strings"

	baseServices "telerad-core-module/internals/base-services"
	"telerad-core-module/internals/constants"
	"telerad-core-module/internals/entities"
	fieldValues "telerad-core-module/internals/entities/field-values"
	objectMappers "telerad-core-module/internals/object-mappers"
	"telerad-core-module/internals/repositories"
	imagingResultTemplateControllerRequests "telerad-core-module/internals/requests/imaging-result-template-controller_requests"
	"telerad-core-module/internals/responses"
	imagingResultTemplateControllerResponses "telerad-core-module/internals/responses/imaging-result-template-controller_responses"

	_errorMessages "telerad-core-module/error"

	_error "telerad-core-module/error"

	"github.com/google/uuid"
)

func StaffGetPaginatedImagingResultTemplates(
	ctx context.Context,
	page, pageSize int,
	modality string,
	search string,
	isActive *bool,
	bodyParts []string,
) (*responses.PaginationResponse, *_error.SystemError) {
	templates, totalCount, err := repositories.FindPaginatedImagingResultTemplates(ctx, bunNoTransaction, page, pageSize, strings.TrimSpace(modality), strings.TrimSpace(search), isActive, bodyParts)
	if err != nil {
		return nil, _error.New(err)
	}

	records := objectMappers.ToStaffGetListImagingResultTemplateSlice(templates)
	response := responses.NewPaginationResponse(totalCount, page, pageSize, records)
	return &response, nil
}

// StaffGetImagingResultTemplateFormOptions trả option cho form mẫu kết quả: loại chụp
// hệ thống hỗ trợ + danh sách bộ phận chụp (PACS_BODY_PART).
func StaffGetImagingResultTemplateFormOptions() imagingResultTemplateControllerResponses.StaffGetImagingResultTemplateFormOptionsResponse {
	bodyParts := fieldValues.GetAllStringTypeByCode(fieldValues.PACS_BODY_PART)

	options := make([]imagingResultTemplateControllerResponses.ImagingResultTemplateOption, 0, len(bodyParts))
	for _, bodyPart := range bodyParts {
		options = append(options, imagingResultTemplateControllerResponses.ImagingResultTemplateOption{
			Value: bodyPart.Value,
			Name:  bodyPart.Name,
		})
	}

	return imagingResultTemplateControllerResponses.StaffGetImagingResultTemplateFormOptionsResponse{
		Modalities: constants.MODALITIES,
		BodyParts:  options,
	}
}

func StaffGetAImagingResultTemplate(ctx context.Context, templateUuid uuid.UUID) (*imagingResultTemplateControllerResponses.StaffGetAImagingResultTemplateResponse, *_error.SystemError) {
	template, systemErr := findImagingResultTemplateOrFail(ctx, templateUuid)
	if systemErr != nil {
		return nil, systemErr
	}

	response := objectMappers.ToStaffGetAImagingResultTemplateResponse(*template)
	return &response, nil
}

func StaffCreateImagingResultTemplate(
	ctx context.Context,
	creatorUuid uuid.UUID,
	request imagingResultTemplateControllerRequests.StaffCreateImagingResultTemplateRequest,
) (*imagingResultTemplateControllerResponses.StaffGetAImagingResultTemplateResponse, *_error.SystemError) {
	request.Name = strings.TrimSpace(request.Name)

	if systemErr := validateImagingResultTemplateBodyParts(request.BodyParts); systemErr != nil {
		return nil, systemErr
	}

	template := baseServices.InitNewImagingResultTemplate(request)
	if err := baseServices.CreateNewImagingResultTemplate(ctx, bunNoTransaction, creatorUuid, &template); err != nil {
		return nil, _error.New(err)
	}

	response := objectMappers.ToStaffGetAImagingResultTemplateResponse(template)
	return &response, nil
}

func StaffUpdateImagingResultTemplate(
	ctx context.Context,
	updaterUuid, templateUuid uuid.UUID,
	request imagingResultTemplateControllerRequests.StaffUpdateImagingResultTemplateRequest,
) (*imagingResultTemplateControllerResponses.StaffGetAImagingResultTemplateResponse, *_error.SystemError) {
	request.Name = strings.TrimSpace(request.Name)

	template, systemErr := findImagingResultTemplateOrFail(ctx, templateUuid)
	if systemErr != nil {
		return nil, systemErr
	}

	if systemErr := validateImagingResultTemplateBodyParts(request.BodyParts); systemErr != nil {
		return nil, systemErr
	}

	template.Modality = request.Modality
	template.Name = request.Name
	template.BodyParts = request.BodyParts
	template.HtmlContent = request.HtmlContent
	template.FontSize = request.FontSize
	template.LineSpacing = request.LineSpacing
	template.DisplayOrder = request.DisplayOrder

	if err := baseServices.UpdateWholeImagingResultTemplateRecord(ctx, bunNoTransaction, updaterUuid, template); err != nil {
		return nil, _error.New(err)
	}

	response := objectMappers.ToStaffGetAImagingResultTemplateResponse(*template)
	return &response, nil
}

func StaffActivateImagingResultTemplate(ctx context.Context, updaterUuid, templateUuid uuid.UUID) *_error.SystemError {
	return changeImagingResultTemplateActive(ctx, updaterUuid, templateUuid, true)
}

func StaffDeactivateImagingResultTemplate(ctx context.Context, updaterUuid, templateUuid uuid.UUID) *_error.SystemError {
	return changeImagingResultTemplateActive(ctx, updaterUuid, templateUuid, false)
}

func changeImagingResultTemplateActive(ctx context.Context, updaterUuid, templateUuid uuid.UUID, isActive bool) *_error.SystemError {
	template, systemErr := findImagingResultTemplateOrFail(ctx, templateUuid)
	if systemErr != nil {
		return systemErr
	}

	if template.IsActive == isActive {
		return nil
	}

	template.IsActive = isActive

	if err := baseServices.UpdateWholeImagingResultTemplateRecord(ctx, bunNoTransaction, updaterUuid, template); err != nil {
		return _error.New(err)
	}

	return nil
}

func findImagingResultTemplateOrFail(ctx context.Context, templateUuid uuid.UUID) (*entities.ImagingResultTemplateEntity, *_error.SystemError) {
	template, err := baseServices.FindOneImagingResultTemplateByUuid(ctx, bunNoTransaction, templateUuid)
	if err != nil {
		return nil, _error.New(err)
	} else if template == nil {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E104_001)
	}

	return template, nil
}

// validateImagingResultTemplateBodyParts đảm bảo mọi mã bộ phận chụp đều thuộc
// field-values PACS_BODY_PART.
func validateImagingResultTemplateBodyParts(bodyParts []string) *_error.SystemError {
	for _, bodyPart := range bodyParts {
		if fieldValues.FromValueAndCodeString(bodyPart, fieldValues.PACS_BODY_PART) == nil {
			return _error.NewErrorByString(_errorMessages.TELERAD_E104_003)
		}
	}
	return nil
}
