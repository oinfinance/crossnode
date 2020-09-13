package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// creates a querier for auth REST endpoints
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {
		default:
			return nil, sdk.ErrUnknownRequest("unknown mapping query endpoint")
		}
	}
}
//
//func BuildResponse(info *types.MappingInfo) types.QueryInfoResponse {
//	t := types.QueryInfoResponse{}
//	if info == nil {
//		t.Found = false
//	} else {
//		t.Found = true
//		t.RemoteAddr = info.RemoteAddr
//		t.ChainId = info.ChainId
//		t.TokenType = info.TokenType
//		t.MyAddress = info.MyAddress
//		t.Balance = info.Balance.String()
//	}
//	return t
//}
