package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/account/types"
)

func CmdListCoordinatorByAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-coordinator-by-address",
		Short: "list all coordinatorByAddress",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllCoordinatorByAddressRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.CoordinatorByAddressAll(context.Background(), params)
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

func CmdShowCoordinatorByAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-coordinator-by-address [address]",
		Short: "shows a coordinatorByAddress",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argsAddress, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetCoordinatorByAddressRequest{
				Address: argsAddress,
			}

			res, err := queryClient.CoordinatorByAddress(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
