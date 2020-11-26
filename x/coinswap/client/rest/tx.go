package rest

import (
	"github.com/oinfinance/crossnode/bridge"
	"github.com/oinfinance/crossnode/x/coinswap/types"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

type MintCoinReq struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	TxHash    string       `json:"txHash"`    // 抵押交易hash
	FromChain string       `json:"fromChain"` // 原始链
	FromAddr  string       `json:"fromAddr"`  // 抵押者地址
	Token     string       `json:"token"`     // 代币类型
	Value     int          `json:"value"`     // 抵押数量
	ToAddr    string       `json:"toAddr"`    // 接收铸币的地址
	ToChain   string       `json:"toChain"`   // 目标链
}

// SendRequestHandlerFn - http request handler to send coins to a address.
func MintCoinRequestHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req MintCoinReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		from := req.BaseReq.From

		if !bridge.SupportedGroup(req.FromChain, req.ToChain, req.Token) {
			return
		}
		fromChainId := bridge.ChainIdByName(req.FromChain)
		toChainId := bridge.ChainIdByName(req.ToChain)
		tokenId := bridge.TokenIdByName(req.Token)

		msgCoinMint := types.NewMsgCoinSwap(from, req.TxHash, int(fromChainId), req.FromAddr, int(tokenId),
			uint64(req.Value), req.ToAddr, int(toChainId), 1)

		if e := msgCoinMint.ValidateBasic(); e != nil {
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msgCoinMint})
	}
}

type BurnCoinReq struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	TxHash    string       `json:"txHash"`    // 抵押交易hash
	FromChain string       `json:"fromChain"` // 原始链
	FromAddr  string       `json:"fromAddr"`  // 抵押者地址
	Token     string       `json:"token"`     // 代币类型
	Value     int          `json:"value"`     // 抵押数量
	ToAddr    string       `json:"toAddr"`    // 接收铸币的地址
	ToChain   string       `json:"toChain"`   // 目标链
}

func BurnCoinRequestHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req MintCoinReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		from := req.BaseReq.From

		if !bridge.SupportedGroup(req.FromChain, req.ToChain, req.Token) {
			return
		}
		fromChainId := bridge.ChainIdByName(req.FromChain)
		toChainId := bridge.ChainIdByName(req.ToChain)
		tokenId := bridge.TokenIdByName(req.Token)

		msgCoinBurn := types.NewMsgCoinSwap(from, req.TxHash, int(fromChainId), req.FromAddr, int(tokenId),
			uint64(req.Value), req.ToAddr, int(toChainId), 0)

		if e := msgCoinBurn.ValidateBasic(); e != nil {
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msgCoinBurn})
	}
}
