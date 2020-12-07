package chain

import (
	. "github.com/oinfinance/crossnode/chain/types"
)

type Group struct {
	SourceChain ChainId
	TargetChain ChainId
	Token       TokenId
}

var SupportList = []Group{
	{
		SourceChain: ChainEth,
		Token:       TokenEth,
		TargetChain: ChainNear,
	},
	{
		SourceChain: ChainRopsten,
		Token:       TokenEth,
		TargetChain: ChainNear,
	},
}

var SupportChainName = map[ChainId]string{
	ChainEth:     "ethereum",
	ChainRopsten: "ropsten",
	ChainNear:    "near",
}

var SupportChainId = map[string]ChainId{
	"eth":      ChainEth,
	"ETH":      ChainEth,
	"ethereum": ChainEth,
	"ropsten":  ChainRopsten,
	"Ropsten":  ChainRopsten,
	"near":     ChainNear,
	"NEAR":     ChainNear,
	"Near":     ChainNear,
}

var SupportTokenName = map[TokenId]string{
	TokenErc20Oin: "OIN ERC20",
	TokenEth:      "ETH",
}

var SupportTokenId = map[string]TokenId{
	"oin":       TokenErc20Oin,
	"Oin":       TokenErc20Oin,
	"OIN":       TokenErc20Oin,
	"ERC20-OIN": TokenErc20Oin,
	"ERC20-Oin": TokenErc20Oin,
	"eth":       TokenEth,
	"ETH":       TokenEth,
	"Eth":       TokenEth,
}

func TokenIdByName(name string) TokenId {
	if id, exist := SupportTokenId[name]; exist {
		return id
	} else {
		return TokenUnknown
	}
}

func ChainIdByName(name string) ChainId {
	if id, exist := SupportChainId[name]; exist {
		return id
	} else {
		return ChainUnknown
	}
}

func SupportedGroup(sourceChain, targetChain, tokenName string) bool {
	sc := ChainIdByName(sourceChain)
	tc := ChainIdByName(targetChain)
	token := TokenIdByName(tokenName)
	for _, g := range SupportList {
		if sc == g.SourceChain && tc == g.TargetChain && token == g.Token {
			return true
		}
	}
	return false
}
