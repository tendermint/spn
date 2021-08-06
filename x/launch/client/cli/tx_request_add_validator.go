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

func CmdRequestAddValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-add-validator [chainID] [consPubKey] [peer]",
		Short: "Send a request for a genesis validator",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsChainID := string(args[0])
			argsConsPubKey := string(args[1])
			argsPeer := string(args[2])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRequestAddValidator(clientCtx.GetFromAddress().String(), string(argsChainID), string(argsConsPubKey), string(argsPeer))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
