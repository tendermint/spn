package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/participation/types"
)

var _ = strconv.Itoa(0)

func CmdShowAvailableAllocations() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-available-allocations [address]",
		Short: "Show the available, unused allocations for an account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetAvailableAllocationsRequest{
				Address: reqAddress,
			}

			res, err := queryClient.AvailableAllocations(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
