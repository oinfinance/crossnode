package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/oinfinance/crossnode/x/coinswap/types"
	"github.com/spf13/cobra"
)

// GetTxCmd returns the transaction commands for this module
func GetQueryCmd(route string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the coinswap module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(client.GetCommands(
		GetCmdQueryCoinSwapReceipt(route, cdc))...)

	return cmd
}
func GetCmdQueryCoinSwapReceipt(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "receipt [hash]",
		Short: "Get coinswap mint/burn transaction's receipt",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clictx := context.NewCLIContext().WithCodec(cdc)
			hash := args[0]

			bz, err := cdc.MarshalJSON(types.QueryReceiptByHashParams{Hash: hash})
			if err != nil {
				return err
			}
			route := fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryReceiptByHash)
			res, _, err := clictx.QueryWithData(route, bz)
			if err != nil {
				return fmt.Errorf("could not resolve: %s", err)
			}
			var resp types.CoinSwapReceipt
			if err = cdc.UnmarshalJSON(res, &resp); err != nil {
				return err
			}
			return clictx.PrintOutput(resp)
		},
	}
}
