package cli

import (
	"errors"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/launch/types"
)

const (
	flagGenesisChainID = "genesis-chain-id"
	flagSourceURL      = "source-url"
	flagSourceHash     = "source-hash"
	flagDefaultGenesis = "default-genesis"
	flagMetadata       = "metadata"
	flagSetCampaignID  = "set-campaign-id"
	//flagModifyCampaignID = "set-campaign-id"
)

func CmdEditChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit-chain [launch-id]",
		Short: "Edit chain information",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				genesisChainID, _ = cmd.Flags().GetString(flagGenesisChainID)
				sourceURL, _      = cmd.Flags().GetString(flagSourceURL)
				sourceHash, _     = cmd.Flags().GetString(flagSourceHash)
				defaultGenesis, _ = cmd.Flags().GetBool(flagDefaultGenesis)
				genesisURL, _     = cmd.Flags().GetString(flagGenesisURL)
				metadata, _       = cmd.Flags().GetString(flagMetadata)
				campaignID, _     = cmd.Flags().GetUint64(flagSetCampaignID)
				//modifyCampaignID, _ = cmd.Flags().GetBool(flagModifyCampaignID)
			)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var initialGenesis *types.InitialGenesis
			if defaultGenesis && genesisURL != "" {
				return errors.New("the initial genesis can't be the default genesis and a custom genesis from URL at the same time")
			}
			if defaultGenesis {
				defaultInitialGenesis := types.NewDefaultInitialGenesis()
				initialGenesis = &defaultInitialGenesis
			} else if genesisURL != "" {
				genesisHash, err := getHashFromURL(cmd.Context(), genesisURL)
				if err != nil {
					return err
				}
				genesisURL := types.NewGenesisURL(genesisURL, genesisHash)
				initialGenesis = &genesisURL
			}

			modifyCampaignID := cmd.Flags().Changed(flagSetCampaignID)

			metadataBytes := []byte(metadata)

			launchID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgEditChain(
				clientCtx.GetFromAddress().String(),
				launchID,
				genesisChainID,
				sourceURL,
				sourceHash,
				initialGenesis,
				modifyCampaignID,
				campaignID,
				metadataBytes,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagGenesisChainID, "", "Chain ID for the genesis of the chain")
	cmd.Flags().String(flagSourceURL, "", "Set a new source URL for the chain")
	cmd.Flags().String(flagSourceHash, "", "Hash from the new source URL for the chain")
	cmd.Flags().Bool(flagDefaultGenesis, false, "Set the initial genesis to the default genesis of the chain")
	cmd.Flags().String(flagGenesisURL, "", "Set the initial genesis from a URL containing a custom genesis")
	cmd.Flags().String(flagMetadata, "", "Set metadata field for the chain")
	cmd.Flags().Uint64(flagSetCampaignID, 0, "Set the campaign ID if the chain is not associated with a campaign")
	// cmd.Flags().Bool(flagModifyCampaignID, false, "Enable to set the campaign ID of the chain")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
