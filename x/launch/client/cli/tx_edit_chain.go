package cli

import (
	"errors"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	codec "github.com/cosmos/cosmos-sdk/codec/types"
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
		Use:   "edit-chain [chainID]",
		Short: "Edit chain information",
		Args:  cobra.ExactArgs(3),
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

			var initialGenesis *codec.Any
			if defaultGenesis && genesisURL != "" {
				return errors.New("the initial genesis can't be the default genesis and a custom genesis from URL at the same time")
			}
			if defaultGenesis {
				initialGenesis, err = codec.NewAnyWithValue(&types.DefaultInitialGenesis{})
				if err != nil {
					return err
				}
			} else if genesisURL != "" {
				// TODO: automatically determine this value by fetching the resource (need to determine the hash before)
				genesisHash, err := cmd.Flags().GetString(flagGenesisHash)
				if err != nil {
					return err
				}

				initialGenesis, err = codec.NewAnyWithValue(&types.GenesisURL{
					Url:  genesisURL,
					Hash: genesisHash,
				})
				if err != nil {
					return err
				}
			}

			msg := types.NewMsgEditChain(
				clientCtx.GetFromAddress().String(),
				args[0],
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
	cmd.Flags().String(flagGenesisHash, "", "hash of the content of the custom genesis from URL")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
