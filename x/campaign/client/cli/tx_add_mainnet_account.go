package cli

import (
	"strconv"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/tendermint/spn/x/campaign/types"
)

var _ = strconv.Itoa(0)

func CmdAddMainnetAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-mainnet-account [campaign-id] [address] [shares]",
		Short: "Add a mainnet account",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argCampaignID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			argAddress := args[1]
			shares, err := types.NewShares(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddShares(
				argCampaignID,
				clientCtx.GetFromAddress().String(),
				argAddress,
				shares,
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
