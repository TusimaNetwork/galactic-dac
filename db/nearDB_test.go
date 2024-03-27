package db

import (
	"context"
	"github.com/0xPolygon/cdk-data-availability/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_StoreNearChainLog(t *testing.T) {
	t.Log("==================")

	initBlockTimeout := 2 * 15 * time.Second
	ctx, _ := context.WithTimeout(context.Background(), initBlockTimeout)
	//ctx, _ := context.WithCancel(context.Background())

	tmp := Config{

		User:      "postgres",
		Password:  "123",
		Name:      "postgres",
		Host:      "127.0.0.1",
		Port:      "5432",
		EnableLog: false,
		MaxConns:  10,
	}
	// Prepare DB
	pg, err := NewSQLDB(tmp)
	require.NoError(t, err)
	storage := New(pg)

	pg2, err := storage.BeginStateTransaction(ctx)
	require.NoError(t, err)

	datas := make([]types.OffChainData, 0)
	da1 := types.OffChainData{
		Key:   common.HexToHash("0x6E6ea6B5ec0Ee3E1Dc0081FC398284019904eC8B"),
		Value: []byte("aaa"),
	}
	datas = append(datas, da1)

	err = storage.StoreNearChainLog(ctx, datas, pg2)
	require.NoError(t, err)

	err = pg2.Commit(ctx)
	require.NoError(t, err)
	t.Log("==================")
}
