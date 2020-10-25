package mining

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oinfinance/crossnode/x/mining/keeper"
)

// GenesisState is the mapping state that must be provided at genesis.
type GenesisState struct {

}

// NewGenesisState creates a new genesis state.
func NewGenesisState() GenesisState {
	return GenesisState{}
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() GenesisState {
	return NewGenesisState()
}

// InitGenesis sets distribution information for genesis.
func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data GenesisState) {
	return
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) GenesisState {
	return NewGenesisState()
}

// ValidateGenesis performs basic validation of mapping genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data GenesisState) error { return nil }
