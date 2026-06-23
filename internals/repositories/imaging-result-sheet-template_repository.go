package repositories

import (
	"context"

	"telerad-core-module/internals/entities"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func FindPaginatedImagingResultSheetTemplates(
	ctx context.Context,
	tx bun.IDB,
	page, pageSize int,
	teleradPartnerUuid *uuid.UUID,
	isActive *bool,
) ([]entities.ImagingResultSheetTemplateEntity, int, error) {
	var records []entities.ImagingResultSheetTemplateEntity

	query := tx.NewSelect().
		Model(&records).
		Relation("TeleradPartner")

	if teleradPartnerUuid != nil {
		query = query.Where("?TableAlias.telerad_partner_uuid = ?", *teleradPartnerUuid)
	}

	if isActive != nil {
		query = query.Where("?TableAlias.is_active = ?", *isActive)
	}

	query = query.OrderExpr("telerad_partner.code ASC, ?TableAlias.created_at DESC, ?TableAlias.uuid ASC")

	totalCount, err := findPaginated(ctx, query, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return records, totalCount, nil
}

// FindOneImagingResultSheetTemplateByPartner tìm phiếu kết quả theo CSYT (mỗi CSYT
// chỉ 1 phiếu). excludeUuid để bỏ qua chính bản ghi khi cần.
func FindOneImagingResultSheetTemplateByPartner(ctx context.Context, tx bun.IDB, teleradPartnerUuid uuid.UUID, excludeUuid *uuid.UUID) (*entities.ImagingResultSheetTemplateEntity, error) {
	var record entities.ImagingResultSheetTemplateEntity

	query := tx.NewSelect().
		Model(&record).
		Where("?TableAlias.telerad_partner_uuid = ?", teleradPartnerUuid)

	if excludeUuid != nil {
		query = query.Where("?TableAlias.uuid <> ?", *excludeUuid)
	}

	return findOne(ctx, query, &record)
}

func FindOneImagingResultSheetTemplateWithPartnerByUuid(ctx context.Context, tx bun.IDB, id uuid.UUID) (*entities.ImagingResultSheetTemplateEntity, error) {
	var record entities.ImagingResultSheetTemplateEntity

	query := tx.NewSelect().
		Model(&record).
		Relation("TeleradPartner").
		Where("?TableAlias.uuid = ?", id)

	return findOne(ctx, query, &record)
}
