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

func CmdEditChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit-chain [chainID] [sourceURL] [sourceHash]",
		Short: "Edit a chain&#39;s information",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsChainID := string(args[0])
			argsSourceURL := string(args[1])
			argsSourceHash := string(args[2])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgEditChain(clientCtx.GetFromAddress().String(), string(argsChainID), string(argsSourceURL), string(argsSourceHash))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
