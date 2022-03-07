package sample

import participation "github.com/tendermint/spn/x/participation/types"

// ParticipationParams  returns a sample of params for the participation module
func ParticipationParams() participation.Params {
	// TODO: randomize params generation
	return participation.DefaultParams()
}

// ParticipationGenesisState  returns a sample genesis state for the participation module
func ParticipationGenesisState() participation.GenesisState {
	return participation.GenesisState{
		Params:                     ParticipationParams(),
		UsedAllocationsList:        []participation.UsedAllocations{},
		AuctionUsedAllocationsList: []participation.AuctionUsedAllocations{},
	}
}
