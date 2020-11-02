package cli

import (
	"fmt"
	"strconv"
	"strings"

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

	cmd.AddCommand(
		CmdCreateChannel(),
		CmdSendMessage(),
		CmdVotePoll(),
	)

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

// CmdSendMessage returns the transaction command to send a new message
func CmdSendMessage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-message [channel-id] [message-content]",
		Short: "Send a new message to a channel",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			channelID, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			// Flags
			payload, _ := cmd.Flags().GetString(FlagPayload)
			tagsString, _ := cmd.Flags().GetString(FlagTags)
			pollOptionsString, _ := cmd.Flags().GetString(FlagPollOptions)

			// Get the tags
			var tags []string
			if tagsString != "" {
				tags = strings.Split(tagsString, ",")
			}

			// Get the poll options
			var pollOptions []string
			if pollOptionsString != "" {
				pollOptions = strings.Split(pollOptionsString, ",")
			}

			// Create and send message
			msg, err := types.NewMsgSendMessage(
				int32(channelID),
				clientCtx.GetFromAddress(),
				args[1],
				tags,
				pollOptions,
				[]byte(payload),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FlagSetPollOptions())
	cmd.Flags().AddFlagSet(FlagSetTags())
	cmd.Flags().AddFlagSet(FlagSetPayload())
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdVotePoll returns the transaction command to vote on a poll
func CmdVotePoll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vote-poll [value] [channel-id] [message-index]",
		Short: "Vote for a poll",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			// Convert value
			voteValue, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			// Convert channel ID
			channelID, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			// Convert message Index
			messageIndex, err := strconv.Atoi(args[3])
			if err != nil {
				return err
			}

			// Flags
			payload, _ := cmd.Flags().GetString(FlagPayload)

			// Vote
			msg, err := types.NewMsgVotePoll(
				int32(channelID),
				int32(messageIndex),
				clientCtx.GetFromAddress(),
				int32(voteValue),
				[]byte(payload),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FlagSetPollOptions())
	cmd.Flags().AddFlagSet(FlagSetTags())
	cmd.Flags().AddFlagSet(FlagSetPayload())
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
