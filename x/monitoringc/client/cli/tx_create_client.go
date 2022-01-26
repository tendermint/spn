package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	spntypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/x/monitoringc/types"
)

func CmdCreateClient() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-client [launch-id] [consensus-state-file] [validator-set-file]",
		Short: "Create a verified client ID to connect to the chain with the specified launch ID",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			launchID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			cs, err := spntypes.ParseConsensusStateFromFile(args[1])
			if err != nil {
				return err
			}

			vs, err := spntypes.ParseValidatorSetFromFile(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateClient(
				clientCtx.GetFromAddress().String(),
				launchID,
				cs,
				vs,
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
