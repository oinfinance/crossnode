package parent

import "github.com/oinfinance/crossnode/chain"

type ParentMainnet struct {
	contractAddr string
	rpclist      []string
	chainId      uint
}

var (
	ethMainNet = ParentMainnet{}
)

func init() {
	// https://cn.etherscan.com/token/0x9aeb50f542050172359a0e1a25a9933bc8c01259
	ethMainNet.chainId = 1
	ethMainNet.contractAddr = "0x9aeb50f542050172359a0e1a25a9933bc8c01259"
	ethMainNet.rpclist = []string{
		"http://127.0.0.1:8545",
	}
}
func GetParentMainNet() chain.ParentChain {
	return ethMainNet
}
func connected(link string) bool {
	return true
}

func (m ParentMainnet) ValidRPC() string {
	for _, link := range m.rpclist {
		if connected(link) {
			return link
		}
	}
	return ""
}

func (m ParentMainnet) ChainId() uint {
	return m.chainId
}

func (m ParentMainnet) ContractAddr() string {
	return m.contractAddr
}
