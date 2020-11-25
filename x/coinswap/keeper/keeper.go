package keeper

import (
	"errors"
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

func (k Keeper) HasRecord(ctx sdk.Context, rhash []byte) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(rhash)
}

func (k Keeper) AddRecord(ctx sdk.Context, rhash []byte, r *types.CoinSwapRecord) error {
	store := ctx.KVStore(k.storeKey)
	if store.Has(rhash) {
		return errors.New("the token on the chain has been maped")
	} else {
		var rs types.CoinSwapRecordStorage
		rs.Record.Value = r.Value
		rs.Record.ToAddr = r.ToAddr
		rs.Record.ToChain = r.ToChain
		rs.Record.AddedBlock = r.AddedBlock
		rs.Record.Token = r.Token
		rs.Record.FromAddr = r.FromAddr
		rs.Record.FromChain = r.FromChain
		rs.Record.TxHash = r.TxHash
		rs.Receipt.Status = types.RecordStatusWaited // wait
		rs.Receipt.Receipt = ""

		data, err := k.cdc.MarshalBinaryBare(rs)
		if err != nil {
			return err
		}
		// save new map info to store.
		store.Set(rhash, data)
	}
	return nil
}

func (k Keeper) GetRecord(ctx sdk.Context, rhash []byte) *types.CoinSwapRecordStorage {
	var rs types.CoinSwapRecordStorage
	store := ctx.KVStore(k.storeKey)
	if data := store.Get(rhash); data == nil {
		return nil
	} else {
		k.cdc.UnmarshalBinaryBare(data, &rs)
	}

	return &rs
}

func (k Keeper) UpdateRecord(ctx sdk.Context, rhash []byte, rs *types.CoinSwapRecordStorage) error {
	store := ctx.KVStore(k.storeKey)
	if store.Has(rhash) {
		data, err := k.cdc.MarshalBinaryBare(rs)
		if err != nil {
			return err
		}
		// save new map info to store.
		store.Set(rhash, data)
	} else {
		return errors.New("not found map record")
	}
	return nil
}

func (k Keeper) GetAllRecord(ctx sdk.Context) []*types.CoinSwapRecordStorage {
	var list = make([]*types.CoinSwapRecordStorage, 0)

	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, []byte{})

	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var rs types.CoinSwapRecordStorage
		k.cdc.MustUnmarshalBinaryBare(iter.Value(), &rs)
		list = append(list, &rs)
	}

	return list
}
