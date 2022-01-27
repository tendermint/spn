package cli

import (
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/profile/types"
)

func CmdSetValidatorConsAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-validator-cons-address [validator-key] [signature]",
		Short: "Associate a consensus address for a specific SPN address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Read validator key file
			valKeyBytes, err := ioutil.ReadFile(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgSetValidatorConsAddress(
				args[1],
				valKeyBytes,
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
