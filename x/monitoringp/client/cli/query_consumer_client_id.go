package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/monitoringp/types"
)

func CmdShowConsumerClientID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-consumer-client-id",
		Short: "shows the consumer client ID used for the IBC connection",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetConsumerClientIDRequest{}

			res, err := queryClient.ConsumerClientID(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
