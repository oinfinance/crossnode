package keeper

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/oinfinance/crossnode/x/mapping/types"
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

func (k Keeper) AddMapping(ctx sdk.Context, info *types.MappingInfo) error {
	store := ctx.KVStore(k.storeKey)
	if store.Has(info.MyAddress) {
		return errors.New("the token on the chain has been maped")
	} else {
		data, err := k.cdc.MarshalBinaryBare(info)
		if err != nil {
			return err
		}
		// save new map info to store.
		store.Set(info.MyAddress, data)
	}
	return nil
}

func (k Keeper) UpdateMapping(ctx sdk.Context, info *types.MappingInfo) error {
	store := ctx.KVStore(k.storeKey)
	if store.Has(info.MyAddress) {
		data, err := k.cdc.MarshalBinaryBare(info)
		if err != nil {
			return err
		}
		// save new map info to store.
		store.Set(info.MyAddress, data)
	} else {
		return errors.New("not found map record")
	}
	return nil
}

func (k Keeper) GetMapInfo(ctx sdk.Context, myAddr []byte) *types.MappingInfo {
	var info types.MappingInfo
	store := ctx.KVStore(k.storeKey)
	if data := store.Get(myAddr); data == nil {
		return nil
	} else {
		k.cdc.UnmarshalBinaryBare(data, &info)
	}

	return &info
}

func (k Keeper) GetAllMapInfo(ctx sdk.Context) []*types.MappingInfo {
	var list = make([]*types.MappingInfo, 0)

	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, []byte{})

	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var info types.MappingInfo
		k.cdc.MustUnmarshalBinaryBare(iter.Value(), &info)
		list = append(list, &info)
	}

	return list
}
