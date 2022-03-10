package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/profile/types"
)

const (
	flagOperatorAddress = "operator-address"
)

func CmdAddValidatorOperatorAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-validator-operator-address",
		Short: "Associate an validator operator address to a validator on SPN",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			validatorAddr := clientCtx.GetFromAddress().String()
			operatorAddr, _ := cmd.Flags().GetString(flagOperatorAddress)
			if operatorAddr == "" {
				operatorAddr = validatorAddr
			}

			msg := types.NewMsgSAddValidatorOperatorAddress(
				validatorAddr,
				operatorAddr,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			// initialize tx
			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags())
			txBuilder, err := tx.BuildUnsignedTx(txf, msg)
			if err != nil {
				return err
			}

			if err := addSignature(clientCtx, txf, txBuilder, validatorAddr); err != nil {
				return err
			}

			// double sign if the operator address is different from the SPN validator address
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
			fmt.Println(resp.String())

			return nil
		},
	}

	cmd.Flags().String(flagOperatorAddress, "", "validator operator address")
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
	fromAddr, fromName, _, err := client.GetFromFields(clientCtx.Keyring, addr, clientCtx.GenerateOnly)
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
