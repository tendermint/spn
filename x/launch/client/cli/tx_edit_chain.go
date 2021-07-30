package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/launch/types"
)

const (
	flagEditInitialGenesis  = "edit-initial-genesis"
)

func CmdEditChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit-chain [chainID] [sourceURL] [sourceHash]",
		Short: "Edit chain information",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			genesisURL, err := cmd.Flags().GetString(flagGenesisURL)
			if err != nil {
				return err
			}

			// TODO: automatically determine this value by fetching the resource (need to determine the hash before)
			genesisHash, err := cmd.Flags().GetString(flagGenesisHash)
			if err != nil {
				return err
			}

			msg := types.NewMsgEditChain(
				clientCtx.GetFromAddress().String(),
				args[0],
				args[1],
				args[2],
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagGenesisURL, "", "URL for a custom genesis")
	cmd.Flags().String(flagGenesisHash, "", "hash of the content of the custom genesis")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
