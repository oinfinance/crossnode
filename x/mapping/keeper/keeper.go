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

func (k Keeper) GetMapInfo(ctx sdk.Context, ercAddr string) *types.MappingInfo {
	var info types.MappingInfo
	key := []byte(ercAddr)
	store := ctx.KVStore(k.mapStoreKey)

	if data := store.Get(key); data == nil {
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

func (k Keeper) AddVerified(ctx sdk.Context, ccAddr string) error {
	store := ctx.KVStore(k.verifyStoreKey)
	addr, _ := sdk.AccAddressFromBech32(ccAddr)
	status := byte(types.MappingWaitVerify)
	if store.Has(addr) {
		return errors.New("the ccAddr on the store has been maped")
	} else {
		// save new map info to store.
		store.Set(addr, []byte{status})
	}
	return nil
}

func (k Keeper) UpdateVerified(ctx sdk.Context, ccAddr string, status byte) error {
	store := ctx.KVStore(k.verifyStoreKey)
	addr, _ := sdk.AccAddressFromBech32(ccAddr)
	if store.Has(addr) {
		// save new map info to store.
		store.Set(addr, []byte{status})
	} else {
		return errors.New("not found map record")
	}
	return nil
}

func (k Keeper) GetVerified(ctx sdk.Context, ccAddr string) int {
	var status int
	store := ctx.KVStore(k.verifyStoreKey)
	addr, _ := sdk.AccAddressFromBech32(ccAddr)
	if data := store.Get(addr); data == nil {
		return 0
	} else {
		status = int(data[0])
	}

	return status
}
