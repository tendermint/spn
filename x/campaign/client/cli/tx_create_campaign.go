package cli

import (
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/tendermint/spn/x/campaign/types"
)

const (
	flagDynamicShares = "dynamic-shares"
)

func CmdCreateCampaign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-campaign [campaign-name] [total-supply]",
		Short: "Create a new campaign",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argCampaignName := args[0]

			argDynamicShares, err := cast.ToBoolE(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateCampaign(
				clientCtx.GetFromAddress().String(),
				argCampaignName,
				argDynamicShares,
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
