package cli

import (
	"encoding/hex"
	"errors"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
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
		MappingVerifyTxCmd(cdc),
	)...)
	return txCmd
}

func MappingVerifyTxCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "verify [crossnode_address] [erc_address] ",
		Short: "verify crossnode bind with erc20 address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContextWithFrom(args[0]).WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			var msg *types.MsgMapVerify
			if ercAddr, e := hex.DecodeString(args[1]); e != nil {
				return errors.New("invalid erc20 address")
			} else {
				msg = types.NewMsgMapVerify(ercAddr, []byte(args[0]))
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
