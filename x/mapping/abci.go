package mapping

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oinfinance/crossnode/bridge"
	"github.com/oinfinance/crossnode/x/mapping/keeper"
)

var (
	RefreshPoint = int64(150)
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	blockNumber := ctx.BlockHeight()
	if blockNumber <= 0 && (blockNumber%RefreshPoint != 0) {
		return
	}

	maplist := k.GetAllMapInfo(ctx)
	for _, info := range maplist {
		gate := bridge.NewBridge(info.ChainId)
		newBalance := gate.GetBalance(info.RemoteAddr)
		if newBalance.Cmp(info.Balance) != 0 {
			info.Balance = newBalance
			k.UpdateMapping(ctx, info)
		}
	}
}
