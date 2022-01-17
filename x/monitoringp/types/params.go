package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/spn/pkg/ibctypes"
	"gopkg.in/yaml.v2"
)

var (
	KeyConsumerConsensusState = []byte("ConsumerConsensusState")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(ccs *ibctypes.ConsensusState) Params {
	return Params{
		ConsumerConsensusState: ccs,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(nil)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(
			KeyConsumerConsensusState,
			&p.ConsumerConsensusState,
			validateConsumerConsensusState,
		),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return validateConsumerConsensusState(p.ConsumerConsensusState)
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateConsumerConsensusState validates consumer consensus state
func validateConsumerConsensusState(i interface{}) error {
	ccs, ok := i.(*ibctypes.ConsensusState)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// if defined, the consumer consensus state must be a valid consensus state
	if ccs != nil {
		tmConsensusState, err := ccs.ToTendermintConsensusState()
		if err != nil {
			return errors.Wrap(err, "consumer consensus state can't be converted")
		}
		if err := tmConsensusState.ValidateBasic(); err != nil {
			return errors.Wrap(err, "invalid consumer consensus state")
		}
	}
	return nil
}
