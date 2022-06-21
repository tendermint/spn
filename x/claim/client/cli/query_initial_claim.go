package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/claim/types"
)

func CmdShowInitialClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-initial-claim",
		Short: "shows information about initial claim",
		Long:  "shows if initial claim is enabled and what is the mission ID completed by initial claim",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetInitialClaimRequest{}

			res, err := queryClient.InitialClaim(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
