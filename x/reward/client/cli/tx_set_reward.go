package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/reward/types"
)

func CmdSetRewards() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-reward [launch-id] [coins] [last-reward-height]",
		Short: "Set rewards for being validator for a chain",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			launchID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			coins, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}
			lastRewardHeight, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetRewards(
				clientCtx.GetFromAddress().String(),
				launchID,
				lastRewardHeight,
				coins,
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
