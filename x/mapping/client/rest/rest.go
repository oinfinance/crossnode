package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

// RegisterRoutes registers the auth module REST routes.
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc("/mapping/register/", RegisterRequestHandlerFn(cliCtx)).Methods("POST")
	r.HandleFunc("/mapping/verify/", VerifyRequestHandlerFn(cliCtx)).Methods("POST")
	r.HandleFunc("/mapping/mapinfo/{address}", QueryMapinfoRequestHandlerFn(storeName, cliCtx)).Methods("GET")
}
