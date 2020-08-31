package ontology

import (
	"github.com/oinfinance/crossnode/bridge"
	"math/big"
)

type OntBridge struct {
}

func (o *OntBridge) GetBalance(addr bridge.Address) *big.Int {
	return big.NewInt(0)
}

func (o *OntBridge) Transfer(from, to bridge.Address, value *big.Int) bridge.Receipt {
	return nil
}

func (o *OntBridge) BatchTransfer(from bridge.Address, to []bridge.Address, value []*big.Int) bridge.Receipts {
	return nil
}
