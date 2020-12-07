package parent

import (
	"github.com/oinfinance/crossnode/bridge"
	"github.com/oinfinance/crossnode/chain/types"
	"math/big"
)

type ParentMainnet struct {
	contractAddr string
	rpclist      []string
	chainId      types.ChainId
	ethClient    *bridge.EthCluster
}

func newEthMainNet_OIN() ParentMainnet {
	// https://cn.etherscan.com/token/0x9aeb50f542050172359a0e1a25a9933bc8c01259
	ethMainNet := ParentMainnet{}
	ethMainNet.chainId = types.ChainEth
	ethMainNet.contractAddr = "0x9aeb50f542050172359a0e1a25a9933bc8c01259"
	ethMainNet.rpclist = []string{
		"http://127.0.0.1:8545",
	}
	ethMainNet.ethClient = bridge.NewClientCluster(ethMainNet.rpclist)
	return ethMainNet
}

func GetParentMainNet(token types.TokenId) types.ParentChain {
	switch token {
	case types.TokenErc20Oin:
		return newEthMainNet_OIN()
	default:
		panic("unsupported token")
	}
	return nil
}

func (m ParentMainnet) ChainId() types.ChainId {
	return m.chainId
}

func (m ParentMainnet) ContractAddr() string {
	return m.contractAddr
}

func (m ParentMainnet) GetBalance(addr string, blockNumber int64) *big.Int {
	if balance, err := m.ethClient.GetBalance(m.contractAddr, addr, blockNumber); err != nil {
		return balance
	} else {
		return balance
	}
}
