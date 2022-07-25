package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/campaign/types"
)

func CmdListMainnetAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-mainnet-account [campaign-id]",
		Short: "List all mainnet accounts for a campaign",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			campaignID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllMainnetAccountRequest{
				CampaignID: campaignID,
				Pagination: pageReq,
			}

			res, err := queryClient.MainnetAccountAll(context.Background(), params)
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

func CmdShowMainnetAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-mainnet-account [campaign-id] [address]",
		Short: "Shows the mainnet account for a campaign",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argsCampaignID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argsAddress := args[1]

			params := &types.QueryGetMainnetAccountRequest{
				CampaignID: argsCampaignID,
				Address:    argsAddress,
			}

			res, err := queryClient.MainnetAccount(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListMainnetAccountBalance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-mainnet-account-balance [campaign-id]",
		Short: "List all mainnet account balances for a campaign",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			campaignID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllMainnetAccountBalanceRequest{
				CampaignID: campaignID,
				Pagination: pageReq,
			}

			res, err := queryClient.MainnetAccountBalanceAll(context.Background(), params)
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

func CmdShowMainnetAccountBalance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-mainnet-account-balance [campaign-id] [address]",
		Short: "Shows the mainnet account balance for a campaign",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argsCampaignID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argsAddress := args[1]

			params := &types.QueryGetMainnetAccountBalanceRequest{
				CampaignID: argsCampaignID,
				Address:    argsAddress,
			}

			res, err := queryClient.MainnetAccountBalance(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
