package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/campaign/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group campaign queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdListCampaign(),
		CmdShowCampaign(),
		CmdShowCampaignChains(),
		CmdListMainnetAccount(),
		CmdShowMainnetAccount(),
		CmdListMainnetVestingAccount(),
		CmdShowMainnetVestingAccount(),
		CmdListCampaignSummary(),
		CmdShowCampaignSummary(),
		CmdQueryParams(),
	)

	// this line is used by starport scaffolding # 1

	return cmd
}
