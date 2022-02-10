package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/campaign/types"
)

func CmdShowCampaignChains() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-campaign-chains [campaign-id]",
		Short: "shows a campaignChains",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argsCampaignID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetCampaignChainsRequest{
				CampaignID: argsCampaignID,
			}

			res, err := queryClient.CampaignChains(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
