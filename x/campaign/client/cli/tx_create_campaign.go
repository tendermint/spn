package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/campaign/types"
)

const (
	flagDynamicShares = "dynamic-shares"
)

func CmdCreateCampaign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-campaign [campaign-name] [total-supply]",
		Short: "Create a new campaign",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				dynamicShares, _ = cmd.Flags().GetBool(flagDynamicShares)
			)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			totalSupply := sdk.NewCoins()
			if len(args) > 1 {
				totalSupply, err = sdk.ParseCoinsNormalized(args[1])
				if err != nil {
					return err
				}
			}

			msg := types.NewMsgCreateCampaign(
				clientCtx.GetFromAddress().String(),
				args[0],
				totalSupply,
				dynamicShares,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Bool(flagDynamicShares, false, "Allows to update the shares supply for the mainnet coins supply")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
