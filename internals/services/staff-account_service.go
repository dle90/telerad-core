package services

import (
	"context"
	"slices"
	"strings"

	baseServices "telerad-core-module/internals/base-services"
	"telerad-core-module/internals/constants"
	"telerad-core-module/internals/entities"
	objectMappers "telerad-core-module/internals/object-mappers"
	"telerad-core-module/internals/repositories"
	staffAccountControllerRequests "telerad-core-module/internals/requests/staff-account-controller_requests"
	"telerad-core-module/internals/responses"
	staffAccountControllerResponses "telerad-core-module/internals/responses/staff-account-controller_responses"
	"telerad-core-module/internals/secure"

	_errorMessages "telerad-core-module/error"

	_error "telerad-core-module/error"

	"github.com/google/uuid"
)

func StaffGetPaginatedStaffAccounts(
	ctx context.Context,
	page, pageSize int,
	search string,
	isActive *bool,
) (*responses.PaginationResponse, *_error.SystemError) {
	staffs, totalCount, err := repositories.FindPaginatedStaffAccounts(ctx, bunNoTransaction, page, pageSize, strings.TrimSpace(search), isActive)
	if err != nil {
		return nil, _error.New(err)
	}

	records := objectMappers.ToStaffGetListStaffAccountSlice(staffs)
	response := responses.NewPaginationResponse(totalCount, page, pageSize, records)
	return &response, nil
}

func StaffGetAStaffAccount(ctx context.Context, staffUuid uuid.UUID) (*staffAccountControllerResponses.StaffGetAStaffAccountResponse, *_error.SystemError) {
	staff, systemErr := findStaffAccountOrFail(ctx, staffUuid)
	if systemErr != nil {
		return nil, systemErr
	}

	response := objectMappers.ToStaffGetAStaffAccountResponse(*staff)
	return &response, nil
}

func StaffCreateStaffAccount(
	ctx context.Context,
	creatorUuid uuid.UUID,
	request staffAccountControllerRequests.StaffCreateStaffAccountRequest,
) (*staffAccountControllerResponses.StaffGetAStaffAccountResponse, *_error.SystemError) {
	request.Code = strings.TrimSpace(request.Code)

	if existing, err := repositories.FindOneStaffAccountByCode(ctx, bunNoTransaction, request.Code, nil); err != nil {
		return nil, _error.New(err)
	} else if existing != nil {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E102_002)
	}

	staff := baseServices.InitNewStaffAccount(request)
	if err := baseServices.CreateNewStaffAccount(ctx, bunNoTransaction, creatorUuid, &staff); err != nil {
		return nil, _error.New(err)
	}

	response := objectMappers.ToStaffGetAStaffAccountResponse(staff)
	return &response, nil
}

func StaffUpdateStaffAccount(
	ctx context.Context,
	updaterUuid, staffUuid uuid.UUID,
	request staffAccountControllerRequests.StaffUpdateStaffAccountRequest,
) (*staffAccountControllerResponses.StaffGetAStaffAccountResponse, *_error.SystemError) {
	request.Code = strings.TrimSpace(request.Code)

	staff, systemErr := findStaffAccountOrFail(ctx, staffUuid)
	if systemErr != nil {
		return nil, systemErr
	}

	// Không cho sửa tài khoản quản trị.
	if staffHasAdminRole(staff) {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E102_006)
	}

	if !strings.EqualFold(request.Code, staff.Code) {
		if duplicate, err := repositories.FindOneStaffAccountByCode(ctx, bunNoTransaction, request.Code, &staff.Uuid); err != nil {
			return nil, _error.New(err)
		} else if duplicate != nil {
			return nil, _error.NewErrorByString(_errorMessages.TELERAD_E102_002)
		}
	}

	staff.Code = request.Code
	staff.FullName = request.FullName
	staff.Gender = request.Gender
	staff.DateOfBirth = request.DateOfBirth
	staff.CitizenIdentityNumber = request.CitizenIdentityNumber
	staff.Phone = request.Phone
	staff.Email = request.Email
	staff.FullAddress = request.FullAddress

	if err := baseServices.UpdateWholeStaffAccountRecord(ctx, bunNoTransaction, updaterUuid, staff); err != nil {
		return nil, _error.New(err)
	}

	response := objectMappers.ToStaffGetAStaffAccountResponse(*staff)
	return &response, nil
}

