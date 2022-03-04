package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/profile/types"
)

func CmdShowCoordinatorByAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-coordinator-by-address [address]",
		Short: "Shows a coordinator by its address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetCoordinatorByAddressRequest{
				Address: args[0],
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
