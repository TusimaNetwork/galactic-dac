package near

//
//import (
//	"fmt"
//	"github.com/stretchr/testify/require"
//	"strings"
//	"testing"
//)
//
//func Test_Defaults(t *testing.T) {
//	res, err := DoCommand("/usr/local/bin/near", "call", "ypenghui7.testnet", "set_greett", `{"greeting":"test667", "greetingValue":"testahowdy666667"}`,
//		"--account-id", "ypenghui7.testnet")
//	require.NoError(t, err)
//	fmt.Println(res)
//
//	/**
//	Receipt: 3JkTTn4KxjZdkfQfzij96GgyuW6Eo5WEXtbDHqqrcvbF
//	        Log [ypenghui7.testnet]: Saving greeting testahowdy
//	        Log [ypenghui7.testnet]: howdy888999jhhlklfds98sfadkliojfsdalk
//	Transaction Id GCGHRKfxWFyEEftV2BCwh2w9vEqadFVo8cpCiDjqkz8a
//	Open the explorer for more info: https://testnet.nearblocks.io/txns/GCGHRKfxWFyEEftV2BCwh2w9vEqadFVo8cpCiDjqkz8a
//	*/
//	idIndex := strings.Index(res, "Transaction Id ")
//
//	fmt.Println("=========================================")
//	fmt.Println(res[idIndex+15 : idIndex+15+44])
//}
