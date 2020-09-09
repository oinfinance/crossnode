package bridge

import (
	"github.com/oinfinance/crossnode/bridge/ethereum"
	"github.com/oinfinance/crossnode/bridge/ontology"
	"math/big"
)

type Address []byte
type Receipt []byte
type Receipts []Receipt

/*
 * Bridge define the interface that could query account's assert and do a transfer.
 */
type Bridge interface {
	GetBalance(account Address) *big.Int
	Transfer(from, to Address, value *big.Int) Receipt
	BatchTransfer(from Address, to []Address, value []*big.Int) Receipts
}

func NewBridge(chainId uint) Bridge {
	switch chainId {
	case 0x01:
		return &ethereum.EthBridge{}
	case 0x02:
		return &ontology.OntBridge{}
	}
}
