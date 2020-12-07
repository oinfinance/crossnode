package sourcechain

import (
	"github.com/oinfinance/crossnode/chain/types"
	"github.com/tendermint/crypto/sha3"
)

type SourceChainEth_Main struct {
	Admin []byte
}

func (t SourceChainEth_Main) ChainId() types.ChainId {
	return types.ChainEth
}

func (t SourceChainEth_Main) MakeReleaseProof(param []byte) []byte {
	s := sha3.New256()
	s.Write(param)
	return s.Sum(nil)
}

func newSourceChain_Eth() SourceChainEth_Main {
	return SourceChainEth_Main{}
}

func GetSourceChain(id types.ChainId) types.SourceChain {
	switch id {
	case types.ChainEth:
		return newSourceChain_Eth()
	default:
		return nil
	}
}
