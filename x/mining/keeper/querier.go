package keeper

import (
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oinfinance/crossnode/x/mining/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// creates a querier for auth REST endpoints
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {
		case types.QueryRewards:
			return queryRewards(ctx, req, k)
		default:
			return nil, sdk.ErrUnknownRequest("unknown mapping query endpoint")
		}
	}
}

func queryRewards(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, sdk.Error) {
	var params types.QueryRewardsParams
	if err := k.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrInternal("unmarshal param failed")
	}
	addr, err := hex.DecodeString(params.LocalAddr)
	if err != nil {
		return nil, sdk.ErrInternal("invalid addr param")
	}

	response := BuildResponse(k.GetRewards(ctx, addr))
	bz, err := codec.MarshalJSONIndent(k.cdc, response)
	if err != nil {
		return nil, sdk.ErrInternal("marshal response info failed")
	}
	return bz, nil
}

func BuildResponse(info *types.RewardsInfo) types.QueryRewardsResponse {
	t := types.QueryRewardsResponse{}
	if info == nil {
		t.Rewards = 0
	} else {
		t.Rewards = info.Rewards.Uint64()
	}
	return t
}
