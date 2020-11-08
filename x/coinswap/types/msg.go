package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/big"
)

// All mapping message define。
var (
	// type check
	_ sdk.Msg = MsgCoinSwap{}
	_ sdk.Tx  = MsgCoinSwap{}
)

const (
	TypeMsgCoinSwap = "coinswap"
)

// All coinswap message define.

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
func (msg MsgCoinSwap) Type() string { return TypeMsgCoinSwap }

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
