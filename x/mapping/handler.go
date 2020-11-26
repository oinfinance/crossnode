package mapping

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oinfinance/crossnode/x/mapping/keeper"
	"github.com/oinfinance/crossnode/x/mapping/types"
	"math/big"
	"strings"
)

// NewHandler returns a handler for "greeter" type messages.
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgRegister:
			return handleMsgRegister(ctx, k, msg)
		case *types.MsgMapVerify:
			return handleMsgMapVerify(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized mapping message type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgRegister(ctx sdk.Context, k keeper.Keeper, msg *types.MsgRegister) sdk.Result {
	if err := msg.ValidateBasic(); err != nil {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidInput, "validate basic failed").Result()
	}
	info := types.MappingInfo{}
	info.ErcAddr = msg.ErcAddr
	info.CCAddr = msg.CCAddr
	info.Status = types.MappingWaitVerify
	info.Balance = big.NewInt(0)
	info.RegisterBlock = uint64(ctx.BlockHeader().Height)

	err := k.AddMapping(ctx, &info)
	if err != nil {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidInput, "has been mapped").Result()
	}

	k.AddVerified(ctx, msg.CCAddr)
	return sdk.Result{}
}

func handleMsgMapVerify(ctx sdk.Context, k keeper.Keeper, msg *types.MsgMapVerify) sdk.Result {
	if err := msg.ValidateBasic(); err != nil {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidInput, "validate basic failed").Result()
	}
	info := k.GetMapInfo(ctx, msg.ErcAddr)
	if info == nil {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidAddress, "have no registered").Result()
	}
	if strings.Compare(string(msg.CCAddr), info.CCAddr) != 0 {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidAddress, "binding ccAddr not matched").Result()
	}

	status := k.GetVerified(ctx, msg.CCAddr)
	if status == types.MappingVerifyPassed {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidAddress, "ccAddr have been band").Result()
	}

	info.Status = types.MappingVerifyPassed
	k.UpdateMapping(ctx, info)
	k.UpdateVerified(ctx, msg.CCAddr, byte(types.MappingVerifyPassed))
	return sdk.Result{}
}
