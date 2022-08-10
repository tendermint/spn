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
		Use:   "add-validator-operator-address [operator-address]",
		Short: "Associate an validator operator address to a validator on SPN",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			validatorAddr := clientCtx.GetFromAddress().String()
			operatorAddr := args[0]

			msg := types.NewMsgSAddValidatorOperatorAddress(
				validatorAddr,
				operatorAddr,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			// initialize tx
			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags())
			txBuilder, err := txf.BuildUnsignedTx(msg)
			if err != nil {
				return err
			}

			// double sign with the operator address
			if err := addSignature(clientCtx, txf, txBuilder, validatorAddr); err != nil {
				return err
			}
			if operatorAddr != validatorAddr {
				if err := addSignature(clientCtx, txf, txBuilder, operatorAddr); err != nil {
					return err
				}
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
			fmt.Print(resp.String())

			return nil
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// addSignature add a signature to the tx builder
func addSignature(
	clientCtx client.Context,
	txf tx.Factory,
	txBuilder client.TxBuilder,
	addr string,
) error {
	fromAddr, fromName, _, err := client.GetFromFields(clientCtx, clientCtx.Keyring, addr)
	if err != nil {
		return err
	}
	clientCtx = clientCtx.WithFrom(addr).WithFromAddress(fromAddr).WithFromName(fromName)

	num, seq, err := txf.AccountRetriever().GetAccountNumberSequence(clientCtx, clientCtx.FromAddress)
	if err != nil {
		return err
	}
	txf = txf.WithAccountNumber(num)
	txf = txf.WithSequence(seq)

	return tx.Sign(txf, clientCtx.FromName, txBuilder, false)
}
