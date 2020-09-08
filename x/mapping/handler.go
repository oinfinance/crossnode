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
		return &sdk.Result{Log: err.Error()}, err
	}

	// parse the chainID from a string to a base-10 integer
	intChainID, ok := new(big.Int).SetString(ctx.ChainID(), 10)
	if !ok {
		return nil, greeter.ErrInvalidChainID(fmt.Sprintf("invalid chainID: %s", ctx.ChainID()))
	}

	st := evm.StateTransition{
		Sender:       ethcmn.BytesToAddress(msg.From.Bytes()),
		AccountNonce: msg.AccountNonce,
		Price:        msg.Price.BigInt(),
		GasLimit:     msg.GasLimit,
		Amount:       msg.Amount.BigInt(),
		Payload:      msg.Payload,
		Csdb:         k.CommitStateDB.WithContext(ctx),
		ChainID:      intChainID,
		Simulate:     ctx.IsCheckTx(),
	}

	if msg.Recipient != nil {
		to := ethcmn.BytesToAddress(msg.Recipient.Bytes())
		st.Recipient = &to
	}

	// Prepare db for logs
	k.CommitStateDB.Prepare(ethcmn.Hash{}, ethcmn.Hash{}, k.TxCount.Get()) // Cannot provide tx hash
	k.TxCount.Increment()

	_, res := st.TransitionCSDB(ctx, false, greeter.DenomDefault)
	return &res, nil
}
