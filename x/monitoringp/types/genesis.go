package types

import (
	host "github.com/cosmos/ibc-go/modules/core/24-host"
	// this line is used by starport scaffolding # genesis/types/import
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:              PortID,
		ConsumerClientID:    nil,
		ConnectionChannelID: nil,
		MonitoringInfo:      nil,
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
	// this line is used by starport scaffolding # genesis/types/validate

	// check monitoring info validity
	if  gs.MonitoringInfo != nil {
		if err := gs.MonitoringInfo.SignatureCounts.Validate(); err != nil {
			return err
		}
	}

	return gs.Params.Validate()
}
