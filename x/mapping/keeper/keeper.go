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
	mapStoreKey    sdk.StoreKey // Unexposed key to access store from sdk.Context
	verifyStoreKey sdk.StoreKey // Unexposed key to access store from sdk.Context
	cdc            *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper returns a guardian keeper
func NewKeeper(cdc *codec.Codec, mapStoreKey sdk.StoreKey, verifyStoreKey sdk.StoreKey) Keeper {
	keeper := Keeper{
		mapStoreKey:    mapStoreKey,
		verifyStoreKey: verifyStoreKey,
		cdc:            cdc,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("%s", types.ModuleName))
}

func (k Keeper) AddMapping(ctx sdk.Context, info *types.MappingInfo) error {
	store := ctx.KVStore(k.mapStoreKey)
	key := []byte(info.ErcAddr)
	if store.Has(key) {
		return errors.New("the record on the chain has been maped")
	} else {
		data, err := k.cdc.MarshalBinaryBare(info)
		if err != nil {
			return err
		}
		// save new map info to store.
		store.Set(key, data)
	}
	return nil
}

func (k Keeper) UpdateMapping(ctx sdk.Context, info *types.MappingInfo) error {
	store := ctx.KVStore(k.mapStoreKey)
	key := []byte(info.ErcAddr)
	if store.Has(key) {
		data, err := k.cdc.MarshalBinaryBare(info)
		if err != nil {
			return err
		}
		// save new map info to store.
		store.Set(key, data)
	} else {
		return errors.New("not found map record")
	}
	return nil
}

func (k Keeper) GetMapInfo(ctx sdk.Context, ercAddr []byte) *types.MappingInfo {
	var info types.MappingInfo
	store := ctx.KVStore(k.mapStoreKey)

	if data := store.Get(ercAddr); data == nil {
		return nil
	} else {
		k.cdc.UnmarshalBinaryBare(data, &info)
	}

	return &info
}

func (k Keeper) GetAllMapInfo(ctx sdk.Context) []*types.MappingInfo {
	var list = make([]*types.MappingInfo, 0)

	store := ctx.KVStore(k.mapStoreKey)
	iter := sdk.KVStorePrefixIterator(store, []byte{})

	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var info types.MappingInfo
		k.cdc.MustUnmarshalBinaryBare(iter.Value(), &info)
		list = append(list, &info)
	}

	return list
}

func (k Keeper) AddVerified(ctx sdk.Context, ccAddr []byte) error {
	store := ctx.KVStore(k.verifyStoreKey)
	key := ccAddr
	status := byte(types.MappingWaitVerify)
	if store.Has(key) {
		return errors.New("the ccAddr on the store has been maped")
	} else {
		// save new map info to store.
		store.Set(key, []byte{status})
	}
	return nil
}

func (k Keeper) UpdateVerified(ctx sdk.Context, ccAddr []byte, status byte) error {
	store := ctx.KVStore(k.verifyStoreKey)
	key := ccAddr
	if store.Has(key) {
		// save new map info to store.
		store.Set(key, []byte{status})
	} else {
		return errors.New("not found map record")
	}
	return nil
}

func (k Keeper) GetVerified(ctx sdk.Context, ccAddr []byte) int {
	var status int
	store := ctx.KVStore(k.verifyStoreKey)

	if data := store.Get(ccAddr); data == nil {
		return 0
	} else {
		status = int(data[0])
	}

	return status
}
