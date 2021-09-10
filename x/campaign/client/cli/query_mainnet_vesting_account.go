package cli

import (
	"context"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/tendermint/spn/x/campaign/types"
)

func CmdListMainnetVestingAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-mainnet-vesting-account",
		Short: "list all MainnetVestingAccount",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllMainnetVestingAccountRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.MainnetVestingAccountAll(context.Background(), params)
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

func CmdShowMainnetVestingAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-mainnet-vesting-account [campaign-id] [address]",
		Short: "shows a MainnetVestingAccount",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argsCampaignID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argsAddress := args[1]

			params := &types.QueryGetMainnetVestingAccountRequest{
				CampaignID: argsCampaignID,
				Address:    argsAddress,
			}

			res, err := queryClient.MainnetVestingAccount(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
