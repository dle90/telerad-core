package services

import (
	"context"
	"strings"

	baseServices "telerad-core-module/internals/base-services"
	"telerad-core-module/internals/entities"
	objectMappers "telerad-core-module/internals/object-mappers"
	"telerad-core-module/internals/repositories"
	teleradPartnerControllerRequests "telerad-core-module/internals/requests/telerad-partner-controller_requests"
	"telerad-core-module/internals/responses"
	teleradPartnerControllerResponses "telerad-core-module/internals/responses/telerad-partner-controller_responses"
	"telerad-core-module/internals/secure"

	_errorMessages "telerad-core-module/error"

	_error "telerad-core-module/error"

	"github.com/google/uuid"
)

func StaffGetPaginatedTeleradPartners(
	ctx context.Context,
	page, pageSize int,
	search string,
	isActive *bool,
) (*responses.PaginationResponse, *_error.SystemError) {
	partners, totalCount, err := repositories.FindPaginatedTeleradPartners(ctx, bunNoTransaction, page, pageSize, strings.TrimSpace(search), isActive)
	if err != nil {
		return nil, _error.New(err)
	}

	records := objectMappers.ToStaffGetListTeleradPartnerSlice(partners)
	response := responses.NewPaginationResponse(totalCount, page, pageSize, records)
	return &response, nil
}

// StaffGetAllTeleradPartners trả TẤT CẢ đối tác (kèm trạng thái) cho màn phân quyền
// đọc phim: user xem danh sách, tick chọn rồi gọi StaffAssignReadingPermission.
func StaffGetAllTeleradPartners(ctx context.Context) ([]teleradPartnerControllerResponses.StaffGetAllTeleradPartnerResponse, *_error.SystemError) {
	partners, err := repositories.FindAllTeleradPartners(ctx, bunNoTransaction)
	if err != nil {
		return nil, _error.New(err)
	}

	return objectMappers.ToStaffGetAllTeleradPartnerSlice(partners), nil
}

func StaffGetATeleradPartner(ctx context.Context, partnerUuid uuid.UUID) (*teleradPartnerControllerResponses.StaffGetATeleradPartnerResponse, *_error.SystemError) {
	partner, systemErr := findTeleradPartnerOrFail(ctx, partnerUuid)
	if systemErr != nil {
		return nil, systemErr
	}

	response := objectMappers.ToStaffGetATeleradPartnerResponse(*partner)
	return &response, nil
}

func StaffCreateTeleradPartner(
	ctx context.Context,
	creatorUuid uuid.UUID,
	request teleradPartnerControllerRequests.StaffCreateTeleradPartnerRequest,
) (*teleradPartnerControllerResponses.StaffCreateTeleradPartnerResponse, *_error.SystemError) {
	request.Code = strings.TrimSpace(request.Code)
	username := strings.ToUpper(strings.TrimSpace(request.Username))

	if existing, err := repositories.FindOneTeleradPartnerByCode(ctx, bunNoTransaction, request.Code, nil); err != nil {
		return nil, _error.New(err)
	} else if existing != nil {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E101_002)
	}

	if existing, err := repositories.FindOneTeleradPartnerByUsername(ctx, bunNoTransaction, username, nil); err != nil {
		return nil, _error.New(err)
	} else if existing != nil {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E101_003)
	}

	passwordHash, err := secure.GenerateBcryptHash(request.Password)
	if err != nil {
		return nil, _error.New(err)
	}

	partner := baseServices.InitNewTeleradPartner(request, username, passwordHash)
	if err := baseServices.CreateNewTeleradPartner(ctx, bunNoTransaction, creatorUuid, &partner); err != nil {
		return nil, _error.New(err)
	}

	return &teleradPartnerControllerResponses.StaffCreateTeleradPartnerResponse{
		Uuid:     partner.Uuid,
		Code:     partner.Code,
		Name:     partner.Name,
		IsActive: partner.IsActive,
		Username: partner.Username,
	}, nil
}

