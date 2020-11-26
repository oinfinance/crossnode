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
	TxHash     string   `json:"txHash"`     // 抵押交易hash
	FromChain  int      `json:"fromChain"`  // 原始链
	FromAddr   string   `json:"fromAddr"`   // 抵押者地址
	Token      int      `json:"token"`      // 代币类型
	Value      *big.Int `json:"value"`      // 抵押数量
	ToAddr     string   `json:"toAddr"`     // 接收铸币的地址
	ToChain    int      `json:"toChain"`    // 目标链
	EventType  int      `json:"eventType"`  // 事件类型(1: 铸币 0：销毁)
	AddedBlock uint64   `json:"addedBlock"` // 记录产生的区块号
}

func (c CoinSwapRecord) Hash() []byte {
	h := sha3.New256()
	h.Write([]byte(c.TxHash))
	h.Write(big.NewInt(int64(c.FromChain)).Bytes())
	h.Write([]byte(c.FromAddr))
	h.Write(big.NewInt(int64(c.Token)).Bytes())
	h.Write(c.Value.Bytes())
	h.Write([]byte(c.ToAddr))
	h.Write(big.NewInt(int64(c.ToChain)).Bytes())
	h.Write(big.NewInt(int64(c.AddedBlock)).Bytes())
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
