package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/tendermint/spn/x/launch/types"
)

var _ = strconv.Itoa(0)

func CmdRequestRemoveValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-remove-validator [chainID] [validatorAddress]",
		Short: "Request to remove a validator",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsChainID := string(args[0])
			argsValidatorAddress := string(args[1])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRequestRemoveValidator(clientCtx.GetFromAddress().String(), string(argsChainID), string(argsValidatorAddress))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
