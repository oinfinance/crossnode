package mining

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oinfinance/crossnode/x/mining/keeper"
)

var (
	RefreshPoint = int64(150)
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	//blockNumber := ctx.BlockHeight()
	//if blockNumber <= 0 && (blockNumber%RefreshPoint != 0) {
	//	return
	//}
	//
	//maplist := k.GetAllMapInfo(ctx)
	//for _, info := range maplist {
	//	gate := bridge.NewBridge(info.ChainId)
	//	newBalance := gate.GetBalance(info.RemoteAddr)
	//	if newBalance.Cmp(info.Balance) != 0 {
	//		info.Balance = newBalance
	//		k.UpdateMapping(ctx, info)
	//	}
	//}
}
