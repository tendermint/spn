package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/campaign/types"
)

func CmdEditCampaign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-campaign-name [campaign-id] [name]",
		Short: "Update the campaign name",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			campaignID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			metadata, err := cmd.Flags().GetString(flagMetadata)
			if err != nil {
				return err
			}
			metadataBytes := []byte(metadata)

			msg := types.NewMsgEditCampaign(
				clientCtx.GetFromAddress().String(),
				args[1],
				campaignID,
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
