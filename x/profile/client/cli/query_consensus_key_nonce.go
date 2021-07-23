package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/profile/types"
)

func CmdShowConsensusKeyNonce() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-consensus-key-nonce [consAddress]",
		Short: "shows the currently required nonce to sign a consensus key proof of ownership",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argsConsAddress, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetConsensusKeyNonceRequest{
				ConsAddress: argsConsAddress,
			}

			res, err := queryClient.ConsensusKeyNonce(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
