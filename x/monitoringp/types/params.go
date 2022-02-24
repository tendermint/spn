package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"

	"github.com/tendermint/spn/pkg/chainid"
	spntypes "github.com/tendermint/spn/pkg/types"
)

var (
	KeyLastBlockHeight         = []byte("LastBlockHeight")
	KeyConsumerConsensusState  = []byte("ConsumerConsensusState")
	KeyConsumerChainID         = []byte("ConsumerChainID")
	KeyConsumerUnbondingPeriod = []byte("ConsumerUnbondingPeriod")
	KeyConsumerRevisionHeight  = []byte("RevisionHeight")
	KeyDebugMode               = []byte("DebugMode")

	DefaultLastBlockHeight int64 = 1
	DefautConsumerChainID        = "spn-1"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	lastBlockHeight int64,
	consumerChainID string,
	ccs spntypes.ConsensusState,
	consumerUnbondingpPeriod int64,
	consumerRevisionHeight uint64,
	debugMode bool,
) Params {
	return Params{
		LastBlockHeight:         lastBlockHeight,
		ConsumerConsensusState:  ccs,
		ConsumerChainID:         consumerChainID,
		ConsumerUnbondingPeriod: consumerUnbondingpPeriod,
		ConsumerRevisionHeight:  consumerRevisionHeight,
		DebugMode:               debugMode,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultLastBlockHeight,
		DefautConsumerChainID,
		spntypes.ConsensusState{},
		spntypes.DefaultUnbondingPeriod,
		spntypes.DefaultRevisionHeight,
		false,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(
			KeyLastBlockHeight,
			&p.LastBlockHeight,
			validateLastBlockHeight,
		),
		paramtypes.NewParamSetPair(
			KeyConsumerConsensusState,
			&p.ConsumerConsensusState,
			validateConsumerConsensusState,
		),
		paramtypes.NewParamSetPair(
			KeyConsumerChainID,
			&p.ConsumerChainID,
			validateConsumerChainID,
		),
		paramtypes.NewParamSetPair(
			KeyConsumerUnbondingPeriod,
			&p.ConsumerUnbondingPeriod,
			validateConsumerUnbondingPeriod,
		),
		paramtypes.NewParamSetPair(
			KeyConsumerRevisionHeight,
			&p.ConsumerRevisionHeight,
			validateConsumerRevisionHeight,
		),
		paramtypes.NewParamSetPair(
			KeyDebugMode,
			&p.DebugMode,
			func(i interface{}) error { return nil },
		),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateLastBlockHeight(p.LastBlockHeight); err != nil {
		return err
	}
	if err := validateConsumerChainID(p.ConsumerChainID); err != nil {
		return err
	}
	if err := validateConsumerConsensusState(p.ConsumerConsensusState); err != nil {
		return err
	}
	if err := validateConsumerUnbondingPeriod(p.ConsumerUnbondingPeriod); err != nil {
		return err
	}
	return validateConsumerRevisionHeight(p.ConsumerRevisionHeight)
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateLastBlockHeight validates last block height
func validateLastBlockHeight(i interface{}) error {
	lastBlockHeight, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if lastBlockHeight <= 0 {
		return errors.New("last block height can't be 0 or negative")
	}

	return nil
}

// validateConsumerConsensusState validates consumer consensus state
func validateConsumerConsensusState(i interface{}) error {
	ccs, ok := i.(spntypes.ConsensusState)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// perform the verification only if the Consumer Consensus State is defined
	// TODO: remove this check and set an official SPN mainnet consensus state as default
	if ccs.Timestamp != "" {
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

// validateConsumerChainID validates consumer chain ID
func validateConsumerChainID(i interface{}) error {
	chainID, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	_, _, err := chainid.ParseGenesisChainID(chainID)
	if err != nil {
		return errors.Wrap(err, "invalid chain ID param")
	}
	return nil
}

// validateConsumerUnbondingPeriod validates consumer unbonding period
func validateConsumerUnbondingPeriod(i interface{}) error {
	unbondingPeriod, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if unbondingPeriod < spntypes.MinimalUnbondingPeriod {
		return fmt.Errorf("minimal unbonding period is %d", spntypes.MinimalUnbondingPeriod)
	}

	return nil
}

// validateConsumerRevisionHeight validates consumer revision height
func validateConsumerRevisionHeight(i interface{}) error {
	revisionHeight, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if revisionHeight == 0 {
		return fmt.Errorf("minimal unbonding period is %d", spntypes.MinimalUnbondingPeriod)
	}

	return nil
}
