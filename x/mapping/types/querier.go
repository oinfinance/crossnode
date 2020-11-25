package types

const (
	QueryInfoByErc = "queryInfoByErc" // QueryFeed
	QueryInfoByCC  = "queryInfoByCC"
)

// QueryInfoByErcParams defines the params to query map info by oin erc20 addr.
type QueryInfoByErcParams struct {
	ErcAddr string
}

// QueryInfoByCCParams defines the params to query map info by crossnode addr.
type QueryInfoByCCParams struct {
	CCAddr string
}
