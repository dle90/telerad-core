package configs

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostGresDB struct {
	Bun  *bun.DB       // High-level API (ORM) — dùng cho list/mutation/transaction
	Gorm *gorm.DB      // Read-only ORM — CHỈ dùng cho detail/print query cần deep relation (Preload sâu)
	Pool *pgxpool.Pool // Low-level API (raw pgx)
}

var TeleradCorePostgresDatabase *PostGresDB

func ConnectDB() {
	// Postgres Connection
	config, err := pgxpool.ParseConfig(EnvTeleradCorePostgresURI())
	if err != nil {
		panic(err)
	}

	config.ConnConfig.ConnectTimeout = GetTeleradCorePostgresConnectTimeout()
	config.MaxConns = GetTeleradCorePostgresMaxConns()
	config.MinConns = GetTeleradCorePostgresMinConns()

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}

	// Wrap pgxpool.Pool into bun.DB
	sqldb := stdlib.OpenDBFromPool(pool)
	bunDB := bun.NewDB(sqldb, pgdialect.New())
	bunDB.AddQueryHook(bundebug.NewQueryHook()) // optional logging

	// Khởi tạo GORM dùng chung *sql.DB với bun → 1 connection pool duy nhất.
	// Tắt magic của GORM: bun đã lo transaction/migration/timestamp.
	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: sqldb}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	TeleradCorePostgresDatabase = &PostGresDB{
		Bun:  bunDB,
		Gorm: gormDB,
		Pool: pool,
	}

	fmt.Println("Connected to Telerad Core Postgres Database")
}
