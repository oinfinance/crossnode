package ethereum

import (
	"github.com/oinfinance/crossnode/bridge"
	"math/big"
)

type EthBridge struct {
}

func (e *EthBridge) GetBalance(addr bridge.Address) *big.Int {
	return big.NewInt(0)
}

func (e *EthBridge) Transfer(from, to bridge.Address, value *big.Int) bridge.Receipt {
	return nil
}

func (e *EthBridge) BatchTransfer(from bridge.Address, to []bridge.Address, value []*big.Int) bridge.Receipts {
	return nil
}
