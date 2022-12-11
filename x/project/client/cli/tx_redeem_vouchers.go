package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/project/types"
)

const flagAccount = "account"

func CmdRedeemVouchers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "redeem-vouchers [project-id] [vouchers]",
		Short: "Redeem vouchers and allocate shares for an account in the mainnet of the project",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			projectID, err := cast.ToUint64E(args[0])
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
				account,
				projectID,
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
