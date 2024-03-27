package near

import (
	"context"
	"errors"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var ErrStateNotFound = errors.New("not found")

type PostgresStorage struct {
	*pgxpool.Pool
}

func NewPostgresStorage(db *pgxpool.Pool) *PostgresStorage {
	return &PostgresStorage{
		db,
	}
}

type execQuerier interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

func (p *PostgresStorage) getExecQuerier(dbTx pgx.Tx) execQuerier {
	if dbTx != nil {
		return dbTx
	}
	return p
}

func (n *NearDA) StoreNearStateLog(ctx context.Context, logId int, tx string, dbTx pgx.Tx) error {
	//if dbTx, err = s.db.BeginStateTransaction(ctx); err != nil {
	//	log.Errorf("Starts checking whether data needs to be uploaded to near DA")
	//	return
	//}
	//
	//if err = dbTx.Commit(ctx); err != nil {
	//	log.Errorf("Starts checking whether data needs to be uploaded to near DA")
	//	return
	//}
	//

	e := n.db.getExecQuerier(dbTx)

	const storeOffChainDataSQL = `
		INSERT INTO near_cache.state_log (log_id, tx)
		VALUES ($1, $2)
		ON CONFLICT (log_id) DO NOTHING;
	`
	if _, err := e.Exec(
		ctx, storeOffChainDataSQL,
		logId,
		tx,
	); err != nil {
		return err
	}

	return nil
}

func (n *NearDA) GetNearStateLog(ctx context.Context, dbTx pgx.Tx) (int, error) {
	e := n.db.getExecQuerier(dbTx)

	const getOffchainDataSQL = `
			SELECT log_id
			FROM near_cache.state_log
			WHERE id = (SELECT MAX(id) FROM near_cache.state_log);;
	`
	var (
		logId int
	)

	if err := e.QueryRow(ctx, getOffchainDataSQL).Scan(&logId); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, ErrStateNotFound
		}
		return 0, err
	}
	return logId, nil
}

type ResData struct {
	Key   string
	Value string
}

// GetOffChainData returns the value identified by the key
func (n *NearDA) GetOffChainData(ctx context.Context, keys []string, dbTx pgx.Tx) ([]ResData, error) {
	e := n.db.getExecQuerier(dbTx)

	getOffchainDataSQL := `SELECT * FROM data_node.offchain_data WHERE key = ANY($1::text[])`

	hexValue := make([]ResData, 0)
	rows, err := e.Query(ctx, getOffchainDataSQL, keys)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrStateNotFound
		}
		return nil, err
	}

	// 遍历结果集
	for rows.Next() {
		var value string
		var key string

		err = rows.Scan(&key, &value)
		if err != nil {
			return nil, err
		}
		hexValue = append(hexValue, ResData{Value: value, Key: key})
	}
	return hexValue, nil
}

type Res struct {
	Id  int
	Key string
}

func (n *NearDA) GetNearChainLog(ctx context.Context, id int, dbTx pgx.Tx) ([]Res, error) {
	e := n.db.getExecQuerier(dbTx)

	const getOffchainDataSQL = `
		SELECT *
		FROM near_cache.offchain_log
		WHERE id > $1
		Order by id;
	`

	logId := make([]Res, 0)

	rows, err := e.Query(ctx, getOffchainDataSQL, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrStateNotFound
		}
		return nil, err
	}

	// 遍历结果集
	for rows.Next() {
		var id int
		var key string

		err = rows.Scan(&id, &key)
		if err != nil {
			return nil, err
		}
		logId = append(logId, Res{Id: id, Key: key})
	}
	return logId, nil
}
