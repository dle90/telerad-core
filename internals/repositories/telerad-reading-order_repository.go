package repositories

import (
	"context"

	"telerad-core-module/internals/entities"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// FindOneReadingOrderByPartnerAndOrderItemId dùng để chống trùng: 1 đối tác không
// được đẩy lặp cùng một order_item_id.
func FindOneReadingOrderByPartnerAndOrderItemId(ctx context.Context, tx bun.IDB, teleradPartnerUuid uuid.UUID, orderItemId string) (*entities.TeleradReadingOrderEntity, error) {
	var record entities.TeleradReadingOrderEntity

	query := tx.NewSelect().
		Model(&record).
		Where("?TableAlias.telerad_partner_uuid = ?", teleradPartnerUuid).
		Where("?TableAlias.order_item_id = ?", orderItemId)

	return findOne(ctx, query, &record)
}
