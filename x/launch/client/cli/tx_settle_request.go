package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/launch/types"
)

var (
	approveMap = map[string]bool{
		"approve": true,
		"reject":  false,
	}
)

func CmdSettleRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "settle-request [approve|reject] [launch-id] [request-id]",
		Short: "Approve or reject a pending request",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			approve, ok := approveMap[args[0]]
			if !ok {
				return fmt.Errorf(
					"invalid approve type '%s'. approvals must be %v",
					args[0], approveMap,
				)
			}

			requestID, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			launchID, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgSettleRequest(
				clientCtx.GetFromAddress().String(),
				launchID,
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
