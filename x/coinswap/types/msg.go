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

	_ sdk.Msg = MsgMint{}
	_ sdk.Tx  = MsgMint{}

	_ sdk.Msg = MsgCoinSwap{}
	_ sdk.Tx  = MsgCoinSwap{}
)

const (
	TypeMsgDeposit = "deposit"
	TypeMsgBurning = "burning"
	TypeMsgMint    = "mint"
)

// All coinswap message define.

// mint msg
type MsgMint struct {
	Chain       uint   `json:"chain"`
	Token       uint   `json:"token"`
	Amount      []byte `json:"amount"`
	TargetChain uint   `json:"target"`
}

func NewMsgMint(chainid uint, tokenType uint, amount *big.Int, targetChain uint) *MsgMint {
	return &MsgMint{
		Chain:       chainid,
		Token:       tokenType,
		Amount:      amount.Bytes(),
		TargetChain: targetChain,
	}
}

// Route should return the name of the module
func (msg MsgMint) Route() string { return RouterKey }

// Type should return the action
func (msg MsgMint) Type() string { return TypeMsgMint }

// GetMsgs returns a single MsgSetAccName as an sdk.Msg.
func (msg MsgMint) GetMsgs() []sdk.Msg {
	return []sdk.Msg{msg}
}

// ValidateBasic runs stateless checks on the message
func (msg MsgMint) ValidateBasic() sdk.Error {
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
func (msg MsgMint) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgMint) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

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

// coinswap msg got from borker node.
type MsgCoinSwap struct {
	TxHash    []byte   `json:"txHash"`    // 抵押交易hash
	FromChain int      `json:"fromChain"` // 原始链
	FromAddr  []byte   `json:"fromAddr"`  // 抵押者地址
	Token     int      `json:"token"`     // 代币类型
	Value     *big.Int `json:"value"`     // 抵押数量
	ToAddr    []byte   `json:"toAddr"`    // 接收铸币的地址
	ToChain   int      `json:"toChain"`   // 目标链
}

func NewMsgCoinSwap(txhash []byte, fromChain int, fromAddr []byte, token int, value *big.Int,
	toAddr []byte, toChain int) *MsgCoinSwap {
	return &MsgCoinSwap{
		TxHash:    txhash,
		FromChain: fromChain,
		FromAddr:  fromAddr,
		Token:     token,
		Value:     value,
		ToAddr:    toAddr,
		ToChain:   toChain,
	}
}

// Route should return the name of the module
func (msg MsgCoinSwap) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCoinSwap) Type() string { return TypeMsgBurning }

// GetMsgs returns a single MsgSetAccName as an sdk.Msg.
func (msg MsgCoinSwap) GetMsgs() []sdk.Msg {
	return []sdk.Msg{msg}
}

// ValidateBasic runs stateless checks on the message
func (msg MsgCoinSwap) ValidateBasic() sdk.Error {

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCoinSwap) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCoinSwap) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg MsgCoinSwap) ToRecord() CoinSwapRecord {
	return CoinSwapRecord{
		TxHash:    msg.TxHash,
		FromChain: msg.FromChain,
		FromAddr:  msg.FromAddr,
		Token:     msg.Token,
		Value:     msg.Value,
		ToAddr:    msg.ToAddr,
		ToChain:   msg.ToChain,
	}
}
