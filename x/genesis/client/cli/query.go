package cli

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/genesis/types"
	"strconv"
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
		CmdProposal(),
		CmdShowProposal(),
		CmdPendingProposals(),
		CmdRejectedProposals(),
		CmdApprovedProposals(),
		CmdLaunchInformation(),
	)

	return cmd
}

// CmdListChains returns the command to list the chains
func CmdListChains() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-chains",
		Short: "list the chains to launch",
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

			return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdProposal returns the command to show the number of proposals
func CmdProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposal-count [chain-id]",
		Short: "number of proposals for the chain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryProposalCountRequest{
				ChainID: args[0],
			}

			res, err := queryClient.ProposalCount(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdShowProposal returns the command to show a proposal
func CmdShowProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-proposal [chain-id] [proposal-id]",
		Short: "show info concerning a proposal for a chain genesis",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			// Convert proposal ID
			proposalID, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			params := &types.QueryShowProposalRequest{
				ChainID:    args[0],
				ProposalID: int32(proposalID),
			}

			res, err := queryClient.ShowProposal(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdPendingProposals returns the command to list pending proposals for a chain genesis
func CmdPendingProposals() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pending-proposals [chain-id]",
		Short: "list the pending proposals for a chain genesis",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryPendingProposalsRequest{
				ChainID: args[0],
			}

			// Perform the request
			res, err := queryClient.PendingProposals(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdApprovedProposals returns the command to list approved proposals for a chain genesis
func CmdApprovedProposals() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approved-proposals [chain-id]",
		Short: "list the approved proposals for a chain genesis",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryApprovedProposalsRequest{
				ChainID: args[0],
			}

			// Perform the request
			res, err := queryClient.ApprovedProposals(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdRejectedProposals returns the command to list rejected proposals for a chain genesis
func CmdRejectedProposals() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rejected-proposals [chain-id]",
		Short: "list the rejected proposals for a chain genesis",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryRejectedProposalsRequest{
				ChainID: args[0],
			}

			// Perform the request
			res, err := queryClient.RejectedProposals(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdLaunchInformation returns the command to show the information to launch a chain
func CmdLaunchInformation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "launch-information [chain-id]",
		Short: "show the information to launch a chain (genesis information and peers)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryLaunchInformationRequest{
				ChainID: args[0],
			}

			// Perform the request
			res, err := queryClient.LaunchInformation(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}