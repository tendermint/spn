package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/project/types"
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
		CmdCreateProject(),
		CmdEditProject(),
		CmdUpdateTotalSupply(),
		CmdUpdateSpecialAllocations(),
		CmdInitializeMainnet(),
		CmdMintVouchers(),
		CmdBurnVouchers(),
		CmdUnredeemVouchers(),
		CmdRedeemVouchers(),
	)

	// this line is used by starport scaffolding # 1

	return cmd
}
