package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/big"
)

// All mapping message define。
var (
	// type check
	_ sdk.Msg = MsgDeposit{}
	_ sdk.Tx  = MsgDeposit{}

	_ sdk.Msg = MsgBurning{}
	_ sdk.Tx  = MsgBurning{}
)

const (
	TypeMsgDeposit = "deposit"
	TypeMsgBurning = "burning"
)

// All coinswap message define.
// deposit msg
type MsgDeposit struct {
	Chain       uint   `json:"chain"`
	Token       uint   `json:"token"`
	Amount      []byte `json:"amount"`
	TargetChain uint   `json:"target"`
}

func NewMsgDeposit(chainid uint, tokenType uint, amount *big.Int, targetChain uint) *MsgDeposit {
	return &MsgDeposit{
		Chain:       chainid,
		Token:       tokenType,
		Amount:      amount.Bytes(),
		TargetChain: targetChain,
	}
}

// Route should return the name of the module
func (msg MsgDeposit) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDeposit) Type() string { return TypeMsgDeposit }

// GetMsgs returns a single MsgSetAccName as an sdk.Msg.
func (msg MsgDeposit) GetMsgs() []sdk.Msg {
	return []sdk.Msg{msg}
}

// ValidateBasic runs stateless checks on the message
func (msg MsgDeposit) ValidateBasic() sdk.Error {
	if msg.Chain < 0 {
		return sdk.NewError(DefaultCodespace, CodeInvalidInput, "invalid chainId")
	}
	if msg.Token < 0 {
		return sdk.NewError(DefaultCodespace, CodeInvalidInput, "invalid tokenType")
	}

	if msg.TargetChain < 0 {
		return sdk.NewError(DefaultCodespace, CodeInvalidInput, "invalid target chainId")
	}
	if big.NewInt(0).SetBytes(msg.Amount).Int64() < 0 {
		return sdk.NewError(DefaultCodespace, CodeInvalidInput, "negative amount")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeposit) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeposit) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

type MsgBurning struct {
	Chain       uint   `json:"chain"`
	Token       uint   `json:"token"`
	Amount      []byte `json:"amount"`
	TargetChain uint   `json:"target"`
}

func NewMsgBurning(chainid uint, tokenType uint, amount *big.Int, targetChain uint) *MsgDeposit {
	return &MsgDeposit{
		Chain:       chainid,
		Token:       tokenType,
		Amount:      amount.Bytes(),
		TargetChain: targetChain,
	}
}

// Route should return the name of the module
func (msg MsgBurning) Route() string { return RouterKey }

// Type should return the action
func (msg MsgBurning) Type() string { return TypeMsgBurning }

// GetMsgs returns a single MsgSetAccName as an sdk.Msg.
func (msg MsgBurning) GetMsgs() []sdk.Msg {
	return []sdk.Msg{msg}
}

// ValidateBasic runs stateless checks on the message
func (msg MsgBurning) ValidateBasic() sdk.Error {
	if msg.Chain < 0 {
		return sdk.NewError(DefaultCodespace, CodeInvalidInput, "invalid chainId")
	}
	if msg.Token < 0 {
		return sdk.NewError(DefaultCodespace, CodeInvalidInput, "invalid tokenType")
	}

	if msg.TargetChain < 0 {
		return sdk.NewError(DefaultCodespace, CodeInvalidInput, "invalid target chainId")
	}
	if big.NewInt(0).SetBytes(msg.Amount).Int64() < 0 {
		return sdk.NewError(DefaultCodespace, CodeInvalidInput, "negative amount")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgBurning) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgBurning) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}
