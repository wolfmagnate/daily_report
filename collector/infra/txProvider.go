package infra

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wolfmagnate/daily_report/collector/infra/db"
)

type Tx interface {
	db.DBTX
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type TxProvider interface {
	BeginTx(ctx context.Context, opts pgx.TxOptions) (Tx, error)
}

type PgxTxProvider struct {
	Pool *pgxpool.Pool
}

var _ TxProvider = (*PgxTxProvider)(nil)

func (p *PgxTxProvider) BeginTx(ctx context.Context, opts pgx.TxOptions) (Tx, error) {
	return p.Pool.BeginTx(ctx, opts)
}
