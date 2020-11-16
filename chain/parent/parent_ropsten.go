package parent

import "github.com/oinfinance/crossnode/chain"

type ParentRopsten struct {
	contractAddr string
	rpclist      []string
	chainId      uint
}

var (
	ethRopsten = ParentRopsten{}
)

func init() {
	// https://ropsten.etherscan.io/token/0x919b1acde3564bcad39efd3fb2955dd98c178e16
	ethRopsten.chainId = 3
	ethRopsten.contractAddr = "0x919b1acde3564bcad39efd3fb2955dd98c178e16"
	ethRopsten.rpclist = []string{
		"http://127.0.0.1:8545",
	}
}
func GetParentRopsten() chain.ParentChain {
	return ethRopsten
}

func (m ParentRopsten) ValidRPC() string {
	for _, link := range m.rpclist {
		if connected(link) {
			return link
		}
	}
	return ""
}

func (m ParentRopsten) ChainId() uint {
	return m.chainId
}

func (m ParentRopsten) ContractAddr() string {
	return m.contractAddr
}
