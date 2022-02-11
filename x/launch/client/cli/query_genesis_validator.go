package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/launch/types"
)

func CmdListGenesisValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-genesis-validator [launch-id]",
		Short: "list all genesisValidator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			launchID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryAllGenesisValidatorRequest{
				LaunchID:   launchID,
				Pagination: pageReq,
			}

			res, err := queryClient.GenesisValidatorAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowGenesisValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-genesis-validator [launch-id] [address]",
		Short: "shows a genesisValidator",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			launchID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetGenesisValidatorRequest{
				LaunchID: launchID,
				Address:  args[1],
			}

			res, err := queryClient.GenesisValidator(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
