package cli

import (
	"io/ioutil"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	valtypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/x/profile/types"
)

const (
	flagSignature = "signature"
)

func CmdSetValidatorConsAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-validator-cons-address [validator-key] [nonce]",
		Short: "Associate a Tendermint consensus address to a specific validator address on SPN",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Read validator key file
			valConsKeyBytes, err := ioutil.ReadFile(args[0])
			if err != nil {
				return err
			}
			valConsKey, err := valtypes.LoadValidatorKey(valConsKeyBytes)
			if err != nil {
				return err
			}

			// check if the signature flag exists, if not, create the signature based in the nonce
			signature, _ := cmd.Flags().GetString(flagSignature)
			if signature == "" {
				nonce, err := strconv.ParseUint(args[1], 10, 64)
				if err != nil {
					return err
				}
				signature, err = valConsKey.Sign(nonce, clientCtx.ChainID)
				if err != nil {
					return err
				}
			}

			msg := types.NewMsgSetValidatorConsAddress(
				clientCtx.GetFromAddress().String(),
				signature,
				valConsKey.PubKey.Type(),
				valConsKey.PubKey.Bytes(),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagSignature, "", "signature of the nonce and chain id. Using this arg invalidate the nonce parameter")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
