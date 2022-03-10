package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/profile/types"
)

func CmdAddValidatorOperatorAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-validator-operator-address [validator-address]",
		Short: "Associate an validator operator address to a validator on SPN",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSAddValidatorOperatorAddress(
				args[0],
				clientCtx.GetFromAddress().String(),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			// return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)


			// sign the tx with both accounts
			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags())

			num, seq, err := txf.AccountRetriever().GetAccountNumberSequence(clientCtx, clientCtx.FromAddress)
			if err != nil {
				return err
			}
			txf = txf.WithAccountNumber(num)
			txf = txf.WithSequence(seq)

			txBuilder, err := tx.BuildUnsignedTx(txf, msg)
			if err != nil {
				return err
			}
			if err := tx.Sign(txf, clientCtx.FromName, txBuilder, false); err != nil {
				return err
			}

			// encode tx
			encoder := clientCtx.TxConfig.TxEncoder()
			encodedTx, err := encoder(txBuilder.GetTx())
			if err != nil {
				return err
			}

			// broadcast tx
			resp, err := clientCtx.BroadcastTxSync(encodedTx)
			if err != nil {
				return err
			}
			fmt.Println(resp.String())

			return nil
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
