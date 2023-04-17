package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/project/types"
)

func CmdListMainnetAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-mainnet-account [project-id]",
		Short: "List all mainnet accounts for a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			projectID, err := strconv.ParseUint(args[0], 10, 64)
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
				ProjectID:  projectID,
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
		Use:   "show-mainnet-account [project-id] [address]",
		Short: "Shows the mainnet account for a project",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argsProjectID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argsAddress := args[1]

			params := &types.QueryGetMainnetAccountRequest{
				ProjectID: argsProjectID,
				Address:   argsAddress,
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
		Use:   "list-mainnet-account-balance [project-id]",
		Short: "List all mainnet account balances for a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			projectID, err := strconv.ParseUint(args[0], 10, 64)
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
				ProjectID:  projectID,
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
		Use:   "show-mainnet-account-balance [project-id] [address]",
		Short: "Shows the mainnet account balance for a project",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argsProjectID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argsAddress := args[1]

			params := &types.QueryGetMainnetAccountBalanceRequest{
				ProjectID: argsProjectID,
				Address:   argsAddress,
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
