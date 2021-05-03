package cli

import (
	"context"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/internal/utils"
	"github.com/tendermint/spn/x/launch/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group launch queries under a subcommand
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
		CmdListProposals(),
		CmdLaunchInformation(),
		CmdSimulatedLaunchInformation(),
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

			return clientCtx.PrintProto(res)
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
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryShowChainRequest{
				ChainID: args[0],
			}

			res, err := queryClient.ShowChain(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
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
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryProposalCountRequest{
				ChainID: args[0],
			}

			res, err := queryClient.ProposalCount(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdShowProposal returns the command to show a proposal
func CmdShowProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-proposal [chain-id] [proposal-id]",
		Short: "show info concerning a proposal for a chain launch",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			// Convert proposal ID
			proposalID, err := utils.ParseInt32(args[1])
			if err != nil {
				return err
			}

			params := &types.QueryShowProposalRequest{
				ChainID:    args[0],
				ProposalID: proposalID,
			}

			res, err := queryClient.ShowProposal(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdListProposals returns the command to list proposals for a chain launch
func CmdListProposals() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-proposals [chain-id]",
		Short: "list the  proposals for a chain launch",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryListProposalsRequest{
				ChainID: args[0],
				Status:  types.ProposalStatus_ANY_STATUS,
				Type:    types.ProposalType_ANY_TYPE,
			}

			// Perform the request
			res, err := queryClient.ListProposals(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdLaunchInformation returns the command to show the information to launch a chain
func CmdLaunchInformation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "launch-information [chain-id]",
		Short: "show the information to launch a chain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryLaunchInformationRequest{
				ChainID: args[0],
			}

			// Perform the request
			res, err := queryClient.LaunchInformation(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdSimulatedLaunchInformation returns the command to show the simulated information to launch a chain with a pending proposal
func CmdSimulatedLaunchInformation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "simulated-launch-information [chain-id] [proposal-ids]",
		Short: "show the simulated launch information with comma-separated proposals to test",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			var proposalIDs []int32
			proposalIDsArg := strings.Split(args[1], ",")
			for _, proposalIDArg := range proposalIDsArg {
				// Convert proposal ID
				proposalID, err := utils.ParseInt32(proposalIDArg)
				if err != nil {
					return err
				}

				proposalIDs = append(proposalIDs, proposalID)
			}

			params := &types.QuerySimulatedLaunchInformationRequest{
				ChainID:     args[0],
				ProposalIDs: proposalIDs,
			}

			// Perform the request
			res, err := queryClient.SimulatedLaunchInformation(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
