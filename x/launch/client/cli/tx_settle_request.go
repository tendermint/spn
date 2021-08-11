package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/launch/types"
)

func CmdSettleRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "settle-request [approve|reject] [chainID] [requestID]",
		Short: "Approve or reject a pending request",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			approve, err := strconv.ParseBool(args[0])
			if err != nil {
				return err
			}

			requestID, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSettleRequest(
				clientCtx.GetFromAddress().String(),
				args[1],
				requestID,
				approve,
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
