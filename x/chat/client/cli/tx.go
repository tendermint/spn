package cli

import (
	"fmt"

	proto "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/tendermint/spn/x/chat/types"
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
		Use:   "create-channel [name] [subject]",
		Short: "Create a new channel",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			// Get and decode payload if defined
			var payload *proto.Message
			payloadData, _ := cmd.Flags().GetString(FlagPayload)
			if payloadData != "" {
				err = jsonpb.UnmarshalString(payloadData, *payload)
				if err != nil {
					return err
				}
			} else {
				payload = nil
			}

			// Create and send message
			msg, err := types.NewMsgCreateChannel(
				clientCtx.GetFromAddress(),
				args[0],
				args[1],
				payload,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FlagSetPaylaod())
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
