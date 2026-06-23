package repositories

import (
	"context"

	"telerad-core-module/internals/entities"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func FindPaginatedImagingResultTemplates(
	ctx context.Context,
	tx bun.IDB,
	page, pageSize int,
	modality string,
	search string,
	isActive *bool,
	bodyParts []string,
) ([]entities.ImagingResultTemplateEntity, int, error) {
	var records []entities.ImagingResultTemplateEntity

	query := tx.NewSelect().
		Model(&records)

	if modality != "" {
		query = query.Where("?TableAlias.modality = ?", modality)
	}

	if search != "" {
		pattern := "%" + search + "%"
		query = query.Where("?TableAlias.name ILIKE ?", pattern)
	}

	if isActive != nil {
		query = query.Where("?TableAlias.is_active = ?", *isActive)
	}

	// Lọc theo bộ phận chụp: mẫu phải trùng ÍT NHẤT 1 body_part (toán tử overlap &&).
	if len(bodyParts) > 0 {
		query = query.Where("?TableAlias.body_parts && ?", pgdialect.Array(bodyParts))
	}

	query = query.OrderExpr("?TableAlias.display_order ASC NULLS LAST, ?TableAlias.name ASC, ?TableAlias.uuid ASC")

	totalCount, err := findPaginated(ctx, query, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return records, totalCount, nil
}
