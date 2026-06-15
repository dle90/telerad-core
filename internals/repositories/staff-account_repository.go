package repositories

import (
	"context"

	"telerad-core-module/internals/constants"
	"telerad-core-module/internals/entities"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func FindPaginatedStaffAccounts(
	ctx context.Context,
	tx bun.IDB,
	page, pageSize int,
	search string,
	isActive *bool,
) ([]entities.StaffAccountEntity, int, error) {
	var records []entities.StaffAccountEntity

	query := tx.NewSelect().
		Model(&records)

	// Ẩn tài khoản quản trị khỏi danh sách nhân sự. COALESCE để xử lý roles NULL
	// (`<> ALL('{}')` đúng cho mảng rỗng nên dòng roles NULL/rỗng vẫn hiển thị).
	query = query.Where("? <> ALL(COALESCE(?TableAlias.roles, '{}'))", constants.ROLE_ADMIN)

	if search != "" {
		pattern := "%" + search + "%"
		query = query.Where("(?TableAlias.full_name ILIKE ? OR ?TableAlias.code ILIKE ? OR ?TableAlias.username ILIKE ? OR ?TableAlias.phone ILIKE ?)", pattern, pattern, pattern, pattern)
	}

	if isActive != nil {
		query = query.Where("?TableAlias.is_active = ?", *isActive)
	}

	query = query.OrderExpr("?TableAlias.code ASC, ?TableAlias.uuid ASC")

	totalCount, err := findPaginated(ctx, query, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return records, totalCount, nil
}

func FindOneStaffAccountByCode(ctx context.Context, tx bun.IDB, code string, excludeUuid *uuid.UUID) (*entities.StaffAccountEntity, error) {
	var record entities.StaffAccountEntity

	query := tx.NewSelect().
		Model(&record).
		Where("?TableAlias.code ILIKE ?", code)

	if excludeUuid != nil {
		query = query.Where("?TableAlias.uuid <> ?", *excludeUuid)
	}

	return findOne(ctx, query, &record)
}

func GetStaffAccountByUsername(ctx context.Context, tx bun.IDB, username string) (*entities.StaffAccountEntity, error) {
	var record entities.StaffAccountEntity

	query := tx.NewSelect().
		Model(&record).
		Where("?TableAlias.username = ?", username)

	return findOne(ctx, query, &record)
}
