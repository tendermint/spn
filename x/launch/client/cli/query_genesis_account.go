package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/launch/types"
)

func CmdListGenesisAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-genesis-account",
		Short: "list all genesisAccount",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllGenesisAccountRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.GenesisAccountAll(context.Background(), params)
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

func CmdShowGenesisAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-genesis-account [chainID] [address]",
		Short: "shows a genesisAccount",
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

			params := &types.QueryGetGenesisAccountRequest{
				ChainID: argsChainID,
				Address: argsAddress,
			}

			res, err := queryClient.GenesisAccount(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
