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

func CmdUnredeemVouchers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unredeem-vouchers [campaign-id] [shares]",
		Short: "Unredeem vouchers that has been redeemed into an account and get vouchers back",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCampaignID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			shares, err := types.NewShares(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUnredeemVouchers(
				clientCtx.GetFromAddress().String(),
				argCampaignID,
				shares,
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
