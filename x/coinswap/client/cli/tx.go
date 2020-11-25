package cli

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/oinfinance/crossnode/bridge"
	"github.com/oinfinance/crossnode/x/coinswap/types"
	"github.com/spf13/cobra"
	"math/big"
	"strconv"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "coinswap transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(client.PostCommands(
		CoinMintCmd(cdc),
		CoinBurnCmd(cdc),
	)...)
	return txCmd
}

func CoinMintCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "mint [from_name_address] [from_chain_name] [txhash] [depoly_addr] [token_name] [token_num] [receiver_addr] [to_chain_name]",
		Short: "mint token on target chain to the receiver",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithFrom(args[0]).WithCodec(cdc)
			//cliCtx := context.NewCLIContext().WithCodec(cdc)

			fromChain := args[1]
			txhash := args[2]
			fromAddr := args[3]
			token := args[4]
			value, e2 := strconv.Atoi(args[5])
			if e2 != nil {
				return errors.New("invalid token_num")
			}
			toAddr := args[6]
			toChain := args[7]

			if !bridge.SupportedGroup(fromChain, toChain, token) {
				return errors.New(fmt.Sprintf("unsupported chain token pair:"+
					"from: <%s> to: <%s> token: <%s>", fromChain, toChain, token))
			}
			fromChainId := bridge.ChainIdByName(fromChain)
			toChainId := bridge.ChainIdByName(toChain)
			tokenId := bridge.TokenIdByName(token)

			msgCoinMint := types.NewMsgCoinSwap([]byte(txhash), int(fromChainId), []byte(fromAddr), int(tokenId),
				big.NewInt(int64(value)), []byte(toAddr), int(toChainId), 1)

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msgCoinMint})
		},
	}
}

func CoinBurnCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "burn [from_name_address] [from_chain_name] [burn_addr] [token_type] [token_num] [receiver_addr] [to_chain_name]",
		Short: "burn token on target chain to the receiver",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithFrom(args[0]).WithCodec(cdc)
			//cliCtx := context.NewCLIContext().WithCodec(cdc)

			fromChain := args[1]
			txhash := args[2]
			fromAddr := args[3]
			token := args[4]
			value, e2 := strconv.Atoi(args[5])
			if e2 != nil {
				return errors.New("invalid token_num")
			}
			toAddr := args[6]
			toChain := args[7]

			if !bridge.SupportedGroup(fromChain, toChain, token) {
				return errors.New(fmt.Sprintf("unsupported chain token pair:"+
					"from: <%s> to: <%s> token: <%s>", fromChain, toChain, token))
			}
			fromChainId := bridge.ChainIdByName(fromChain)
			toChainId := bridge.ChainIdByName(toChain)
			tokenId := bridge.TokenIdByName(token)

			msgCoinMint := types.NewMsgCoinSwap([]byte(txhash), int(fromChainId), []byte(fromAddr), int(tokenId),
				big.NewInt(int64(value)), []byte(toAddr), int(toChainId), 0)

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msgCoinMint})
		},
	}
}
