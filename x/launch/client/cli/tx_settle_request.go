package cli

import (
	"fmt"
	"strconv"
	"strings"

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

			ids, err := parseList(args[1])
			if err != nil {
				return err
			}

			approve, err := strconv.ParseBool(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSettleRequest(
				clientCtx.GetFromAddress().String(),
				args[0],
				ids,
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

// parseList parses comma separated numbers and range to []uint64.
func parseList(arg string) ([]uint64, error) {
	list := make([]uint64, 0)
	for _, numberRange := range strings.Split(arg, ",") {
		trimmedRange := strings.TrimSpace(numberRange)
		if trimmedRange == "" {
			continue
		}

		numbers := strings.Split(trimmedRange, "/")
		switch len(numbers) {
		case 1:
			trimmed := strings.TrimSpace(numbers[0])
			i, err := strconv.ParseUint(trimmed, 10, 32)
			if err != nil {
				return nil, err
			}
			list = append(list, i)
		case 2:
			var (
				startN = strings.TrimSpace(numbers[0])
				endN   = strings.TrimSpace(numbers[1])
			)
			if startN == "" {
				startN = endN
			}
			if endN == "" {
				endN = startN
			}
			if startN == "" {
				continue
			}
			start, err := strconv.ParseUint(startN, 10, 32)
			if err != nil {
				return nil, err
			}
			end, err := strconv.ParseUint(endN, 10, 32)
			if err != nil {
				return nil, err
			}
			if start > end {
				start, end = end, start
			}
			for ; start <= end; start++ {
				list = append(list, start)
			}
		default:
			return nil, fmt.Errorf("cannot parse the number range: %s", trimmedRange)
		}
	}
	return list, nil
}
