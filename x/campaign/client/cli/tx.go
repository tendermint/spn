package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/campaign/types"
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

	cmd.AddCommand(CmdCreateCampaign())
	cmd.AddCommand(CmdUpdateCampaignName())
	cmd.AddCommand(CmdUpdateTotalSupply())
	cmd.AddCommand(CmdUpdateTotalShares())
	cmd.AddCommand(CmdInitializeMainnet())
	cmd.AddCommand(CmdAddShares())
	cmd.AddCommand(CmdAddVestingOptions())
	cmd.AddCommand(CmdMintVouchers())
	cmd.AddCommand(CmdBurnVouchers())
	cmd.AddCommand(CmdUnredeemVouchers())
	cmd.AddCommand(CmdRedeemVouchers())
	// this line is used by starport scaffolding # 1

	return cmd
}
