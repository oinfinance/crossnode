package mapping

import (
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oinfinance/crossnode/bridge"
	"github.com/oinfinance/crossnode/x/mapping/keeper"
	"github.com/oinfinance/crossnode/x/mapping/types"
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
		if info.Status == types.MappingVerifyPassed {
			gate := bridge.NewBridge(0)
			addr, _ := hex.DecodeString(info.ErcAddr)
			newBalance := gate.GetBalance(addr)
			if newBalance.Cmp(info.Balance) != 0 {
				info.Balance = newBalance
				k.UpdateMapping(ctx, info)
			}
		}
		// todo: 1.过期未验证的记录定为超时，删除
		// 2. 验证失败的记录，删除
	}
}
