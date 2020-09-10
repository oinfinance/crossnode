package cli

import (
	"errors"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/oinfinance/crossnode/bridge"
	"github.com/oinfinance/crossnode/x/mapping/types"
	"github.com/spf13/cobra"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "mapping transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(client.PostCommands(
		GetCmdRegister(cdc),
	)...)
	return txCmd
}

func GetCmdRegister(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "register [token_address] [chainName] [tokenName] to [crossnode_address]",
		Short:"register remote chain token to corssnode address",
		Args: cobra.ExactArgs(4),
		RunE: func (cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			oinAddress, err := sdk.AccAddressFromHex(args[0])
			if err != nil {
				return err
			}
			chainName := args[1]
			tokenName := args[2]
			if !bridge.SupportedByName(chainName, tokenName) {
				return errors.New("unsupported chain token pair")
			}

			crossAddress, err := sdk.AccAddressFromHex(args[3])
			if err != nil {
				return err
			}

			msg := types.NewMsgRegister(oinAddress.Bytes(), uint(bridge.ChainIdByName(chainName)), uint(bridge.TokenIdByName(tokenName)),
				crossAddress.Bytes())

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
