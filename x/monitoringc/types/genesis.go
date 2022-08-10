package types

import (
	"fmt"

	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
)

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:                           PortID,
		VerifiedClientIDList:             []VerifiedClientID{},
		ProviderClientIDList:             []ProviderClientID{},
		LaunchIDFromVerifiedClientIDList: []LaunchIDFromVerifiedClientID{},
		LaunchIDFromChannelIDList:        []LaunchIDFromChannelID{},
		MonitoringHistoryList:            []MonitoringHistory{},
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
	clientIDMap := make(map[string]struct{})
	for _, elem := range gs.VerifiedClientIDList {
		index := string(VerifiedClientIDKey(elem.LaunchID))
		if _, ok := verifiedClientIDIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for verifiedClientID")
		}
		verifiedClientIDIndexMap[index] = struct{}{}

		// Check for duplicated client id
		for _, clientID := range elem.ClientIDs {
			key := clientIDKey(elem.LaunchID, clientID)
			if _, ok := clientIDMap[key]; ok {
				return fmt.Errorf("duplicated client id")
			}
			clientIDMap[key] = struct{}{}
		}
	}

	// Check for duplicated index in providerClientID
	providerClientIDIndexMap := make(map[string]struct{})
	for _, elem := range gs.ProviderClientIDList {
		index := string(ProviderClientIDKey(elem.LaunchID))
		if _, ok := providerClientIDIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for providerClientID")
		}
		// Check if the client id exist
		key := clientIDKey(elem.LaunchID, elem.ClientID)
		if _, ok := clientIDMap[key]; !ok {
			return fmt.Errorf("client id from providerClientID list not found")
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
		// Check if the client id exist
		key := clientIDKey(elem.LaunchID, elem.ClientID)
		if _, ok := clientIDMap[key]; !ok {
			return fmt.Errorf("client id from launchIDFromVerifiedClientID list not found")
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

	// Check for duplicated index in monitoringHistory
	monitoringHistoryIndexMap := make(map[string]struct{})
	for _, elem := range gs.MonitoringHistoryList {
		index := string(MonitoringHistoryKey(elem.LaunchID))
		if _, ok := monitoringHistoryIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for monitoringHistory")
		}
		monitoringHistoryIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

// clientIDKey creates a string key for launch id and client id
func clientIDKey(launchID uint64, clientID string) string {
	return fmt.Sprintf("%d-%s", launchID, clientID)
}
