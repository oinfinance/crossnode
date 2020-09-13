package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

// RegisterRoutes registers the auth module REST routes.
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	//r.HandleFunc(
	//	"/mapping/accounts/{address}", QueryAccountRequestHandlerFn(storeName, cliCtx),
	//).Methods("GET")
}
