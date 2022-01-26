package sample

import "github.com/tendermint/spn/pkg/types"

const ConsensusStateNb = 2

// ConsensusState returns a sample ConsensusState
// nb allows to select a consensus state with a matching validator set
// consensus state 0 match with validator set 0
// nb is 0 if above max value
func ConsensusState(nb int) types.ConsensusState {
	if nb >= ConsensusStateNb {
		nb = 0
	}
	return []types.ConsensusState{
		types.NewConsensusState(
			"2022-01-12T12:25:19.523109Z",
			"48C4C20AC5A7BD99A45AEBAB92E61F5667253A2C51CCCD84D20327D3CB8737C9",
			"47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
		),
		types.NewConsensusState(
			"2022-01-12T14:15:12.981874Z",
			"65BD4CB5502F7C926228F4A929E4FAF07384B3E5A0EC553A4230B8AB5A1022ED",
			"47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
		),
	}[nb]
}

// ValidatorSet returns a sample ValidatorSet
// nb allows to select a consensus state with a matching validator set
// consensus state 0 match with validator set 0
// nb is 0 if above max value
func ValidatorSet(nb int) types.ValidatorSet {
	if nb >= ConsensusStateNb {
		nb = 0
	}
	return []types.ValidatorSet{
		types.NewValidatorSet(
			types.NewValidator("fYaox+q+N3XkGZdcQ5f3MH4/5J4oh6FRoYdW0vxRdIg=", 0, 100),
		),
		types.NewValidatorSet(
			types.NewValidator("rQMyKjkzXXUhYsAdII6fSlTkFdf24hiSPGrSCBub5Oc=", 0, 100),
		),
	}[nb]
}
