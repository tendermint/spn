package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/ignite/cli/ignite/pkg/cosmoscmd"

	"github.com/tendermint/spn/app"
	spntypes "github.com/tendermint/spn/pkg/types"
)

func main() {
	rootCmd, _ := cosmoscmd.NewRootCmd(
		spntypes.Name,
		spntypes.AccountAddressPrefix,
		app.DefaultNodeHome,
		spntypes.DefaultChainID,
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
	)
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
