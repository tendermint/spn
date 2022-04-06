package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/monitoringc/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group monitoringc queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.FullModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdShowVerifiedClientIds(),
		CmdShowProviderClientID(),
		CmdListProviderClientID(),
		CmdShowLaunchIDFromChannelID(),
		CmdListLaunchIDFromChannelID(),
		CmdShowMonitoringHistory(),
		CmdQueryParams(),
	)

	// this line is used by starport scaffolding # 1

	return cmd
}
