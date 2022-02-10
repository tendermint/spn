package cli

import (
	"context"
	"encoding/base64"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/profile/types"
)

func CmdShowConsensusKeyNonce() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-consensus-key-nonce [consensus-address]",
		Short: "shows a ConsensusKeyNonce",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			consAddr, err := base64.StdEncoding.DecodeString(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetConsensusKeyNonceRequest{
				ConsensusAddress: consAddr,
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
