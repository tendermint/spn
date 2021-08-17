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
		Address: AccAddress(),
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
				AccAddress(),
				ValidatorDescription(String(10)),
			},
			{
				AccAddress(),
				ValidatorDescription(String(10)),
			},
		},
		CoordinatorList: []profile.Coordinator{
			{
				0,
				coordAddr1,
				CoordinatorDescription(),
			},
			{
				1,
				coordAddr2,
				CoordinatorDescription(),
			},
		},
		CoordinatorByAddressList: []profile.CoordinatorByAddress{
			{
				coordAddr1,
				0,
			},
			{
				coordAddr2,
				1,
			},
		},
		CoordinatorCount: 2,
	}
}