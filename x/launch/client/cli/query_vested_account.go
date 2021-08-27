package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/launch/types"
)

func CmdListVestedAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-vested-account [chain-id]",
		Short: "list all vestedAccount",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			chainID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryAllVestedAccountRequest{
				ChainID:    chainID,
				Pagination: pageReq,
			}

			res, err := queryClient.VestedAccountAll(context.Background(), params)
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

func CmdShowVestedAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-vested-account [chain-id] [address]",
		Short: "shows a vestedAccount",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			chainID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetVestedAccountRequest{
				ChainID: chainID,
				Address: args[1],
			}

			res, err := queryClient.VestedAccount(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
