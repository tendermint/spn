package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/chat/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group chat queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdShowChannel(),
		CmdListChannels(),
		CmdListMessages(),
		CmdSearchMessages(),
	)

	return cmd
}

// CmdShowChannel returns the command to show a channel
func CmdShowChannel() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-channel [channel-id]",
		Short: "show info concerning a channel",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			// Convert channel ID
			channelID, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryShowChannelRequest{
				Id: int32(channelID),
			}

			res, err := queryClient.ShowChannel(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdListChannels returns the command to list channels
func CmdListChannels() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-channels",
		Short: "list channels",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryListChannelsRequest{}

			res, err := queryClient.ListChannels(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdListMessages returns the command to list messages in a channel
func CmdListMessages() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-messages [channel-id]",
		Short: "list the messages in a channel",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			// Convert channel ID
			channelID, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryListMessagesRequest{
				ChannelId: int32(channelID),
			}

			res, err := queryClient.ListMessages(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdSearchMessages returns the command to search for tagged messages in a channel
func CmdSearchMessages() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search-messages [tag] [channel-id]",
		Short: "search the messages in a channel that contain a specific tag",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			// Convert channel ID
			channelID, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			params := &types.QuerySearchMessagesRequest{
				ChannelId: int32(channelID),
				Tag:       args[0],
			}

			res, err := queryClient.SearchMessages(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
