package parent

import (
	"github.com/oinfinance/crossnode/bridge"
	"github.com/oinfinance/crossnode/chain/types"
	"math/big"
)

type ParentRopsten struct {
	contractAddr string
	rpclist      []string
	chainId      types.ChainId
	ethClient    *bridge.EthCluster
}

var (
	ethRopsten = ParentRopsten{}
)

func newEthRopsten_OIN() ParentRopsten {
	// https://ropsten.etherscan.io/token/0x919b1acde3564bcad39efd3fb2955dd98c178e16
	ethRopsten := ParentRopsten{}
	ethRopsten.chainId = types.ChainRopsten
	ethRopsten.contractAddr = "0x919b1acde3564bcad39efd3fb2955dd98c178e16"
	ethRopsten.rpclist = []string{
		"http://127.0.0.1:8545",
	}
	ethRopsten.ethClient = bridge.NewClientCluster(ethRopsten.rpclist)
	return ethRopsten
}

func GetParentRopsten(token types.TokenId) types.ParentChain {
	switch token {
	case types.TokenErc20Oin:
		return newEthRopsten_OIN()
	default:
		panic("unsupported token")
	}
	return nil
}

func (m ParentRopsten) ChainId() types.ChainId {
	return m.chainId
}

func (m ParentRopsten) ContractAddr() string {
	return m.contractAddr
}

func (m ParentRopsten) GetBalance(addr string, blockNumber int64) *big.Int {
	if balance, err := m.ethClient.GetBalance(m.contractAddr, addr, blockNumber); err != nil {
		return balance
	} else {
		return balance
	}
}
