package types

const (
	QueryReceiptByHash = "queryReceiptByHash"
)

// QueryReceiptByHashParams defines the params to query map info by oin addr.
type QueryReceiptByHashParams struct {
	Hash string
}
