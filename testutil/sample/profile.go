package sample

import profile "github.com/tendermint/spn/x/profile/types"

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
		Description: &profile.CoordinatorDescription{
			Identity: String(10),
			Website:  String(10),
			Details:  String(30),
		},
	}
}
