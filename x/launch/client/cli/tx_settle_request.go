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

func CmdSettleRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "settle-request [chainID] [requestID] [approve]",
		Short: "Approve or reject a pending request",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsRequestID, _ := strconv.ParseUint(args[1], 10, 64)
			argsApprove, _ := strconv.ParseBool(args[2])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSettleRequest(
				clientCtx.GetFromAddress().String(),
				args[0],
				argsRequestID,
				argsApprove,
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
