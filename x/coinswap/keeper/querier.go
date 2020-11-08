package keeper

import (
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oinfinance/crossnode/x/coinswap/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// creates a querier for auth REST endpoints
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {
		case types.QueryReceiptByHash:
			return queryReceiptByHash(ctx, req, k)
		default:
			return nil, sdk.ErrUnknownRequest("unknown coin query endpoint")
		}
	}
}

func queryReceiptByHash(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, sdk.Error) {
	var params types.QueryReceiptByHashParams
	if err := k.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrInternal("unmarshal failed")
	}

	rhash, err := hex.DecodeString(params.Hash)
	if err != nil {
		return nil, sdk.ErrInternal("invalid hash param")
	}
	crs := k.GetRecord(ctx, rhash)
	if crs == nil {
		return nil, sdk.ErrInternal("have no record with given hash " + params.Hash)
	}

	bz, err := codec.MarshalJSONIndent(k.cdc, crs.CoinSwapReceipt)
	if err != nil {
		return nil, sdk.ErrInternal("marshal mapping info failed")
	}
	return bz, nil
}
