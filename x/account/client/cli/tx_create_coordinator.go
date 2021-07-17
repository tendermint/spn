package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/tendermint/spn/x/account/types"
)

var _ = strconv.Itoa(0)

func CmdCreateCoordinator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-coordinator [address] [identity] [website] [details]",
		Short: "Broadcast message create-coordinator",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsAddress := string(args[0])
			argsIdentity := string(args[1])
			argsWebsite := string(args[2])
			argsDetails := string(args[3])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateCoordinator(clientCtx.GetFromAddress().String(), string(argsAddress), string(argsIdentity), string(argsWebsite), string(argsDetails))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
