package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/tendermint/spn/x/profile/types"
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
		CmdUpdateValidatorDescription(),
		CmdCreateCoordinator(),
		CmdUpdateCoordinatorDescription(),
		CmdUpdateCoordinatorAddress(),
		CmdDisableCoordinator(),
		CmdAddValidatorOperatorAddress(),
	)
	// this line is used by starport scaffolding # 1

	return cmd
}
