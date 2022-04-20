package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/launch/types"
)

const (
	flagMetadata = "metadata"
)

func CmdEditChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit-chain [launch-id]",
		Short: "Edit chain information",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				metadata, _   = cmd.Flags().GetString(flagMetadata)
				campaignID, _ = cmd.Flags().GetUint64(flagCampaignID)
			)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			setCampaignID := cmd.Flags().Changed(flagCampaignID)

			metadataBytes := []byte(metadata)

			launchID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgEditChain(
				clientCtx.GetFromAddress().String(),
				launchID,
				setCampaignID,
				campaignID,
				metadataBytes,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagMetadata, "", "Set metadata field for the chain")
	cmd.Flags().Uint64(flagCampaignID, 0, "Set the campaign ID if the chain is not associated with a campaign")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
