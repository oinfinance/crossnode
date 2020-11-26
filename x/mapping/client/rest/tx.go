package rest

import (
	"encoding/hex"
	"github.com/oinfinance/crossnode/x/mapping/types"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

type RegisterReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	ErcAddr string       `json:"erc_addr"` // erc20 address
	CCAddr  string       `json:"cc_addr"`  // cross chain address
}

// SendRequestHandlerFn - http request handler to send coins to a address.
func RegisterRequestHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		ercAddr, err := hex.DecodeString(req.ErcAddr)
		if err != nil {
			return
		}
		msg := types.NewMsgRegister(ercAddr, []byte(req.CCAddr))
		if e := msg.ValidateBasic(); e != nil {
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

type VerifyReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	ErcAddr string       `json:"erc_addr"` // erc20 address
	CCAddr  string       `json:"cc_addr"`  // cross chain address
}

func VerifyRequestHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req VerifyReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		ercAddr, err := hex.DecodeString(req.ErcAddr)
		if err != nil {
			return
		}
		msg := types.NewMsgMapVerify(ercAddr, []byte(req.CCAddr))
		if e := msg.ValidateBasic(); e != nil {
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}
