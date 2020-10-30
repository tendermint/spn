package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	// FlagPayload describes a JSON string describing the payload data
	FlagPayload = "payload"

	// FlagDescription gives the description of a channel
	FlagDescription = "description"
)

// FlagSetPayload returns the FlagSet for payloads
func FlagSetPayload() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagPayload, "", "JSON string describing the payload data")
	return fs
}

// FlagSetDescription returns the FlagSet for description
func FlagSetDescription() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagDescription, "", "Description of the channel")
	return fs
}
