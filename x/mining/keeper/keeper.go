package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oinfinance/crossnode/x/mining/types"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper store user map info, (ERC20(oin) ===> localCoins)
type Keeper struct {
	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context
	cdc      *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper returns a guardian keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	keeper := Keeper{
		storeKey: key,
		cdc:      cdc,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("%s", types.ModuleName))
}

func (k Keeper) GetRewards(ctx sdk.Context, myAddr []byte) *types.RewardsInfo {
	var info types.RewardsInfo
	store := ctx.KVStore(k.storeKey)
	if data := store.Get(myAddr); data == nil {
		return &info
	} else {
		k.cdc.UnmarshalBinaryBare(data, &info)
	}

	return &info
}

func (k Keeper) UpdateRewards(ctx sdk.Context, myAddr []byte, rewards *types.RewardsInfo) error {
	store := ctx.KVStore(k.storeKey)
	data, err := k.cdc.MarshalBinaryBare(rewards)
	if err != nil {
		return err
	}
	// save new map info to store.
	store.Set(myAddr, data)
	return nil
}
