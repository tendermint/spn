package cli

import (
	"fmt"
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	tmjson "github.com/tendermint/tendermint/libs/json"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/client"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"

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
		CmdProposalAddAccount(),
		CmdProposalAddValidator(),
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
			genesis, err := tmjson.Marshal(*genesisDoc)
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

// CmdProposalAddAccount returns the transaction command to add a new account into the genesis
func CmdProposalAddAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposal-add-account [chain-id] [coins]",
		Short: "Add a proposal to add a genesis account, [coins] must be comma separated coin denominations: 1000atom,1000stake",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			// Parse coins
			coins, err := sdk.ParseCoins(args[1])
			if err != nil {
				return fmt.Errorf("failed to parse coins: %v", err.Error())
			}

			// Construct payload
			payload := types.NewProposalAddAccountPayload(
				clientCtx.GetFromAddress(),
				coins,
			)

			// Create and send message
			msg := types.NewMsgProposalAddAccount(
				args[0],
				clientCtx.GetFromAddress(),
				payload,
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

// CmdProposalAddValidator returns the transaction command to add a new validator into the genesis
func CmdProposalAddValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposal-add-validator [chain-id] [peer] [gentx-file]",
		Short: "Add a proposal to add a gentx to add a validator during chain initialization",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			// Read gentxFile
			gentxBytes, err := ioutil.ReadFile(args[2])
			if err != nil {
				return err
			}

			// Parse gentx
			var gentx txtypes.Tx
			err = clientCtx.JSONMarshaler.UnmarshalJSON(gentxBytes, &gentx)
			if err != nil {
				return err
			}

			// Construct payload
			payload := types.NewProposalAddValidatorPayload(
				gentx,
				args[1],
			)

			// Create and send message
			msg := types.NewMsgProposalAddValidator(
				args[0],
				clientCtx.GetFromAddress(),
				payload,
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
