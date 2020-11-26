package bridge

import "github.com/oinfinance/crossnode/bridge/ethereum"

type ChainId uint
type TokenId uint

type Token struct {
	id       TokenId
	contract string
}

const (
	ChainStart ChainId = iota + 0x00
	ChainEth
	ChainONT
	ChainNear
)
const (
	TokenStart TokenId = iota + 0x10000
	TokenErc20Oin
)

var SupportList = map[ChainId][]Token{
	ChainEth: {{TokenErc20Oin, ethereum.OinToken}},
}

var SupportChainName = map[ChainId]string{
	ChainEth:  "ethereum",
	ChainONT:  "ontology",
	ChainNear: "near",
}

var SupportChainId = map[string]ChainId{
	"eth":      ChainEth,
	"ETH":      ChainEth,
	"ethereum": ChainEth,
	"ont":      ChainONT,
	"ontology": ChainONT,
	"ONT":      ChainONT,
	"near":     ChainNear,
	"NEAR":     ChainNear,
}

var SupportTokenName = map[TokenId]string{
	TokenErc20Oin: "OIN ERC20",
}

var SupportTokenId = map[string]TokenId{
	"oin":       TokenErc20Oin,
	"Oin":       TokenErc20Oin,
	"OIN":       TokenErc20Oin,
	"ERC20-OIN": TokenErc20Oin,
	"ERC20-Oin": TokenErc20Oin,
}

func TokenIdByName(name string) TokenId {
	return SupportTokenId[name]
}

func ChainIdByName(name string) ChainId {
	return SupportChainId[name]
}
