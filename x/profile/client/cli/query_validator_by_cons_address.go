package cli

import (
	"context"
	"encoding/base64"

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

			consAddr, err := base64.StdEncoding.DecodeString(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetValidatorByConsAddressRequest{
				ConsensusAddress: consAddr,
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
