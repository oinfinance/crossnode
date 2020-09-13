package types

const (
	QueryInfoByOin   = "queryInfoByOin" // QueryFeed
	QueryInfoByLocal = "queryInfoByLocal"
)

// QueryInfoByOinParams defines the params to query map info by oin addr.
type QueryInfoByOinParams struct {
	OinAddr string
}

// QueryInfoByLocalParams defines the params to query map info by crossnode addr.
type QueryInfoByLocalParams struct {
	LocalAddr string
}
