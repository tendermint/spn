package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/project/types"
)

func CmdMintVouchers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint-vouchers [project-id] [shares]",
		Short: "Mint vouchers from project shares",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			projectID, err := cast.ToUint64E(args[0])
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

			msg := types.NewMsgMintVouchers(
				clientCtx.GetFromAddress().String(),
				projectID,
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
