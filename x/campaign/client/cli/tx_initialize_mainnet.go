package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/campaign/types"
)

func CmdInitializeMainnet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "initialize-mainnet [campaign-id] [source-url] [source-hash] [mainnet-chain-id]",
		Short: "Initialize the mainnet of the campaign to open gentxs submissions and fix total supply",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			campaignID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgInitializeMainnet(
				clientCtx.GetFromAddress().String(),
				campaignID,
				args[1],
				args[2],
				args[3],
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
