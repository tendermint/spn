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

func CmdCreateChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-chain [chainName] [sourceURL] [sourceHash] [genesisURL] [genesisHash]",
		Short: "Broadcast message create-chain",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsChainName := string(args[0])
			argsSourceURL := string(args[1])
			argsSourceHash := string(args[2])
			argsGenesisURL := string(args[3])
			argsGenesisHash := string(args[4])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateChain(clientCtx.GetFromAddress().String(), string(argsChainName), string(argsSourceURL), string(argsSourceHash), string(argsGenesisURL), string(argsGenesisHash))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
