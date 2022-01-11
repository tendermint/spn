package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/tendermint/spn/x/launch/types"
)

var _ = strconv.Itoa(0)

func CmdRequestSelection() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-selection [launch-id] [request-ids]",
		Short: "Query request-selection",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqRequestIDs := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			reqLaunchID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryRequestSelectionRequest{
				LaunchID:   reqLaunchID,
				RequestIDs: reqRequestIDs,
			}

			res, err := queryClient.RequestSelection(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
