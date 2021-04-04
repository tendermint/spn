package cli

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/launch/types"
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
		CmdApprove(),
		CmdReject(),
		CmdProposalAddAccount(),
		CmdProposalAddValidator(),
	)

	return cmd
}

const genesisURLFlag = "genesis"

// CmdChainCreate returns the transaction command to create a new chain
func CmdChainCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-chain [chain-id] [source-URL] [source-hash]",
		Short: "Create a new chain to launch",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// If custom initial genesis, try out url and compute genesis hash
			var genesisHash string
			genesisURL, err := cmd.Flags().GetString(genesisURLFlag)
			if err != nil {
				return err
			}
			if genesisURL != "" {
				res, err := http.Get(genesisURL)
				if err != nil {
					return err
				}
				if res.StatusCode != http.StatusOK {
					return fmt.Errorf("genesis url fetch error %s", res.Status)
				}
				initialGenesis, err := ioutil.ReadAll(res.Body)
				if err != nil {
					return err
				}
				genesisHash = types.GenesisURLHash(initialGenesis)
			}

			// Create and send message
			msg := types.NewMsgChainCreate(
				args[0],
				clientCtx.GetFromAddress(),
				args[1],
				args[2],
				genesisURL,
				genesisHash,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().String(genesisURLFlag, "", "URL to a custom initial genesis")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdApprove returns the transaction command to approve a specific proposal
func CmdApprove() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposal-approve [chain-id] [proposal-id]",
		Short: "Approve a proposal",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Convert value for proposal ID
			proposalID, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			// Create and send message
			msg := types.NewMsgApprove(
				args[0],
				int32(proposalID),
				clientCtx.GetFromAddress(),
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

// CmdApprove returns the transaction command to reject a specific proposal
func CmdReject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposal-reject [chain-id] [proposal-id]",
		Short: "Reject a proposal",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Convert value for proposal ID
			proposalID, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			// Create and send message
			msg := types.NewMsgReject(
				args[0],
				int32(proposalID),
				clientCtx.GetFromAddress(),
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
		Use:   "proposal-add-account [chain-id] [address] [coins]",
		Short: "Add a proposal to add a genesis account, [coins] must be comma separated coin denominations: 1000atom,1000stake",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Parse address
			address, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			// Parse coins
			coins, err := sdk.ParseCoinsNormalized(args[2])
			if err != nil {
				return fmt.Errorf("failed to parse coins: %v", err.Error())
			}

			// Construct payload
			payload := types.NewProposalAddAccountPayload(
				address,
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
		Use:   "proposal-add-validator [chain-id] [peer] [address] [self-delegation] [gentx-file]",
		Short: "Add a proposal to add a gentx to add a validator during chain initialization",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Parse address
			address, err := sdk.AccAddressFromBech32(args[2])
			if err != nil {
				return err
			}

			// Read self-delegation
			selfDelegation, err := sdk.ParseCoinNormalized(args[3])
			if err != nil {
				return err
			}

			// Read gentxFile
			gentxBytes, err := ioutil.ReadFile(args[4])
			if err != nil {
				return err
			}

			// Construct payload
			payload := types.NewProposalAddValidatorPayload(
				gentxBytes,
				sdk.ValAddress(address),
				selfDelegation,
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
