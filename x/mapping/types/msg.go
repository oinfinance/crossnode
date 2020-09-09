package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// All mapping message define。
var (
	// type check
	_ sdk.Msg = MsgRegister{}
	_ sdk.Tx  = MsgRegister{}
)

const (
	TypeMsgRegister = "register"
)

type MsgRegister struct {
	RemoteAccount []byte `json:"account"`   // source chain account
	ChainId       uint   `json:"chainId"`   // chain identification
	TokenType     uint   `json:"token"`     // token type identification
	MyAddress     []byte `json:"myAddress"` // binding address with local address
}

func NewMsgRegister(sourceAccount []byte, chainid uint, tokenType uint, myAddress []byte) *MsgRegister {
	return &MsgRegister{
		RemoteAccount: sourceAccount,
		ChainId:       chainid,
		TokenType:     tokenType,
		MyAddress:     myAddress,
	}
}

// Route should return the name of the module
func (msg MsgRegister) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRegister) Type() string { return TypeMsgRegister }

// GetMsgs returns a single MsgSetAccName as an sdk.Msg.
func (msg MsgRegister) GetMsgs() []sdk.Msg {
	return []sdk.Msg{msg}
}

// ValidateBasic runs stateless checks on the message
func (msg MsgRegister) ValidateBasic() sdk.Error {
	if msg.ChainId < 0 {
		return sdk.NewError(DefaultCodespace, CodeInvalidInput, "invalid chainId")
	}
	if msg.TokenType < 0 {
		return sdk.NewError(DefaultCodespace, CodeInvalidInput, "invalid tokenType")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRegister) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRegister) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.MyAddress}
}
