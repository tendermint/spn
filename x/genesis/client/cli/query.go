package cli

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/genesis/types"
	"io/ioutil"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group genesis queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdListChains(),
		CmdShowChain(),
	)

	return cmd
}

// CmdListChains returns the command to list the chains
func CmdListChains() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-chains",
		Short: "show info concerning a channel",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			// Get page
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			params := &types.QueryListChainsRequest{
				Pagination: pageReq,
			}

			// Perform the request
			res, err := queryClient.ListChains(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "chains")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdShowChain returns the command to show a chain
func CmdShowChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-chain [chain-id]",
		Short: "show info concerning a chain to launch",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryShowChainRequest{
				ChainID: args[0],
			}

			res, err := queryClient.ShowChain(context.Background(), params)
			if err != nil {
				return err
			}

			// Check if the genesis must be save in a file
			genesisFile, _ := cmd.Flags().GetString(FlagSaveGenesis)
			if genesisFile != "" {
				genesis := res.Chain.Genesis
				ioutil.WriteFile(genesisFile, genesis, 0644)
			}

			return clientCtx.PrintOutput(res)
		},
	}

	cmd.Flags().AddFlagSet(FlagSetSaveGenesis())
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}