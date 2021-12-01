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
func ProfileGenesisState(addresses ...string) profile.GenesisState {
	for len(addresses) < 7 {
		addresses = append(addresses, Address())
	}
	return profile.GenesisState{
		CoordinatorList: []profile.Coordinator{
			{
				CoordinatorID: 0,
				Address:       addresses[0],
				Description:   CoordinatorDescription(),
			},
			{
				CoordinatorID: 1,
				Address:       addresses[1],
				Description:   CoordinatorDescription(),
			},
			{
				CoordinatorID: 2,
				Address:       addresses[2],
				Description:   CoordinatorDescription(),
			},
			{
				CoordinatorID: 3,
				Address:       addresses[3],
				Description:   CoordinatorDescription(),
			},
			{
				CoordinatorID: 4,
				Address:       addresses[4],
				Description:   CoordinatorDescription(),
			},
		},
		CoordinatorByAddressList: []profile.CoordinatorByAddress{
			{
				Address:       addresses[0],
				CoordinatorID: 0,
			},
			{
				Address:       addresses[1],
				CoordinatorID: 1,
			},
			{
				Address:       addresses[2],
				CoordinatorID: 2,
			},
			{
				Address:       addresses[3],
				CoordinatorID: 3,
			},
			{
				Address:       addresses[4],
				CoordinatorID: 4,
			},
		},
		CoordinatorCounter: 5,
		ValidatorList: []profile.Validator{
			{
				Address:     addresses[5],
				Description: ValidatorDescription(String(10)),
			},
			{
				Address:     addresses[6],
				Description: ValidatorDescription(String(10)),
			},
		},
	}
}
