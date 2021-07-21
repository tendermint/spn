package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/launch/types"
)

func CmdListRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-request [chainID]",
		Short: "list all request",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argsChainID, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryAllRequestRequest{
				ChainID:   argsChainID,
				Pagination: pageReq,
			}

			res, err := queryClient.RequestAll(context.Background(), params)
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

func CmdShowRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-request [chainID] [requestID]",
		Short: "shows a request",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argsChainID, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}
			argsRequestID, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			params := &types.QueryGetRequestRequest{
				ChainID:   argsChainID,
				RequestID: argsRequestID,
			}

			res, err := queryClient.Request(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