func StaffUpdateTeleradPartner(
	ctx context.Context,
	updaterUuid, partnerUuid uuid.UUID,
	request teleradPartnerControllerRequests.StaffUpdateTeleradPartnerRequest,
) (*teleradPartnerControllerResponses.StaffGetATeleradPartnerResponse, *_error.SystemError) {
	request.Code = strings.TrimSpace(request.Code)

	partner, systemErr := findTeleradPartnerOrFail(ctx, partnerUuid)
	if systemErr != nil {
		return nil, systemErr
	}

	if !strings.EqualFold(request.Code, partner.Code) {
		if duplicate, err := repositories.FindOneTeleradPartnerByCode(ctx, bunNoTransaction, request.Code, &partner.Uuid); err != nil {
			return nil, _error.New(err)
		} else if duplicate != nil {
			return nil, _error.NewErrorByString(_errorMessages.TELERAD_E101_002)
		}
	}

	partner.Code = request.Code
	partner.Name = request.Name
	partner.Contact = request.Contact
	partner.Modalities = request.Modalities

	if err := baseServices.UpdateWholeTeleradPartnerRecord(ctx, bunNoTransaction, updaterUuid, partner); err != nil {
		return nil, _error.New(err)
	}

	response := objectMappers.ToStaffGetATeleradPartnerResponse(*partner)
	return &response, nil
}

func StaffActivateTeleradPartner(ctx context.Context, updaterUuid, partnerUuid uuid.UUID) *_error.SystemError {
	return changeTeleradPartnerActive(ctx, updaterUuid, partnerUuid, true)
}

func StaffDeactivateTeleradPartner(ctx context.Context, updaterUuid, partnerUuid uuid.UUID) *_error.SystemError {
	return changeTeleradPartnerActive(ctx, updaterUuid, partnerUuid, false)
}

func StaffGetTeleradPartnerPartnerConfig(ctx context.Context, partnerUuid uuid.UUID) (*teleradPartnerControllerResponses.StaffTeleradPartnerPartnerConfigResponse, *_error.SystemError) {
	partner, systemErr := findTeleradPartnerOrFail(ctx, partnerUuid)
	if systemErr != nil {
		return nil, systemErr
	}

	response := objectMappers.ToStaffTeleradPartnerPartnerConfigResponse(*partner)
	return &response, nil
}

func StaffUpdateTeleradPartnerPartnerConfig(
	ctx context.Context,
	updaterUuid, partnerUuid uuid.UUID,
	request teleradPartnerControllerRequests.StaffUpdateTeleradPartnerPartnerConfigRequest,
) (*teleradPartnerControllerResponses.StaffTeleradPartnerPartnerConfigResponse, *_error.SystemError) {
	partner, systemErr := findTeleradPartnerOrFail(ctx, partnerUuid)
	if systemErr != nil {
		return nil, systemErr
	}

	partner.Callback = request.Callback
	partner.CallbackUrl = request.CallbackUrl
	partner.PartnerUsername = request.PartnerUsername
	partner.PartnerPassword = request.PartnerPassword

	if err := baseServices.UpdateWholeTeleradPartnerRecord(ctx, bunNoTransaction, updaterUuid, partner); err != nil {
		return nil, _error.New(err)
	}

	response := objectMappers.ToStaffTeleradPartnerPartnerConfigResponse(*partner)
	return &response, nil
}

// StaffChangeTeleradPartnerPassword đổi mật khẩu phía telerad (admin cung cấp mật khẩu mới).
// Không cho đổi username.
func StaffChangeTeleradPartnerPassword(ctx context.Context, updaterUuid, partnerUuid uuid.UUID, request teleradPartnerControllerRequests.StaffChangeTeleradPartnerPasswordRequest) *_error.SystemError {
	partner, systemErr := findTeleradPartnerOrFail(ctx, partnerUuid)
	if systemErr != nil {
		return systemErr
	}

	passwordHash, err := secure.GenerateBcryptHash(request.Password)
	if err != nil {
		return _error.New(err)
	}

	partner.PasswordHash = passwordHash

	if err := baseServices.UpdateWholeTeleradPartnerRecord(ctx, bunNoTransaction, updaterUuid, partner); err != nil {
		return _error.New(err)
	}

	return nil
}

func changeTeleradPartnerActive(ctx context.Context, updaterUuid, partnerUuid uuid.UUID, isActive bool) *_error.SystemError {
	partner, systemErr := findTeleradPartnerOrFail(ctx, partnerUuid)
	if systemErr != nil {
		return systemErr
	}

	if partner.IsActive == isActive {
		return nil
	}

	partner.IsActive = isActive

	if err := baseServices.UpdateWholeTeleradPartnerRecord(ctx, bunNoTransaction, updaterUuid, partner); err != nil {
		return _error.New(err)
	}

	return nil
}

func findTeleradPartnerOrFail(ctx context.Context, partnerUuid uuid.UUID) (*entities.TeleradPartnerEntity, *_error.SystemError) {
	partner, err := baseServices.FindOneTeleradPartnerByUuid(ctx, bunNoTransaction, partnerUuid)
	if err != nil {
		return nil, _error.New(err)
	} else if partner == nil {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E101_001)
	}

	return partner, nil
}
