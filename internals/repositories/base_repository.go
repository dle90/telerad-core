package repositories

import (
	"context"
	"database/sql"
	"errors"
	"telerad-core-module/configs"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/uptrace/bun"
	"gorm.io/gorm"
)

var BunTeleradCore *bun.DB
var GormTeleradCore *gorm.DB
var PoolTeleradCore *pgxpool.Pool

const SchemaMaster = "master"

// Table* và Sequence* constants được khai báo ở từng file repository theo domain
// (vd diagnostic-imaging-study_repository.go). Skeleton chưa có domain nên để trống.

func InitDatabase() {
	BunTeleradCore = configs.TeleradCorePostgresDatabase.Bun
	GormTeleradCore = configs.TeleradCorePostgresDatabase.Gorm
	PoolTeleradCore = configs.TeleradCorePostgresDatabase.Pool

	// Đăng ký các junction model (m2m) khi domain được thêm vào, ví dụ:
	// BunTeleradCore.RegisterModel((*clusterEntities.XxxEntity)(nil))
}

func WithSchemaTableSelect(schema, tableName string) func(*bun.SelectQuery) *bun.SelectQuery {
	return func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.ModelTableExpr("?.? AS ?TableAlias", bun.Ident(schema), bun.Ident(tableName))
	}
}

// WithSchemaTableUpdate giống WithSchemaTable nhưng dùng cho NewUpdate (UpdateQuery có
// chữ ký Apply khác SelectQuery nên cần helper riêng).
func WithSchemaTableUpdate(schema, tableName string) func(*bun.UpdateQuery) *bun.UpdateQuery {
	return func(q *bun.UpdateQuery) *bun.UpdateQuery {
		return q.ModelTableExpr("?.? AS ?TableAlias", bun.Ident(schema), bun.Ident(tableName))
	}
}

///////////////////////////////// Schema trong context /////////////////////////////////

type contextKey string

const schemaContextKey contextKey = "tenant_schema"

// WithSchema gắn tenant schema vào context.
func WithSchema(ctx context.Context, schema string) context.Context {
	return context.WithValue(ctx, schemaContextKey, schema)
}

// SchemaFromContext lấy tenant schema đã gắn trong context (rỗng nếu chưa set).
func SchemaFromContext(ctx context.Context) string {
	schema, _ := ctx.Value(schemaContextKey).(string)
	return schema
}

///////////////////////////////// Các hàm dùng chung cấp 1 /////////////////////////////////

// SetSearchPath đặt search_path cho transaction hiện tại thành tenant schema + master + public.
// Postgres sẽ tự resolve các unqualified table name theo thứ tự trong search_path,
// nhờ vậy bun's belongs-to / has-one Relation join sẽ hoạt động đúng với cluster tables
// mà không cần Apply(WithSchemaTableSelect(...)) ở từng query.
//
// LƯU Ý: phải gọi BÊN TRONG transaction (bun.Tx). Nếu gọi trên connection pool
// (bunNoTransaction), SET LOCAL không có tác dụng và Postgres sẽ chỉ log warning.
// Khi transaction COMMIT/ROLLBACK, search_path tự reset → connection trả về pool sạch.
func SetSearchPath(ctx context.Context, tx bun.IDB, schema string) error {
	_, err := tx.NewRaw("SET LOCAL search_path TO ?", bun.Ident(schema)).Exec(ctx)

	return err
}

// SetSearchPathGorm — phiên bản GORM của SetSearchPath. Cùng cơ chế (SET LOCAL
// trong transaction), nhưng tự quote identifier bằng pgx.Identifier vì GORM
// không có sẵn helper escape identifier như bun.Ident.
//
// LƯU Ý: phải gọi BÊN TRONG transaction (gorm tx). SET LOCAL trên connection pool
// sẽ bị Postgres ignore (chỉ log warning).
func SetSearchPathGorm(tx *gorm.DB, schema string) error {
	quoted := pgx.Identifier{schema}.Sanitize()
	return tx.Exec("SET LOCAL search_path TO " + quoted).Error
}

func findPaginated(ctx context.Context, query *bun.SelectQuery, page int, pageSize int) (int, error) {
	if page < 1 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	} else if pageSize > 100 {
		pageSize = 100
	}

	offset := (page - 1) * pageSize
	query.Limit(pageSize).Offset(offset)

	return query.ScanAndCount(ctx)
}

func InsertOneWithSchema(ctx context.Context, tx bun.IDB, schema, tableName string, model any) error {
	_, err := tx.NewInsert().
		Model(model).
		ModelTableExpr("?.?", bun.Ident(schema), bun.Ident(tableName)).
		Exec(ctx)
	return err
}

func InsertOne(ctx context.Context, tx bun.IDB, model any) error {
	_, err := tx.NewInsert().
		Model(model).
		Exec(ctx)
	return err
}

func InsertManyWithSchema(ctx context.Context, tx bun.IDB, schema, tableName string, models any) error {
	_, err := tx.NewInsert().
		Model(models).
		ModelTableExpr("?.?", bun.Ident(schema), bun.Ident(tableName)).
		Exec(ctx)
	return err
}

func InsertMany(ctx context.Context, tx bun.IDB, model any) error {
	_, err := tx.NewInsert().
		Model(model).
		Exec(ctx)
	return err
}

func FindOneByUuidWithSchema[T any](ctx context.Context, tx bun.IDB, schema, tableName string, id uuid.UUID) (*T, error) {
	var record T
	err := tx.NewSelect().
		Model(&record).
		Apply(WithSchemaTableSelect(schema, tableName)).
		Where("?TableAlias.uuid = ?", id).
		Limit(1).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &record, nil
}

func FindManyByUuidsWithSchema[T any](ctx context.Context, tx bun.IDB, schema, tableName string, ids []uuid.UUID) ([]T, error) {
	records := []T{}
	if len(ids) == 0 {
		return records, nil
	}

	err := tx.NewSelect().
		Model(&records).
		Apply(WithSchemaTableSelect(schema, tableName)).
		Where("?TableAlias.uuid IN (?)", bun.List(ids)).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return records, nil
		}
		return nil, err
	}
	return records, nil
}

func FindOneByUuid[T any](ctx context.Context, tx bun.IDB, id uuid.UUID) (*T, error) {
	var record T
	err := tx.NewSelect().
		Model(&record).
		Where("?TableAlias.uuid = ?", id).
		Limit(1).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &record, nil
}

func FindOneById[T any](ctx context.Context, tx bun.IDB, id int64) (*T, error) {
	var record T
	err := tx.NewSelect().
		Model(&record).
		Where("?TableAlias.id = ?", id).
		Limit(1).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &record, nil
}

func UpdateWholeRecordWithSchema(ctx context.Context, tx bun.IDB, schema, tableName string, model any) error {
	_, err := tx.NewUpdate().
		Model(model).
		Apply(WithSchemaTableUpdate(schema, tableName)).
		WherePK().
		Exec(ctx)
	return err
}

func UpdateWholeRecord(ctx context.Context, tx bun.IDB, model any) error {
	_, err := tx.NewUpdate().
		Model(model).
		WherePK().
		Exec(ctx)
	return err
}

// findOne scans the query into dest, which MUST be the same model instance
// already bound to the query via .Model(dest). It deliberately does NOT call
// .Model() again: re-binding rebuilds bun's tableModel and drops deferred m2m
// relations (their post-scan query is lost), so re-binding would silently
// return empty m2m slices while keeping belongs-to/has-one joins working.
func findOne[T any](ctx context.Context, query *bun.SelectQuery, dest *T) (*T, error) {
	if err := query.Limit(1).Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return dest, nil
}
