package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/launch/types"
)

func CmdListChainNameCount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-chain-name-count",
		Short: "list all chainNameCount",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllChainNameCountRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ChainNameCountAll(context.Background(), params)
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

func CmdShowChainNameCount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-chain-name-count [chainName]",
		Short: "shows a chainNameCount",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argsChainName, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetChainNameCountRequest{
				ChainName: argsChainName,
			}

			res, err := queryClient.ChainNameCount(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
