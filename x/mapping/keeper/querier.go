package keeper

import (
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oinfinance/crossnode/x/mapping/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"strings"
)

// creates a querier for auth REST endpoints
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {
		case types.QueryInfoByErc:
			return queryInfoByErc(ctx, req, k)
		case types.QueryInfoByCC:
			return queryInfoByCC(ctx, req, k)
		default:
			return nil, sdk.ErrUnknownRequest("unknown mapping query endpoint")
		}
	}
}

func queryInfoByErc(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, sdk.Error) {
	var params types.QueryInfoByErcParams
	if err := k.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrInternal("unmarshal failed")
	}

	ercAddr, err := hex.DecodeString(params.ErcAddr)
	if err != nil {
		return nil, sdk.ErrInternal("invalid addr param")
	}
	pinfo := k.GetMapInfo(ctx, ercAddr)

	response := BuildResponse(pinfo)
	bz, err := codec.MarshalJSONIndent(k.cdc, response)
	if err != nil {
		return nil, sdk.ErrInternal("marshal mapping info failed")
	}
	return bz, nil
}

func queryInfoByCC(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, sdk.Error) {
	var params types.QueryInfoByCCParams
	if err := k.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrInternal("unmarshal failed")
	}

	all := k.GetAllMapInfo(ctx)
	var pinfo *types.MappingInfo
	for _, info := range all {
		if strings.Compare(info.CCAddr, params.CCAddr) == 0 {
			pinfo = info
			break
		}
	}

	response := BuildResponse(pinfo)
	bz, err := codec.MarshalJSONIndent(k.cdc, response)
	if err != nil {
		return nil, sdk.ErrInternal("marshal mapping info failed")
	}
	return bz, nil
}

func BuildResponse(info *types.MappingInfo) types.QueryInfoResponse {
	t := types.QueryInfoResponse{}
	if info == nil {
		t.Found = false
	} else {
		t.Found = true
		t.ErcAddr = info.ErcAddr
		t.Status = info.Status
		t.CCAddr = info.CCAddr
		t.RegisterBlock = info.RegisterBlock
		t.Balance = info.Balance.Uint64()
	}
	return t
}
