package cli

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/tendermint/spn/x/genesis/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdChainCreate(),
	)

	return cmd 
}

// CmdChainCreate returns the transaction command to create a new chain
func CmdChainCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-chain [chain-id] [source-URL] [source-hash] [genesis-file]",
		Short: "Create a new chain to launch",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			// Read genesis file
			genesisDoc, err := tmtypes.GenesisDocFromFile(args[3])
			if err != nil {
				return err
			}

			// Convert genesis
			genesis, err := json.Marshal(*genesisDoc)
			if err != nil {
				return err
			}

			// Create and send message
			msg := types.NewMsgChainCreate(
				args[0],
				clientCtx.GetFromAddress(),
				args[1],
				args[2],
				genesis,
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
