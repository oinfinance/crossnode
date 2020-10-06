package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/oinfinance/crossnode/x/mining/types"
	"github.com/spf13/cobra"
	"strconv"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "withdraw transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(client.PostCommands(
		GetCmdWithDraw(cdc),
	)...)
	return txCmd
}

func GetCmdWithDraw(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "mining [oin_address] [crossnode_address] [amount]",
		Short:"withdraw rewards to oin-address",
		Args: cobra.ExactArgs(3),
		RunE: func (cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			oinAddress, err := sdk.AccAddressFromHex(args[0])
			if err != nil {
				return err
			}

			crossAddress, err := sdk.AccAddressFromHex(args[1])

			amount, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgWithDraw(oinAddress.Bytes(), crossAddress.Bytes(), amount)

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
