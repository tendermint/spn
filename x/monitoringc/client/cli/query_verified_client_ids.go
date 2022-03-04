package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/monitoringc/types"
)

func CmdShowVerifiedClientIds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verified-client-ids [launch-id]",
		Short: "Retrieves all verified client IDs for a launch",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argLaunchID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetVerifiedClientIdsRequest{
				LaunchID: argLaunchID,
			}

			res, err := queryClient.VerifiedClientIds(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
