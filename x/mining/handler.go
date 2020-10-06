package mining

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oinfinance/crossnode/x/mining/keeper"
	"github.com/oinfinance/crossnode/x/mining/types"
	"math/big"
)

// NewHandler returns a handler for "greeter" type messages.
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgWithDraw:
			return handleMsgWithDraw(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized mapping message type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgWithDraw(ctx sdk.Context, k keeper.Keeper, msg *types.MsgWithDraw) sdk.Result {
	if err := msg.ValidateBasic(); err != nil {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidInput, "validate basic failed").Result()
	}
	info := k.GetRewards(ctx, msg.MyAddress)
	amount := big.NewInt(0).SetUint64(msg.Amount)
	if info.Rewards.Uint64() < msg.Amount {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidInput, "amount is invalid").Result()
	}

	info.Rewards.Sub(&info.Rewards, amount)
	k.UpdateRewards(ctx, msg.MyAddress, info)

	// todo: send erc20 on ethereum

	return sdk.Result{}
}
