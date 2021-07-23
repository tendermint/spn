package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/profile/types"
)

func CmdShowValidatorByConsAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-validator-by-cons-address [consAddress]",
		Short: "shows a validatorByConsAddress",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argsConsAddress, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetValidatorByConsAddressRequest{
				ConsAddress: argsConsAddress,
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
