package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/reward/types"
)

func CmdListRewardPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-reward-pool",
		Short: "List all reward pools",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllRewardPoolRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.RewardPoolAll(context.Background(), params)
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

func CmdShowRewardPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-reward-pool [launch-id]",
		Short: "Shows the reward pool for a launch",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argLaunchID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			request := &types.QueryGetRewardPoolRequest{
				LaunchID: argLaunchID,
			}

			res, err := queryClient.RewardPool(context.Background(), request)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
