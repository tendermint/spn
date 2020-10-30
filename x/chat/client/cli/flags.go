package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	// FlagPayload describes a JSON string describing the payload data
	FlagPayload = "payload"

	// FlagDescription gives the description of a channel
	FlagDescription = "description"

	// FlagPollOptions describes the options for a new poll
	FlagPollOptions = "poll-options"

	// FlagTags describes the tags attached to a message
	FlagTags = "tags"
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

// FlagSetPollOptions returns the FlagSet for poll options
func FlagSetPollOptions() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagPollOptions, "", "Comma-separated list of options for a new poll")
	return fs
}

// FlagSetTags returns the FlagSet for tags
func FlagSetTags() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagTags, "", "Comma-separated list of tags")
	return fs
}
