package cli

import (
	"strconv"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/tendermint/spn/x/campaign/types"
)

var _ = strconv.Itoa(0)

func CmdUpdateTotalSupply() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-total-supply [campaign-id]",
		Short: "Update the total of the mainnet of a campaign",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argCampaignID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateTotalSupply(
				clientCtx.GetFromAddress().String(),
				argCampaignID,
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