// StaffAssignReadingPermission phân quyền đọc phim: cập nhật modalities + danh sách
// đối tác telerad mà nhân viên được đọc phim.
func StaffAssignReadingPermission(
	ctx context.Context,
	updaterUuid, staffUuid uuid.UUID,
	request staffAccountControllerRequests.StaffAssignReadingPermissionRequest,
) (*staffAccountControllerResponses.StaffGetAStaffAccountResponse, *_error.SystemError) {
	staff, systemErr := findStaffAccountOrFail(ctx, staffUuid)
	if systemErr != nil {
		return nil, systemErr
	}

	staff.Modalities = request.Modalities
	staff.TeleradPartnerUuids = request.TeleradPartnerUuids

	if err := baseServices.UpdateWholeStaffAccountRecord(ctx, bunNoTransaction, updaterUuid, staff); err != nil {
		return nil, _error.New(err)
	}

	response := objectMappers.ToStaffGetAStaffAccountResponse(*staff)
	return &response, nil
}

// StaffAssignRoles phân roles cho nhân viên.
func StaffAssignRoles(
	ctx context.Context,
	updaterUuid, staffUuid uuid.UUID,
	request staffAccountControllerRequests.StaffAssignRolesRequest,
) (*staffAccountControllerResponses.StaffGetAStaffAccountResponse, *_error.SystemError) {
	staff, systemErr := findStaffAccountOrFail(ctx, staffUuid)
	if systemErr != nil {
		return nil, systemErr
	}

	roles := request.Roles
	if roles == nil {
		roles = []string{}
	}
	staff.Roles = roles

	if err := baseServices.UpdateWholeStaffAccountRecord(ctx, bunNoTransaction, updaterUuid, staff); err != nil {
		return nil, _error.New(err)
	}

	response := objectMappers.ToStaffGetAStaffAccountResponse(*staff)
	return &response, nil
}

func StaffActivateStaffAccount(ctx context.Context, updaterUuid, staffUuid uuid.UUID) *_error.SystemError {
	return changeStaffAccountActive(ctx, updaterUuid, staffUuid, true)
}

func StaffDeactivateStaffAccount(ctx context.Context, updaterUuid, staffUuid uuid.UUID) *_error.SystemError {
	return changeStaffAccountActive(ctx, updaterUuid, staffUuid, false)
}

// StaffCreateAccount cấp tài khoản đăng nhập cho nhân viên CHƯA có username.
// Admin cung cấp username; mật khẩu tự sinh, trả về 1 lần.
func StaffCreateAccount(
	ctx context.Context,
	updaterUuid, staffUuid uuid.UUID,
	request staffAccountControllerRequests.StaffCreateAccountRequest,
) (*staffAccountControllerResponses.StaffAccountCredentialResponse, *_error.SystemError) {
	staff, systemErr := findStaffAccountOrFail(ctx, staffUuid)
	if systemErr != nil {
		return nil, systemErr
	}

	if staff.Username != nil {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E102_004)
	}

	username := strings.ToUpper(strings.TrimSpace(request.Username))

	if existing, err := repositories.GetStaffAccountByUsername(ctx, bunNoTransaction, username); err != nil {
		return nil, _error.New(err)
	} else if existing != nil {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E102_003)
	}

	plainPassword := secure.GenerateRandomPassword()
	passwordHash, err := secure.GenerateBcryptHash(plainPassword)
	if err != nil {
		return nil, _error.New(err)
	}

	staff.Username = &username
	staff.PasswordHash = &passwordHash

	if err := baseServices.UpdateWholeStaffAccountRecord(ctx, bunNoTransaction, updaterUuid, staff); err != nil {
		return nil, _error.New(err)
	}

	return &staffAccountControllerResponses.StaffAccountCredentialResponse{
		Uuid:     staff.Uuid,
		Username: username,
		Password: plainPassword,
	}, nil
}

