package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// All mapping message define。
var (
	// type check
	_ sdk.Msg = MsgWithDraw{}
	_ sdk.Tx  = MsgWithDraw{}
)

const (
	TypeMsgRegister = "withDraw"
)

type MsgWithDraw struct {
	RemoteAccount []byte `json:"account"`   // source chain account
	MyAddress     []byte `json:"myAddress"` // binding address with local address
	Amount 		  uint64 `json:"amount"`	// the number
}

func NewMsgWithDraw(sourceAccount []byte, myAddress []byte, amount uint64) *MsgWithDraw {
	return &MsgWithDraw{
		RemoteAccount: sourceAccount,
		MyAddress:     myAddress,
		Amount: amount,
	}
}

// Route should return the name of the module
func (msg MsgWithDraw) Route() string { return RouterKey }

// Type should return the action
func (msg MsgWithDraw) Type() string { return TypeMsgRegister }

// GetMsgs returns a single MsgSetAccName as an sdk.Msg.
func (msg MsgWithDraw) GetMsgs() []sdk.Msg {
	return []sdk.Msg{msg}
}

// ValidateBasic runs stateless checks on the message
func (msg MsgWithDraw) ValidateBasic() sdk.Error {
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgWithDraw) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgWithDraw) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.MyAddress}
}
