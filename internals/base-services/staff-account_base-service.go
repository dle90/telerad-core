package baseServices

import (
	"context"
	"strings"
	"time"

	"telerad-core-module/internals/entities"
	"telerad-core-module/internals/repositories"
	staffAccountControllerRequests "telerad-core-module/internals/requests/staff-account-controller_requests"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func FindOneStaffAccountByUuid(ctx context.Context, tx bun.IDB, id uuid.UUID) (*entities.StaffAccountEntity, error) {
	return repositories.FindOneByUuid[entities.StaffAccountEntity](ctx, tx, id)
}

// GetStaffAccountByUsername — uppercase username trước khi query để lookup
// case-insensitive (username luôn lưu uppercase).
func GetStaffAccountByUsername(ctx context.Context, tx bun.IDB, username string) (*entities.StaffAccountEntity, error) {
	return repositories.GetStaffAccountByUsername(ctx, tx, strings.ToUpper(username))
}

// InitNewStaffAccount dựng hồ sơ nhân viên mới (chưa có tài khoản đăng nhập:
// username/password = nil). is_active mặc định true.
func InitNewStaffAccount(request staffAccountControllerRequests.StaffCreateStaffAccountRequest) entities.StaffAccountEntity {
	roles := request.Roles
	if roles == nil {
		roles = []string{}
	}

	return entities.StaffAccountEntity{
		Code:                  request.Code,
		FullName:              request.FullName,
		Gender:                request.Gender,
		DateOfBirth:           request.DateOfBirth,
		CitizenIdentityNumber: request.CitizenIdentityNumber,
		Phone:                 request.Phone,
		Email:                 request.Email,
		FullAddress:           request.FullAddress,
		IsActive:              true,
		Roles:                 roles,
	}
}

func CreateNewStaffAccount(ctx context.Context, tx bun.IDB, creatorUuid uuid.UUID, newRecord *entities.StaffAccountEntity) error {
	newRecord.CreatedAt = time.Now()
	newRecord.CreatedBy = &creatorUuid

	return repositories.InsertOne(ctx, tx, newRecord)
}

func UpdateWholeStaffAccountRecord(ctx context.Context, tx bun.IDB, updaterUuid uuid.UUID, staff *entities.StaffAccountEntity) error {
	now := time.Now()
	staff.UpdatedAt = &now
	staff.UpdatedBy = &updaterUuid

	return repositories.UpdateWholeRecord(ctx, tx, staff)
}
