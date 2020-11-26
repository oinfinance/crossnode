package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// All mapping message define。
var (
	// type check
	_ sdk.Msg = MsgRegister{}
	_ sdk.Tx  = MsgRegister{}

	_ sdk.Msg = MsgMapVerify{}
	_ sdk.Tx  = MsgMapVerify{}
)

const (
	TypeMsgRegister  = "register"
	TypeMsgMapVerify = "verify"
)

type MsgRegister struct {
	Sender  string `json:"sender"`
	ErcAddr string `json:"erc_addr"`   // 用户的ERC20地址
	CCAddr  string `json:"cross_addr"` // cross chain address to binding
}

func NewMsgRegister(sender string, ercAddr string, ccAddr string) *MsgRegister {
	return &MsgRegister{
		Sender:  sender,
		ErcAddr: ercAddr,
		CCAddr:  ccAddr,
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
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRegister) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRegister) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

type MsgMapVerify struct {
	ErcAddr string `json:"erc_addr"`   // 用户的ERC20地址
	CCAddr  string `json:"cross_addr"` // cross chain address to binding
}

func NewMsgMapVerify(ercAddr string, ccAddr string) *MsgMapVerify {
	return &MsgMapVerify{
		ErcAddr: ercAddr,
		CCAddr:  ccAddr,
	}
}

// Route should return the name of the module
func (msg MsgMapVerify) Route() string { return RouterKey }

// Type should return the action
func (msg MsgMapVerify) Type() string { return TypeMsgMapVerify }

// GetMsgs returns a single MsgSetAccName as an sdk.Msg.
func (msg MsgMapVerify) GetMsgs() []sdk.Msg {
	return []sdk.Msg{msg}
}

// ValidateBasic runs stateless checks on the message
func (msg MsgMapVerify) ValidateBasic() sdk.Error {
	if msg.CCAddr == "" || msg.ErcAddr == "" {
		return sdk.NewError(DefaultCodespace, CodeInvalidAddress, "invalid ccaddr or ercaddr")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgMapVerify) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgMapVerify) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.CCAddr)
	return []sdk.AccAddress{addr}
}
