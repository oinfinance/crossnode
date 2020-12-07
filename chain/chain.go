package chain

import (
	"github.com/oinfinance/crossnode/chain/parent"
	"github.com/oinfinance/crossnode/chain/types"
)

func GetParentChain() types.ParentChain {
	switch parentConfig.ParentId {
	case types.ChainRopsten:
		return parent.GetParentRopsten(parentConfig.TokenId)
	case types.ChainEth:
		return parent.GetParentMainNet(parentConfig.TokenId)
	default:
		return nil
	}
}
