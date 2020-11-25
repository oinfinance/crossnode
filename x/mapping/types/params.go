package types

import "github.com/cosmos/cosmos-sdk/x/params"

// Default parameter namespace
const (
	DefaultParamspace = ModuleName
)

var (
	// key for constant fee parameter
	RefreshPoint      = []byte("RefreshPoint")
	WaitVerifyTimeout = []byte("WaitVerifyTimeout")
)

// type declaration for parameters
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable(
		RefreshPoint, int64(150),
		WaitVerifyTimeout, int64(10000),
	)
}
