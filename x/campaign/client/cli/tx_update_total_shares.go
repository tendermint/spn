package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/campaign/types"
)

func CmdUpdateTotalShares() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-total-shares [campaign-id] [total-shares]",
		Short: "Update the shares supply for the campaign",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			campaignID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			totalShares, err := types.NewShares(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateTotalShares(
				clientCtx.GetFromAddress().String(),
				campaignID,
				totalShares,
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