// StaffResetPassword reset mật khẩu cho nhân viên ĐÃ có username (tự sinh mới, trả 1 lần).
func StaffResetPassword(
	ctx context.Context,
	updaterUuid, staffUuid uuid.UUID,
) (*staffAccountControllerResponses.StaffAccountCredentialResponse, *_error.SystemError) {
	staff, systemErr := findStaffAccountOrFail(ctx, staffUuid)
	if systemErr != nil {
		return nil, systemErr
	}

	// Không cho reset mật khẩu tài khoản quản trị.
	if staffHasAdminRole(staff) {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E102_006)
	}

	if staff.Username == nil {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E102_005)
	}

	plainPassword := secure.GenerateRandomPassword()
	passwordHash, err := secure.GenerateBcryptHash(plainPassword)
	if err != nil {
		return nil, _error.New(err)
	}

	staff.PasswordHash = &passwordHash

	if err := baseServices.UpdateWholeStaffAccountRecord(ctx, bunNoTransaction, updaterUuid, staff); err != nil {
		return nil, _error.New(err)
	}

	return &staffAccountControllerResponses.StaffAccountCredentialResponse{
		Uuid:     staff.Uuid,
		Username: *staff.Username,
		Password: plainPassword,
	}, nil
}

func UserGetMe(ctx context.Context, userUuid uuid.UUID) (*staffAccountControllerResponses.UserGetMeResponse, *_error.SystemError) {
	staff, systemErr := findStaffAccountOrFail(ctx, userUuid)
	if systemErr != nil {
		return nil, systemErr
	}

	response := objectMappers.ToUserGetMeResponse(*staff)
	return &response, nil
}

// UserChangePassword — user tự đổi mật khẩu: xác minh mật khẩu hiện tại trước khi đổi.
func UserChangePassword(
	ctx context.Context,
	userUuid uuid.UUID,
	request staffAccountControllerRequests.UserChangePasswordRequest,
) *_error.SystemError {
	staff, systemErr := findStaffAccountOrFail(ctx, userUuid)
	if systemErr != nil {
		return systemErr
	}

	if staff.Username == nil || staff.PasswordHash == nil {
		return _error.NewErrorByString(_errorMessages.TELERAD_E102_005)
	}

	if !secure.VerifyBcryptPassword(*staff.PasswordHash, request.OldPassword) {
		return _error.NewErrorByString(_errorMessages.TELERAD_E001_001)
	}

	passwordHash, err := secure.GenerateBcryptHash(request.NewPassword)
	if err != nil {
		return _error.New(err)
	}

	staff.PasswordHash = &passwordHash

	if err := baseServices.UpdateWholeStaffAccountRecord(ctx, bunNoTransaction, staff.Uuid, staff); err != nil {
		return _error.New(err)
	}

	return nil
}

func changeStaffAccountActive(ctx context.Context, updaterUuid, staffUuid uuid.UUID, isActive bool) *_error.SystemError {
	staff, systemErr := findStaffAccountOrFail(ctx, staffUuid)
	if systemErr != nil {
		return systemErr
	}

	if staff.IsActive == isActive {
		return nil
	}

	staff.IsActive = isActive

	if err := baseServices.UpdateWholeStaffAccountRecord(ctx, bunNoTransaction, updaterUuid, staff); err != nil {
		return _error.New(err)
	}

	return nil
}

// staffHasAdminRole — tài khoản quản trị được bảo vệ: không cho sửa / reset mật
// khẩu và bị ẩn khỏi danh sách nhân sự.
func staffHasAdminRole(staff *entities.StaffAccountEntity) bool {
	return staff != nil && slices.Contains(staff.Roles, constants.ROLE_ADMIN)
}

func findStaffAccountOrFail(ctx context.Context, staffUuid uuid.UUID) (*entities.StaffAccountEntity, *_error.SystemError) {
	staff, err := baseServices.FindOneStaffAccountByUuid(ctx, bunNoTransaction, staffUuid)
	if err != nil {
		return nil, _error.New(err)
	} else if staff == nil {
		return nil, _error.NewErrorByString(_errorMessages.TELERAD_E102_001)
	}

	return staff, nil
}
