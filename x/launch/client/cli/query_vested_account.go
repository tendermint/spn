package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/launch/types"
)

func CmdListVestedAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-vested-account",
		Short: "list all vestedAccount",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllVestedAccountRequest{
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
		Use:   "show-vested-account [chainID] [address]",
		Short: "shows a vestedAccount",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argsChainID, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}
			argsAddress, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}

			params := &types.QueryGetVestedAccountRequest{
				ChainID: argsChainID,
				Address: argsAddress,
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
