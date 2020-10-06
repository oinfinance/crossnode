package mining

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oinfinance/crossnode/x/mapping/keeper"
	"github.com/oinfinance/crossnode/x/mapping/types"
)

// GenesisState is the mapping state that must be provided at genesis.
type GenesisState struct {
	MapList []types.MappingInfo `json:"maplist" yaml:"maplist"`
}

// NewGenesisState creates a new genesis state.
func NewGenesisState(maplist []types.MappingInfo) GenesisState {
	return GenesisState{MapList: maplist}
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() GenesisState {
	return NewGenesisState(nil)
}

// InitGenesis sets distribution information for genesis.
func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data GenesisState) {
	for _, info := range data.MapList {
		keeper.AddMapping(ctx, &info)
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) GenesisState {
	var maplist []types.MappingInfo
	l := keeper.GetAllMapInfo(ctx)
	for _, info := range l {
		maplist = append(maplist, *info)
	}
	return NewGenesisState(maplist)
}

// ValidateGenesis performs basic validation of mapping genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data GenesisState) error { return nil }
