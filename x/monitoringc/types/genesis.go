package types

import (
	"fmt"

	host "github.com/cosmos/ibc-go/modules/core/24-host"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:                           PortID,
		VerifiedClientIDList:             []VerifiedClientID{},
		ProviderClientIDList:             []ProviderClientID{},
		LaunchIDFromVerifiedClientIDList: []LaunchIDFromVerifiedClientID{},
		LaunchIDFromChannelIDList:        []LaunchIDFromChannelID{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	// Check for duplicated index in verifiedClientID
	verifiedClientIDIndexMap := make(map[string]struct{})

	for _, elem := range gs.VerifiedClientIDList {
		index := string(VerifiedClientIDKey(elem.LaunchID, elem.ClientID))
		if _, ok := verifiedClientIDIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for verifiedClientID")
		}
		verifiedClientIDIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in providerClientID
	providerClientIDIndexMap := make(map[string]struct{})

	for _, elem := range gs.ProviderClientIDList {
		index := string(ProviderClientIDKey(elem.LaunchID))
		if _, ok := providerClientIDIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for providerClientID")
		}
		providerClientIDIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in launchIDFromVerifiedClientID
	launchIDFromVerifiedClientIDIndexMap := make(map[string]struct{})

	for _, elem := range gs.LaunchIDFromVerifiedClientIDList {
		index := string(LaunchIDFromVerifiedClientIDKey(elem.ClientID))
		if _, ok := launchIDFromVerifiedClientIDIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for launchIDFromVerifiedClientID")
		}
		launchIDFromVerifiedClientIDIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in launchIDFromChannelID
	launchIDFromChannelIDIndexMap := make(map[string]struct{})

	for _, elem := range gs.LaunchIDFromChannelIDList {
		index := string(LaunchIDFromChannelIDKey(elem.ChannelID))
		if _, ok := launchIDFromChannelIDIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for launchIDFromChannelID")
		}
		launchIDFromChannelIDIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
