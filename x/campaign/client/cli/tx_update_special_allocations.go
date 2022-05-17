package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/campaign/types"
)

func CmdUpdateSpecialAllocations() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-special-allocations [campaign-id] [genesis-distribution] [claimable-airdrop]",
		Short: "update the special allocations for the campaign",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCampaignID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			genesisDistribution, err := types.NewShares(args[1])
			if err != nil {
				return err
			}

			claimableAirdrop, err := types.NewShares(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateSpecialAllocations(
				clientCtx.GetFromAddress().String(),
				argCampaignID,
				types.NewSpecialAllocations(genesisDistribution, claimableAirdrop),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
