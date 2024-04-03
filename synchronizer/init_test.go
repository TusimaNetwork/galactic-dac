package synchronizer

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_BlockByNumber(t *testing.T) {
	initBlockTimeout := 2 * 15 * time.Second
	ctx, _ := context.WithTimeout(context.Background(), initBlockTimeout)
	//ctx, _ := context.WithCancel(context.Background())

	t.Log("000000000000000000000000000")

	//"https://eth-sepolia.g.alchemy.com/v2/Lod__GhU0jNhhvpBTma2dKuwZp3Y4-b2"
	eth, err := ethclient.DialContext(ctx, "https://eth-sepolia.g.alchemy.com/v2/Lod__GhU0jNhhvpBTma2dKuwZp3Y4-b2")

	//eth, err := ethclient.DialContext(ctx, "https://sepolia.infura.io/v3/ae5e8b83be5c4b6eb032a18f66d37f87")

	require.NoError(t, err)

	t.Log("111111111111111111111111111")
	latestBlock, err := eth.BlockByNumber(ctx, nil)
	require.NoError(t, err)
	t.Log("222222222222222222222222222:", latestBlock)
}
