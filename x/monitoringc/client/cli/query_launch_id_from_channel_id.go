package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/monitoringc/types"
)

func CmdListLaunchIDFromChannelID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-launch-id-from-channel-id",
		Short: "list all LaunchIDFromChannelID",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllLaunchIDFromChannelIDRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.LaunchIDFromChannelIDAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowLaunchIDFromChannelID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-launch-id-from-channel-id [channel-id]",
		Short: "shows a LaunchIDFromChannelID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argChannelID := args[0]

			params := &types.QueryGetLaunchIDFromChannelIDRequest{
				ChannelID: argChannelID,
			}

			res, err := queryClient.LaunchIDFromChannelID(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
