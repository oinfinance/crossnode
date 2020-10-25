package coinswap

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oinfinance/crossnode/x/coinswap/keeper"
	"github.com/oinfinance/crossnode/x/coinswap/types"
)

// NewHandler returns a handler for "greeter" type messages.
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgMint:
			return handleMsgMint(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized mapping message type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgMint(ctx sdk.Context, k keeper.Keeper, msg *types.MsgMint) sdk.Result {
	if err := msg.ValidateBasic(); err != nil {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidInput, "validate basic failed").Result()
	}

	return sdk.Result{}
}

