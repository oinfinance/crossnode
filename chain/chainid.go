package chain

type ParentChain interface {
	ValidRPC() string
	ChainId() uint
	ContractAddr() string
}

type SourceChain interface {
	ChainId() uint
	MakeReleaseProof(interface{}) []byte
}

type TargetChain interface {
	ChainId() uint
	MakeMintProof(interface{}) []byte
}
