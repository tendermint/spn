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
func ValidatorDescription(desc string) profile.ValidatorDescription {
	return profile.ValidatorDescription{
		Identity:        desc,
		Moniker:         "moniker " + desc,
		Website:         "https://cosmos.network/" + desc,
		SecurityContact: "foo",
		Details:         desc + " details",
	}
}

// Coordinator returns a sample Coordinator
func Coordinator(address string) profile.Coordinator {
	return profile.Coordinator{
		Address:     address,
		Description: CoordinatorDescription(),
	}
}

// CoordinatorDescription returns a sample CoordinatorDescription
func CoordinatorDescription() profile.CoordinatorDescription {
	return profile.CoordinatorDescription{
		Identity: String(10),
		Website:  String(10),
		Details:  String(30),
	}
}

// ProfileGenesisState returns a sample genesis state for the profile module
func ProfileGenesisState(coordinators ...string) profile.GenesisState {
	for len(coordinators) < 3 {
		coordinators = append(coordinators, Address())
	}
	return profile.GenesisState{
		ValidatorList: []profile.Validator{
			{
				Address:     Address(),
				Description: ValidatorDescription(String(10)),
			},
			{
				Address:     Address(),
				Description: ValidatorDescription(String(10)),
			},
		},
		CoordinatorList: []profile.Coordinator{
			{
				CoordinatorId: 0,
				Address:       coordinators[0],
				Description:   CoordinatorDescription(),
			},
			{
				CoordinatorId: 1,
				Address:       coordinators[1],
				Description:   CoordinatorDescription(),
			},
			{
				CoordinatorId: 2,
				Address:       coordinators[2],
				Description:   CoordinatorDescription(),
			},
		},
		CoordinatorByAddressList: []profile.CoordinatorByAddress{
			{
				Address:       coordinators[0],
				CoordinatorId: 0,
			},
			{
				Address:       coordinators[1],
				CoordinatorId: 1,
			},
			{
				Address:       coordinators[2],
				CoordinatorId: 1,
			},
		},
		CoordinatorCount: 3,
	}
}
