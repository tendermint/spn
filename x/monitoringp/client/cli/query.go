package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/tendermint/spn/x/monitoringp/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group monitoringp queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.FullModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdShowConsumerClientID(),
		CmdShowConnectionChannelID(),
		CmdShowMonitoringInfo(),
		CmdQueryParams(),
	)

	// this line is used by starport scaffolding # 1
	return cmd
}
