package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
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

	cmd.AddCommand(CmdCreateChain())
	cmd.AddCommand(CmdEditChain())
	cmd.AddCommand(CmdRequestAddAccount())
	cmd.AddCommand(CmdRequestAddVestingAccount())
	cmd.AddCommand(CmdRequestRemoveAccount())
	cmd.AddCommand(CmdRequestAddValidator())
	cmd.AddCommand(CmdRequestRemoveValidator())
	cmd.AddCommand(CmdSettleRequest())
	cmd.AddCommand(CmdTriggerLaunch())
	cmd.AddCommand(CmdRevertLaunch())
	// this line is used by starport scaffolding # 1

	return cmd
}
