package types

import (
	"encoding/json"
	"github.com/tendermint/crypto/sha3"
	"math/big"
)

const (
	RecordStatusWaited  = 1
	RecordStatusSucceed = 2
	RecordStatusFailed  = 3
)

type CoinSwapRecord struct {
	TxHash     []byte   `json:"txHash"`     // 抵押交易hash
	FromChain  int      `json:"fromChain"`  // 原始链
	FromAddr   []byte   `json:"fromAddr"`   // 抵押者地址
	Token      int      `json:"token"`      // 代币类型
	Value      *big.Int `json:"value"`      // 抵押数量
	ToAddr     []byte   `json:"toAddr"`     // 接收铸币的地址
	ToChain    int      `json:"toChain"`    // 目标链
	EventType  int      `json:"eventType"`  // 事件类型(1: 铸币 0：销毁)
	AddedBlock *big.Int `json:"addedBlock"` // 记录产生的区块号
}

func (c CoinSwapRecord) Hash() []byte {
	h := sha3.New256()
	h.Write(c.TxHash)
	h.Write(big.NewInt(int64(c.FromChain)).Bytes())
	h.Write(c.FromAddr)
	h.Write(big.NewInt(int64(c.Token)).Bytes())
	h.Write(c.Value.Bytes())
	h.Write(c.ToAddr)
	h.Write(big.NewInt(int64(c.ToChain)).Bytes())
	h.Write(c.AddedBlock.Bytes())
	return h.Sum([]byte{})
}
func (c CoinSwapRecord) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// 代币兑换交易收条，用于返回一段凭据，用户使用凭据去目标链提取兑换的币
type CoinSwapReceipt struct {
	Status  int    `json:"status"`  // 兑换状态(1:wait, 2:finished, 0:invalid)
	Receipt string `json:"receipt"` // 用户提币凭据
}

func (c CoinSwapReceipt) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

type CoinSwapRecordStorage struct {
	Record  CoinSwapRecord  `json:"record"`
	Receipt CoinSwapReceipt `json:"receipt"`
}

func (c CoinSwapRecordStorage) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}
