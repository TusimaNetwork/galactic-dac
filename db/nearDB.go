package db

import (
	"context"
	"github.com/0xPolygon/cdk-data-availability/types"
	"github.com/jackc/pgx/v4"
)

func (db *DB) StoreNearChainLog(ctx context.Context, od []types.OffChainData, dbTx pgx.Tx) error {
	const storeOffChainDataSQL = `
		INSERT INTO near_cache.offchain_log (key)
		VALUES ($1)
		ON CONFLICT (key) DO NOTHING;
	`
	for _, d := range od {
		if _, err := dbTx.Exec(
			ctx, storeOffChainDataSQL,
			d.Key.Hex(),
		); err != nil {
			return err
		}
	}
	if len(NearChan) <= 0 {
		NearChan <- struct{}{}
	}
	return nil
}
