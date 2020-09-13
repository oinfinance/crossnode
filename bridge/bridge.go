package bridge

import (
	"github.com/oinfinance/crossnode/bridge/ethereum"
	"github.com/oinfinance/crossnode/bridge/ontology"
	"github.com/oinfinance/crossnode/bridge/types"
	"math/big"
)


/*
 * Bridge define the interface that could query account's assert and do a transfer.
 */
type Bridge interface {
	GetBalance(account types.Address) *big.Int
	Transfer(from, to types.Address, value *big.Int) types.Receipt
	BatchTransfer(from types.Address, to []types.Address, value []*big.Int) types.Receipts
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