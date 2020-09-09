package mapping

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oinfinance/crossnode/x/mapping/keeper"
	"github.com/oinfinance/crossnode/x/mapping/types"
	"math/big"
)

// NewHandler returns a handler for "greeter" type messages.
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgRegister:
			return handleMsgRegister(ctx, k, msg)
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
	info.RemoteAddr = msg.RemoteAccount
	info.ChainId = msg.ChainId
	info.TokenType = msg.TokenType
	info.MyAddress = msg.MyAddress
	info.Balance = big.NewInt(0)
	err := k.AddMapping(ctx, &info)
	if err != nil {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidInput, "has been mapped").Result()
	}
	return sdk.Result{}
}
