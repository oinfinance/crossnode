package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

// RegisterRoutes registers the auth module REST routes.
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc("/coinswap/mint/", MintCoinRequestHandlerFn(cliCtx)).Methods("POST")
	r.HandleFunc("/coinswap/burn/", BurnCoinRequestHandlerFn(cliCtx)).Methods("POST")
	r.HandleFunc("/coinswap/receipt/{hash}", QueryCoinSwapReceiptRequestHandlerFn(storeName, cliCtx)).Methods("GET")
}
