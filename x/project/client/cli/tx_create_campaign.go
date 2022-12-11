package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/project/types"
)

const (
	flagMetadata = "metadata"
)

func CmdCreateCampaign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-campaign [campaign-name] [total-supply]",
		Short: "Create a new campaign",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
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

			metadata, err := cmd.Flags().GetString(flagMetadata)
			if err != nil {
				return err
			}
			metadataBytes := []byte(metadata)

			msg := types.NewMsgCreateCampaign(
				clientCtx.GetFromAddress().String(),
				args[0],
				totalSupply,
				metadataBytes,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagMetadata, "", "Set metadata field for the campaign")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
