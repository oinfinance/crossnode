package types

const (
	QueryRewards   = "queryRewards"
)

// QueryInfoByOinParams defines the params to query map info by oin addr.
type QueryRewardsParams struct {
	LocalAddr string
}
