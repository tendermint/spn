package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/launch/types"
)

var _ = strconv.Itoa(0)

func CmdTriggerLaunch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trigger-launch [chainID] [remainingTime]",
		Short: "Trigger the launch of a chain",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsRemainingTime, _ := strconv.ParseUint(args[1], 10, 64)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTriggerLaunch(clientCtx.GetFromAddress().String(), args[0], argsRemainingTime)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
