package services

import (
	"context"
	"database/sql"
	"fmt"
	"telerad-core-module/internals/repositories"

	"github.com/uptrace/bun"
	"gorm.io/gorm"
)

var bunNoTransaction = repositories.BunTeleradCore
var gormNoTransaction = repositories.GormTeleradCore

func InitNoTransaction() {
	bunNoTransaction = repositories.BunTeleradCore
	gormNoTransaction = repositories.GormTeleradCore
}

// transactionBun chạy callback trong một bun transaction. Nếu schema != nil, tự động
// SetSearchPath về tenant schema đó ngay đầu transaction (cần cho các Relation join
// trên cluster table); truyền nil khi mọi truy vấn đã chèn schema tường minh.
func transactionBun[T any](
	ctx context.Context,
	schema *string,
	callback func(ctx context.Context, tx bun.Tx) (T, error),
) (T, error) {
	var result T

	err := repositories.BunTeleradCore.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		if schema != nil {
			ctx = repositories.WithSchema(ctx, *schema)
			if err := repositories.SetSearchPath(ctx, tx, *schema); err != nil {
				return err
			}
		}

		res, err := callback(ctx, tx)
		if err != nil {
			return err
		}

		result = res
		return nil
	})

	if err != nil {
		var zero T
		return zero, fmt.Errorf("transaction failed | %s", err.Error())
	}

	return result, nil
}

// transactionGorm — phiên bản GORM của transactionBun. CHỈ dùng cho read-only
// detail/print query cần Preload sâu (deep belongs-to chain mà bun bị truncate alias).
// Mọi mutation/list/transaction phức tạp vẫn dùng transactionBun.
func transactionGorm[T any](
	ctx context.Context,
	schema *string,
	callback func(ctx context.Context, tx *gorm.DB) (T, error),
) (T, error) {
	var result T

	err := repositories.GormTeleradCore.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if schema != nil {
			ctx = repositories.WithSchema(ctx, *schema)
			if err := repositories.SetSearchPathGorm(tx, *schema); err != nil {
				return err
			}
		}

		res, err := callback(ctx, tx)
		if err != nil {
			return err
		}

		result = res
		return nil
	})

	if err != nil {
		var zero T
		return zero, fmt.Errorf("transaction failed | %s", err.Error())
	}

	return result, nil
}
