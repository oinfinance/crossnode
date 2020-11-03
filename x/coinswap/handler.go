package coinswap

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oinfinance/crossnode/x/coinswap/keeper"
	"github.com/oinfinance/crossnode/x/coinswap/types"
	"math/big"
)

// NewHandler returns a handler for "greeter" type messages.
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgCoinSwap:
			return handleMsgCoinSwap(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized mapping message type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgCoinSwap(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCoinSwap) sdk.Result {
	if err := msg.ValidateBasic(); err != nil {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidInput, "validate basic failed").Result()
	}
	record := msg.ToRecord()
	record.AddedBlock = big.NewInt(ctx.BlockHeight())
	hash := record.Hash()

	if k.HasRecord(ctx, hash) {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidInput, "already existed").Result()
	}

	if e := k.AddRecord(ctx, hash, &record); e != nil {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidInput, "add error %s", e.Error()).Result()
	}

	return sdk.Result{Data: hash}
}
