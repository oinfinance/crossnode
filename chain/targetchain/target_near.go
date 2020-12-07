package targetchain

import (
	"github.com/oinfinance/crossnode/chain/types"
	"github.com/tendermint/crypto/sha3"
)

type TargetChainNear struct {
	Admin []byte
}

func (t TargetChainNear) ChainId() types.ChainId {
	return types.ChainNear
}

func (t TargetChainNear) MakeMintProof(param []byte) []byte {
	s := sha3.New256()
	s.Write(param)
	return s.Sum(nil)
}

func newTargetChain_Near() TargetChainNear {
	return TargetChainNear{}
}

func GetTargetChain(id types.ChainId) types.TargetChain {
	switch id {
	case types.ChainNear:
		return newTargetChain_Near()
	default:
		return nil
	}
}
