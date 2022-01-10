package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	ibctmtypes "github.com/cosmos/ibc-go/modules/light-clients/07-tendermint/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/monitoringc/types"
)

func CmdCreateClient() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-client [launch-id]",
		Short: "Create a verified client ID to connect to the chain with the specified launch ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			launchID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateClient(
				clientCtx.GetFromAddress().String(),
				launchID,
				ibctmtypes.ClientState{},
				ibctmtypes.ConsensusState{},
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
