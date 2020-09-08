package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oinfinance/crossnode/x/coinswap/types"
)

// Keeper store user info.

type Keeper struct {
	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context
	cdc      *codec.Codec // The wire codec for binary encoding/decoding.

	// subspace
}

// NewKeeper returns a guardian keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	keeper := Keeper{
		storeKey: key,
		cdc:      cdc,
	}
	return keeper
}

func (k Keeper) AddUserInfo(ctx sdk.Context, info types.UserInfo) {

}

func (k Keeper) GetUserInfo(ctx sdk.Context, myAddr []byte) *types.UserInfo {
	return &types.UserInfo{}
}
