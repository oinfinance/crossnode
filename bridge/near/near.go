package near

import (
	"encoding/hex"
	"github.com/oinfinance/crossnode/bridge/types"
	"math/big"
)

var (
	rpcaddr = "http://127.0.0.1:8545"
)

type NearBridge struct {
	cluster *NearCluster
}

func NewNearBridge() *NearBridge {
	return &NearBridge{NewClientCluster([]string{rpcaddr})}
}

func (e *NearBridge) GetBalance(addr types.Address) *big.Int {
	balance, _ := e.cluster.GetBalance("", addr[:])
	return balance

}

func (e *NearBridge) Transfer(from, to types.Address, value *big.Int) types.Receipt {
	txhash, err := e.cluster.SendTransaction("", hex.EncodeToString(from), hex.EncodeToString(to), value)
	if err != nil {
		return types.Receipt{}
	}
	return types.Receipt(txhash)
}

func (e *NearBridge) BatchTransfer(from types.Address, to []types.Address, value []*big.Int) types.Receipts {
	return nil
}
