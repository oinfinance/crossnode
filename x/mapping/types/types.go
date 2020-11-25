package types

import (
	"encoding/json"
	"golang.org/x/crypto/sha3"
	"math/big"
)

const (
	MappingWaitVerify   int = iota
	MappingVerifyFailed     // 注册并且verify失败
	MappingVerifyPassed     // 注册并且verify成功
	MappingOutline          // 超过了verify的时间，注册已经无效
)

type MappingInfo struct {
	ErcAddr       string   `json:"erc_addr"`
	CCAddr        string   `json:"cc_addr"`
	RegisterBlock uint64   `json:"register_block"`
	Balance       *big.Int `json:"balance"`
	Status        int      `json:"status"`
}

func MapedKey(info MappingInfo) []byte {
	d := sha3.New256()
	d.Write([]byte(info.ErcAddr))
	d.Write([]byte(info.CCAddr))
	hash := d.Sum(nil)
	return hash
}

type QueryInfoResponse struct {
	Found         bool   `json:"found"`
	ErcAddr       string `json:"ercAddress"`
	CCAddr        string `json:"ccAddress"`
	RegisterBlock uint64 `json:"registerBlock"`
	Balance       uint64 `json:"balance"`
	Status        int    `json:"status"`
}

func (q QueryInfoResponse) String() string {
	d, _ := json.Marshal(q)
	return string(d)
}
