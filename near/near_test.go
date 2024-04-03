package near

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_GetOffChainData(t *testing.T) {
	t.Log("==================")

	initBlockTimeout := 2 * 15 * time.Second
	ctx, _ := context.WithTimeout(context.Background(), initBlockTimeout)
	//ctx, _ := context.WithCancel(context.Background())

	var (
		User     = "postgres"
		Password = "123"
		Name     = "postgres"
		Host     = "127.0.0.1"
		Port     = "5432"
		MaxConns = 10
	)
	// Prepare DB
	config, err := pgxpool.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?pool_max_conns=%d", User, Password, Host, Port, Name, MaxConns))
	require.NoError(t, err)
	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	require.NoError(t, err)
	storage := NewPostgresStorage(conn)

	near := NearDA{
		db: storage,
	}

	Keys := []string{"0x6E6ea6B5ec0Ee3E1Dc0081FC398284019904eC8B"}
	res, err := near.GetOffChainData(ctx, Keys, nil)
	require.NoError(t, err)

	t.Log("==================", len(res))
}
