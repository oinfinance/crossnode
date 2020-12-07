package types

import "math/big"

type ChainId uint
type TokenId uint

const (
	ChainStart ChainId = iota + 0x00
	ChainEth
	ChainRopsten
	ChainNear

	ChainUnknown = 0x99
)
const (
	TokenStart TokenId = iota + 0x10000
	TokenErc20Oin
	TokenEth

	TokenUnknown = 0x10099
)

type ParentChain interface {
	ChainId() ChainId
	ContractAddr() string
	GetBalance(addr string, blockNumber int64) *big.Int
}

type SourceChain interface {
	ChainId() ChainId
	MakeReleaseProof([]byte) []byte
}

type TargetChain interface {
	ChainId() ChainId
	MakeMintProof([]byte) []byte
}
