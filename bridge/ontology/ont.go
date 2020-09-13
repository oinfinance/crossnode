package ontology

import (
	"github.com/oinfinance/crossnode/bridge/types"
	"math/big"
)

type OntBridge struct {
}

func (o *OntBridge) GetBalance(addr types.Address) *big.Int {
	return big.NewInt(0)
}

func (o *OntBridge) Transfer(from, to types.Address, value *big.Int) types.Receipt {
	return nil
}

func (o *OntBridge) BatchTransfer(from types.Address, to []types.Address, value []*big.Int) types.Receipts {
	return nil
}
