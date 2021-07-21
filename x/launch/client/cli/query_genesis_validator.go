package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/launch/types"
)

func CmdListGenesisValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-genesis-validator",
		Short: "list all genesisValidator",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllGenesisValidatorRequest{
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
		Use:   "show-genesis-validator [chainID] [address]",
		Short: "shows a genesisValidator",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argsChainID, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}
			argsAddress, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}

			params := &types.QueryGetGenesisValidatorRequest{
				ChainID: argsChainID,
				Address: argsAddress,
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
