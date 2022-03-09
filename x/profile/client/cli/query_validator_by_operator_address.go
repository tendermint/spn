package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/profile/types"
)

func CmdShowValidatorByOperatorAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-validator-by-operator-address [operator-address]",
		Short: "Shows a validator address by an associated operator address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			opAddr := args[0]

			params := &types.QueryGetValidatorByOperatorAddressRequest{
				OperatorAddress: opAddr,
			}

			res, err := queryClient.ValidatorByOperatorAddress(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
