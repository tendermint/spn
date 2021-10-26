package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/campaign/types"
)

func CmdBurnVouchers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-vouchers [campaign-id] [vouchers]",
		Short: "Burn vouchers and decrease allocated shares of the campaign",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			campaignID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			vouchers, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBurnVouchers(
				clientCtx.GetFromAddress().String(),
				campaignID,
				vouchers,
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
