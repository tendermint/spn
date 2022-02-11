package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/monitoringc/types"
)

func CmdListLaunchIDFromVerifiedClientID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-launch-id-from-verified-client-id",
		Short: "list all LaunchIDFromVerifiedClientID",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllLaunchIDFromVerifiedClientIDRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.LaunchIDFromVerifiedClientIDAll(context.Background(), params)
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

func CmdShowLaunchIDFromVerifiedClientID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-launch-id-from-verified-client-id [client-id]",
		Short: "shows a LaunchIDFromVerifiedClientID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argClientID := args[0]

			params := &types.QueryGetLaunchIDFromVerifiedClientIDRequest{
				ClientID: argClientID,
			}

			res, err := queryClient.LaunchIDFromVerifiedClientID(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
