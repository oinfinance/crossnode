package chain

import "github.com/oinfinance/crossnode/chain/types"

type ParentConfig struct {
	ParentId types.ChainId
	TokenId  types.TokenId
}

var (
	parentConfig = ParentConfig{ParentId: types.ChainRopsten, TokenId: types.TokenErc20Oin}
)
