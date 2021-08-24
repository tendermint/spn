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
	flagSourceURL      = "source-url"
	flagSourceHash     = "source-hash"
	flagDefaultGenesis = "default-genesis"
)

func CmdEditChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit-chain [chain-id]",
		Short: "Edit chain information",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sourceURL, err := cmd.Flags().GetString(flagSourceURL)
			if err != nil {
				return err
			}
			sourceHash, err := cmd.Flags().GetString(flagSourceHash)
			if err != nil {
				return err
			}
			defaultGenesis, err := cmd.Flags().GetBool(flagDefaultGenesis)
			if err != nil {
				return err
			}
			genesisURL, err := cmd.Flags().GetString(flagGenesisURL)
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

			chainID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgEditChain(
				clientCtx.GetFromAddress().String(),
				chainID,
				sourceURL,
				sourceHash,
				initialGenesis,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagSourceURL, "", "Set a new source URL for the chain")
	cmd.Flags().String(flagSourceHash, "", "Hash from the new source URL for the chain")
	cmd.Flags().Bool(flagDefaultGenesis, false, "Set the initial genesis to the default genesis of the chain")
	cmd.Flags().String(flagGenesisURL, "", "Set the initial genesis from a URL containing a custom genesis")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
