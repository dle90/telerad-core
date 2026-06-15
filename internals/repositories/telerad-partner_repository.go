package repositories

import (
	"context"

	"telerad-core-module/internals/entities"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

// FindTeleradPartnersForReading lấy đối tác cho cây màn "Đọc ca" của user thường:
// uuid nằm trong danh sách được phân (partnerUuids) VÀ có ít nhất 1 loại chụp giao
// với modalities được phân (toán tử overlap mảng `&&` của Postgres). Đẩy lọc xuống
// DB thay vì load tất cả rồi lọc in-memory.
func FindTeleradPartnersForReading(
	ctx context.Context,
	tx bun.IDB,
	isAdmin bool,
	partnerUuids []uuid.UUID,
	modalities []string,
) ([]entities.TeleradPartnerEntity, error) {
	records := []entities.TeleradPartnerEntity{}

	query := tx.NewSelect().
		Model(&records).
		OrderExpr("?TableAlias.code ASC, ?TableAlias.uuid ASC")

	if !isAdmin {
		if len(partnerUuids) == 0 || len(modalities) == 0 {
			return records, nil
		}
		query = query.Where("?TableAlias.uuid IN (?)", bun.List(partnerUuids)).
			Where("?TableAlias.modalities && ?", pgdialect.Array(modalities))
	}

	if err := query.Scan(ctx); err != nil {
		return nil, err
	}

	return records, nil
}

func FindPaginatedTeleradPartners(
	ctx context.Context,
	tx bun.IDB,
	page, pageSize int,
	search string,
	isActive *bool,
) ([]entities.TeleradPartnerEntity, int, error) {
	var records []entities.TeleradPartnerEntity

	query := tx.NewSelect().
		Model(&records)

	if search != "" {
		pattern := "%" + search + "%"
		query = query.Where("(?TableAlias.name ILIKE ? OR ?TableAlias.code ILIKE ? OR ?TableAlias.username ILIKE ?)", pattern, pattern, pattern)
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

func FindAllTeleradPartners(ctx context.Context, tx bun.IDB) ([]entities.TeleradPartnerEntity, error) {
	var records []entities.TeleradPartnerEntity

	err := tx.NewSelect().
		Model(&records).
		OrderExpr("?TableAlias.code ASC, ?TableAlias.uuid ASC").
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return records, nil
}

func FindOneTeleradPartnerByCode(ctx context.Context, tx bun.IDB, code string, excludeUuid *uuid.UUID) (*entities.TeleradPartnerEntity, error) {
	var record entities.TeleradPartnerEntity

	query := tx.NewSelect().
		Model(&record).
		Where("?TableAlias.code ILIKE ?", code)

	if excludeUuid != nil {
		query = query.Where("?TableAlias.uuid <> ?", *excludeUuid)
	}

	return findOne(ctx, query, &record)
}

func FindOneTeleradPartnerByUsername(ctx context.Context, tx bun.IDB, username string, excludeUuid *uuid.UUID) (*entities.TeleradPartnerEntity, error) {
	var record entities.TeleradPartnerEntity

	query := tx.NewSelect().
		Model(&record).
		Where("?TableAlias.username = ?", username)

	if excludeUuid != nil {
		query = query.Where("?TableAlias.uuid <> ?", *excludeUuid)
	}

	return findOne(ctx, query, &record)
}
