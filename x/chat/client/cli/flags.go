package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	// FlagPayload describes a JSON string describing the payload data
	FlagPayload = "payload"
)

// FlagSetPaylaod Returns the FlagSet for payloads
func FlagSetPaylaod() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagPayload, "", "JSON string describing the payload data")
	return fs
}
