package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/tendermint/spn/app"
	"github.com/tendermint/spn/cmd/spnd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome("")); err != nil {
		os.Exit(1)
	}
}
