package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/tendermint/spn/x/chat/types"

	"github.com/spf13/cobra"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreateChannel())

	return cmd
}

// CmdCreateChannel returns the transaction command to create a new channel
func CmdCreateChannel() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-channel [name]",
		Short: "Create a new channel",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			// Flags
			payload, _ := cmd.Flags().GetString(FlagPayload)
			description, _ := cmd.Flags().GetString(FlagDescription)

			// Create and send message
			msg, err := types.NewMsgCreateChannel(
				clientCtx.GetFromAddress(),
				args[0],
				description,
				[]byte(payload),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FlagSetDescription())
	cmd.Flags().AddFlagSet(FlagSetPayload())
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
