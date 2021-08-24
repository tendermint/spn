package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/launch/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group launch queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdShowChain())
	cmd.AddCommand(CmdListChain())
	cmd.AddCommand(CmdShowGenesisValidator())
	cmd.AddCommand(CmdListGenesisValidator())
	cmd.AddCommand(CmdShowVestedAccount())
	cmd.AddCommand(CmdListVestedAccount())
	cmd.AddCommand(CmdShowGenesisAccount())
	cmd.AddCommand(CmdListGenesisAccount())
	cmd.AddCommand(CmdShowRequest())
	cmd.AddCommand(CmdListRequest())
	cmd.AddCommand(CmdQueryParams())
	// this line is used by starport scaffolding # 1

	return cmd
}
