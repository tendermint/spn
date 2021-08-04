package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/launch/types"
)

const (
	flagChainID = "chain-id"
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

			chainID, _ := cmd.Flags().GetString(flagChainID)

			params := &types.QueryAllGenesisAccountRequest{
				ChainID:    chainID,
				Pagination: pageReq,
			}

			res, err := queryClient.GenesisAccountAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String(flagChainID, "", "filter by chain id")
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

			params := &types.QueryGetGenesisAccountRequest{
				ChainID: args[0],
				Address: args[1],
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
