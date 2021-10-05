package cli

import (
	"strconv"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/campaign/types"
)

var _ = strconv.Itoa(0)

const (
	flagAccount = "account"
)

func CmdRedeemVouchers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "redeem-vouchers [campaign-id] [vouchers]",
		Short: "Redeem vouchers and allocate shares for an account in the mainnet of the campaign",
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

			account, err := cmd.Flags().GetString(flagAccount)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			if account == "" {
				account = clientCtx.GetFromAddress().String()
			}

			msg := types.NewMsgRedeemVouchers(
				clientCtx.GetFromAddress().String(),
				campaignID,
				account,
				vouchers,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagAccount, "", "Account address that receives shares allocation from redeemed vouchers")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
