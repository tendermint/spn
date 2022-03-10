package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/participation/types"
)

func CmdListAuctionUsedAllocations() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-auction-used-allocations [address]",
		Short: "List all used allocations for auctions for an address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			argAddress := args[0]

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllAuctionUsedAllocationsRequest{
				Address:    argAddress,
				Pagination: pageReq,
			}

			res, err := queryClient.AuctionUsedAllocationsAll(context.Background(), params)
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

func CmdShowAuctionUsedAllocations() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-auction-used-allocations [address] [auction-id]",
		Short: "Shows used allocations for an auction ",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argAddress := args[0]
			argAuctionID, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetAuctionUsedAllocationsRequest{
				Address:   argAddress,
				AuctionID: argAuctionID,
			}

			res, err := queryClient.AuctionUsedAllocations(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
