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

func NewBridge(chainid uint) Bridge {
	switch ChainId(chainid) {
	case ChainEth:
		return ethereum.NewEthBridge()
	case ChainONT:
		return &ontology.OntBridge{}
	}
	return nil
}

func SupportedByName(chain string, token string) bool {
	chainId,exChain := SupportChainId[chain]
	if !exChain {
		return false
	}
	tid, extoken := SupportTokenId[token]
	if !extoken {
		return false
	}

	tokenlist := SupportList[chainId]
	for _,info := range tokenlist {
		if info.id == tid {
			return true
		}
	}
	return false
}