package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/oinfinance/crossnode/x/coinswap/types"
	"github.com/spf13/cobra"
)

// GetTxCmd returns the transaction commands for this module
func GetQueryCmd(route string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the auth module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	//
	//cmd.AddCommand(client.GetCommands(
	//	GetCmdQueryMapinfo(route, cdc))...)

	return cmd
}
