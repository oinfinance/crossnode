package types

import (
	"math/big"
)

type Object struct {
	Coins map[int]big.Int // chain/token ==> amount
}

type State struct {
	Objects map[string]Object // all user object
}

// Deposit Receipt 记录用户在 链A上的抵押记录
type MakeReceipt struct {
	TxHash []byte   `json:"txhash"` // 抵押交易hash
	From   []byte   `json:"from"`   // 抵押者地址
	Value  *big.Int `json:"value"`  // 抵押数量
	To     []byte   `json:"to"`     // 接收铸币的地址
	Token  int      `json:"token"`  // 代币类型
	Chain  int      `json:"chain"`  // 链类型
}
