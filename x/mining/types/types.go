package types

import (
	"encoding/json"
	"math/big"
)

type RewardsInfo struct {
	Rewards	big.Int		// 未提现收益
}

type QueryRewardsResponse struct {
	Rewards      uint64   `json:"rewards"`
}

func (q QueryRewardsResponse)String() string{
	d,_ := json.Marshal(q)
	return string(d)
}
