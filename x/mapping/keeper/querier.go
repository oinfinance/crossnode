package keeper

import (
	"bytes"
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oinfinance/crossnode/x/mapping/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// creates a querier for auth REST endpoints
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {
		case types.QueryInfoByOin:
			return queryInfoByOin(ctx, req, k)
		case types.QueryInfoByLocal:
			return queryInfoByLocal(ctx, req, k)
		default:
			return nil, sdk.ErrUnknownRequest("unknown mapping query endpoint")
		}
	}
}

func queryInfoByOin(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, sdk.Error) {
	var params types.QueryInfoByOinParams
	if err := k.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrInternal("unmarshal failed")
	}

	remoteAddr, err := hex.DecodeString(params.OinAddr)
	if err != nil {
		return nil, sdk.ErrInternal("invalid addr param")
	}

	all := k.GetAllMapInfo(ctx)
	var pinfo *types.MappingInfo
	for _, info := range all {
		if bytes.Compare(info.RemoteAddr, remoteAddr) == 0 {
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

func queryInfoByLocal(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, sdk.Error) {
	var params types.QueryInfoByLocalParams
	if err := k.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrInternal("unmarshal failed")
	}
	addr, err := hex.DecodeString(params.LocalAddr)
	if err != nil {
		return nil, sdk.ErrInternal("invalid addr param")
	}

	response := BuildResponse(k.GetMapInfo(ctx, addr))
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
		t.RemoteAddr = info.RemoteAddr
		t.ChainId = info.ChainId
		t.TokenType = info.TokenType
		t.MyAddress = info.MyAddress
		t.Balance = info.Balance.String()
	}
	return t
}
