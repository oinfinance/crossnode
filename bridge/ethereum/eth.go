package ethereum

import (
	"encoding/hex"
	"github.com/oinfinance/crossnode/bridge"
	"math/big"
)

var (
	rpcaddr = "http://127.0.0.1:8545"
)

type EthBridge struct {
	cluster *EthCluster
}

func NewEthBridge() bridge.Bridge {
	return &EthBridge{NewClientCluster([]string{rpcaddr})}
}

func (e *EthBridge) GetBalance(addr bridge.Address) *big.Int {
	balance, _ := e.cluster.GetBalance("", hex.EncodeToString(addr))
	return balance

}

func (e *EthBridge) Transfer(from, to bridge.Address, value *big.Int) bridge.Receipt {
	txhash, err := e.cluster.SendTransaction("", hex.EncodeToString(from), hex.EncodeToString(to), value)
	if err != nil {
		return bridge.Receipt{}
	}
	return bridge.Receipt(txhash)
}

func (e *EthBridge) BatchTransfer(from bridge.Address, to []bridge.Address, value []*big.Int) bridge.Receipts {
	return nil
}
