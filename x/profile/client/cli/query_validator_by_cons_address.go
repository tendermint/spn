package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/profile/types"
)

func CmdShowValidatorByConsAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-validator-by-cons-address [consensus-address]",
		Short: "shows a ValidatorByConsAddress",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argConsensusAddress := args[0]

			params := &types.QueryGetValidatorByConsAddressRequest{
				ConsensusAddress: argConsensusAddress,
			}

			res, err := queryClient.ValidatorByConsAddress(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
