package coinswap

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/oinfinance/crossnode/x/coinswap/client/cli"
	"github.com/oinfinance/crossnode/x/coinswap/client/rest"
	"github.com/oinfinance/crossnode/x/coinswap/keeper"
	"github.com/oinfinance/crossnode/x/coinswap/types"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
)

// coinswap module implement coin deposit and burning.
var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// app module basics object
type AppModuleBasic struct{}

// module name
func (ab AppModuleBasic) Name() string {
	return types.ModuleName
}

// register module codec
func (ab AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	types.RegisterCodec(cdc)
}

// default genesis state
func (ab AppModuleBasic) DefaultGenesis() json.RawMessage {
	return types.ModuleCdc.MustMarshalJSON(DefaultGenesisState())
}

// module validate genesis
func (ab AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	var data GenesisState
	if err := types.ModuleCdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %s", types.ModuleName, err)
	}

	return ValidateGenesis(data)
}

// register rest routes
func (ab AppModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router) {
	rest.RegisterRoutes(ctx, rtr, types.StoreKey)
}

// get the root tx command of this module
func (ab AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetTxCmd(cdc)
}

// get the root query command of this module
func (ab AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetQueryCmd(types.StoreKey, cdc)
}

//___________________________
// app module object
type AppModule struct {
	AppModuleBasic
	SwapKeep keeper.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(keep keeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		SwapKeep:       keep,
	}
}

// module name
func (am AppModule) Name() string {
	return types.ModuleName
}

// register invariants
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// module message route name
func (am AppModule) Route() string { return types.RouterKey }

// module handler
func (am AppModule) NewHandler() sdk.Handler {
	return NewHandler(am.SwapKeep)
}

// module querier route name
func (am AppModule) QuerierRoute() string {
	return types.QuerierRoute
}

// module querier
func (am AppModule) NewQuerierHandler() sdk.Querier {
	return keeper.NewQuerier(am.SwapKeep)
}

// module init-genesis
func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState GenesisState
	types.ModuleCdc.MustUnmarshalJSON(data, &genesisState)
	// Todo: call InitGenesis
	return []abci.ValidatorUpdate{}
}

// module export genesis
func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	gs := ExportGenesis(ctx)
	return types.ModuleCdc.MustMarshalJSON(gs)
}

// module begin-block
func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
	BeginBlocker(ctx, am.SwapKeep)
}

// module end-block
func (AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {

	return []abci.ValidatorUpdate{}
}
