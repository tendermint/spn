package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/project/types"
)

func CmdShowProjectChains() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-project-chains [project-id]",
		Short: "List all chains of a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argsProjectID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetProjectChainsRequest{
				ProjectID: argsProjectID,
			}

			res, err := queryClient.ProjectChains(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
