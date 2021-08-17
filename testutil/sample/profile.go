package sample

import (
	profile "github.com/tendermint/spn/x/profile/types"
)

// MsgCreateCoordinator returns a sample MsgCreateCoordinator
func MsgCreateCoordinator(coordAddress string) profile.MsgCreateCoordinator {
	return *profile.NewMsgCreateCoordinator(
		coordAddress,
		coordAddress,
		"https://cosmos.network/"+coordAddress,
		coordAddress+" details",
	)
}

// ValidatorDescription returns a sample ValidatorDescription
func ValidatorDescription(desc string) *profile.ValidatorDescription {
	return &profile.ValidatorDescription{
		Identity:        desc,
		Moniker:         "moniker " + desc,
		Website:         "https://cosmos.network/" + desc,
		SecurityContact: "foo",
		Details:         desc + " details",
	}
}

// Coordinator returns a sample Coordinator
func Coordinator() profile.Coordinator {
	return profile.Coordinator{
		Address:     AccAddress(),
		Description: CoordinatorDescription(),
	}
}

// CoordinatorDescription returns a sample CoordinatorDescription
func CoordinatorDescription() *profile.CoordinatorDescription {
	return &profile.CoordinatorDescription{
		Identity: String(10),
		Website:  String(10),
		Details:  String(30),
	}
}

// ProfileGenesisState returns a sample genesis state for the profile module
func ProfileGenesisState() profile.GenesisState {
	coordAddr1, coordAddr2 := AccAddress(), AccAddress()

	return profile.GenesisState{
		ValidatorList: []profile.Validator{
			{
				Address: AccAddress(),
				Description: ValidatorDescription(String(10)),
			},
			{
				Address: AccAddress(),
				Description: ValidatorDescription(String(10)),
			},
		},
		CoordinatorList: []profile.Coordinator{
			{
				CoordinatorId: 0,
				Address: coordAddr1,
				Description: CoordinatorDescription(),
			},
			{
				CoordinatorId: 1,
				Address: coordAddr2,
				Description: CoordinatorDescription(),
			},
		},
		CoordinatorByAddressList: []profile.CoordinatorByAddress{
			{
				Address: coordAddr1,
				CoordinatorId: 0,
			},
			{
				Address: coordAddr2,
				CoordinatorId: 1,
			},
		},
		CoordinatorCount: 2,
	}
}
