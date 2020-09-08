package types

import "github.com/cosmos/cosmos-sdk/x/params"

// Default parameter namespace
const (
	DefaultParamspace = ModuleName
)

var (
	// key for constant fee parameter
	ParamStoreKeyTokenCount = []byte("TokenTypeCount")
)

// type declaration for parameters
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable(
		ParamStoreKeyTokenCount, int32(0),
	)
}
