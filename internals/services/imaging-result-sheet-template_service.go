package services

import (
	"context"

	baseServices "telerad-core-module/internals/base-services"
	"telerad-core-module/internals/entities"
	objectMappers "telerad-core-module/internals/object-mappers"
	"telerad-core-module/internals/repositories"
	imagingResultSheetTemplateControllerRequests "telerad-core-module/internals/requests/imaging-result-sheet-template-controller_requests"
	"telerad-core-module/internals/responses"
	imagingResultSheetTemplateControllerResponses "telerad-core-module/internals/responses/imaging-result-sheet-template-controller_responses"

	_errorMessages "telerad-core-module/error"

	_error "telerad-core-module/error"

	"github.com/google/uuid"
)

func StaffGetPaginatedImagingResultSheetTemplates(
	ctx context.Context,
	page, pageSize int,
	teleradPartnerUuid *uuid.UUID,
	isActive *bool,
) (*responses.PaginationResponse, *_error.SystemError) {
	templates, totalCount, err := repositories.FindPaginatedImagingResultSheetTemplates(ctx, bunNoTransaction, page, pageSize, teleradPartnerUuid, isActive)
	if err != nil {
		return nil, _error.New(err)
	}

	records := objectMappers.ToStaffGetListImagingResultSheetTemplateSlice(templates)
	response := responses.NewPaginationResponse(totalCount, page, pageSize, records)
	return &response, nil
}

func StaffGetAImagingResultSheetTemplate(ctx context.Context, templateUuid uuid.UUID) (*imagingResultSheetTemplateControllerResponses.StaffGetAImagingResultSheetTemplateResponse, *_error.SystemError) {
	template, systemErr := findImagingResultSheetTemplateWithPartnerOrFail(ctx, templateUuid)
	if systemErr != nil {
		return nil, systemErr
	}

	response := objectMappers.ToStaffGetAImagingResultSheetTemplateResponse(*template)
	return &response, nil
}

func StaffCreateImagingResultSheetTemplate(
	ctx context.Context,
	creatorUuid uuid.UUID,
	request imagingResultSheetTemplateControllerRequests.StaffCreateImagingResultSheetTemplateRequest,
) (*imagingResultSheetTemplateControllerResponses.StaffGetAImagingResultSheetTemplateResponse, *_error.SystemError) {
	// CSYT (telerad_partner) phải tồn tại.
	if _, systemErr := findTeleradPartnerOrFail(ctx, request.TeleradPartnerUuid); systemErr != nil {
		return nil, systemErr
	}

	// Mỗi CSYT chỉ có 1 phiếu kết quả.
	if existing, err := repositories.FindOneImagingResultSheetTemplateByPartner(ctx, bunNoTransaction, request.TeleradPartnerUuid, nil); err != nil {
		return nil, _error.New(err)
	} else if existing != nil {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E105_002)
	}

	template := baseServices.InitNewImagingResultSheetTemplate(request.TeleradPartnerUuid, request.HtmlContent, request.ResultFontSize, request.ResultLineSpacing)
	if err := baseServices.CreateNewImagingResultSheetTemplate(ctx, bunNoTransaction, creatorUuid, &template); err != nil {
		return nil, _error.New(err)
	}

	created, systemErr := findImagingResultSheetTemplateWithPartnerOrFail(ctx, template.Uuid)
	if systemErr != nil {
		return nil, systemErr
	}

	response := objectMappers.ToStaffGetAImagingResultSheetTemplateResponse(*created)
	return &response, nil
}

func StaffUpdateImagingResultSheetTemplate(
	ctx context.Context,
	updaterUuid, templateUuid uuid.UUID,
	request imagingResultSheetTemplateControllerRequests.StaffUpdateImagingResultSheetTemplateRequest,
) (*imagingResultSheetTemplateControllerResponses.StaffGetAImagingResultSheetTemplateResponse, *_error.SystemError) {
	template, systemErr := findImagingResultSheetTemplateOrFail(ctx, templateUuid)
	if systemErr != nil {
		return nil, systemErr
	}

	template.HtmlContent = request.HtmlContent
	template.ResultFontSize = request.ResultFontSize
	template.ResultLineSpacing = request.ResultLineSpacing

	if err := baseServices.UpdateWholeImagingResultSheetTemplateRecord(ctx, bunNoTransaction, updaterUuid, template); err != nil {
		return nil, _error.New(err)
	}

	updated, systemErr := findImagingResultSheetTemplateWithPartnerOrFail(ctx, template.Uuid)
	if systemErr != nil {
		return nil, systemErr
	}

	response := objectMappers.ToStaffGetAImagingResultSheetTemplateResponse(*updated)
	return &response, nil
}

func StaffActivateImagingResultSheetTemplate(ctx context.Context, updaterUuid, templateUuid uuid.UUID) *_error.SystemError {
	return changeImagingResultSheetTemplateActive(ctx, updaterUuid, templateUuid, true)
}

func StaffDeactivateImagingResultSheetTemplate(ctx context.Context, updaterUuid, templateUuid uuid.UUID) *_error.SystemError {
	return changeImagingResultSheetTemplateActive(ctx, updaterUuid, templateUuid, false)
}

func changeImagingResultSheetTemplateActive(ctx context.Context, updaterUuid, templateUuid uuid.UUID, isActive bool) *_error.SystemError {
	template, systemErr := findImagingResultSheetTemplateOrFail(ctx, templateUuid)
	if systemErr != nil {
		return systemErr
	}

	if template.IsActive == isActive {
		return nil
	}

	template.IsActive = isActive

	if err := baseServices.UpdateWholeImagingResultSheetTemplateRecord(ctx, bunNoTransaction, updaterUuid, template); err != nil {
		return _error.New(err)
	}

	return nil
}

func findImagingResultSheetTemplateOrFail(ctx context.Context, templateUuid uuid.UUID) (*entities.ImagingResultSheetTemplateEntity, *_error.SystemError) {
	template, err := baseServices.FindOneImagingResultSheetTemplateByUuid(ctx, bunNoTransaction, templateUuid)
	if err != nil {
		return nil, _error.New(err)
	} else if template == nil {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E105_001)
	}

	return template, nil
}

func findImagingResultSheetTemplateWithPartnerOrFail(ctx context.Context, templateUuid uuid.UUID) (*entities.ImagingResultSheetTemplateEntity, *_error.SystemError) {
	template, err := repositories.FindOneImagingResultSheetTemplateWithPartnerByUuid(ctx, bunNoTransaction, templateUuid)
	if err != nil {
		return nil, _error.New(err)
	} else if template == nil {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E105_001)
	}

	return template, nil
}
