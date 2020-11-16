package chain

type ParentChain interface {
	ValidRPC() string
	ChainId() uint
	ContractAddr() string
}
