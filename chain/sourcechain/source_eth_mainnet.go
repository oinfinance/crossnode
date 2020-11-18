package sourcechain

type SourceChainEthMainnet struct {
}

var (
	ethMainnetSourceChain = SourceChainEthMainnet{}
)

func init() {

}

func GetSourceChainEthMainNet() SourceChainEthMainnet {
	return ethMainnetSourceChain
}
