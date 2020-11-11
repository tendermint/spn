package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	// FlagSaveGenesis tells the file where to save the genesis on showing a chain
	FlagSaveGenesis = "save-genesis"
)

// FlagSetSaveGenesis returns the FlagSet for saving genesis
func FlagSetSaveGenesis() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagSaveGenesis, "", "file where to save the genesis")
	return fs
}
