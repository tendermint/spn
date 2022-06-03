package types

import (
	"errors"
	"fmt"
	"strings"

	yaml "gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys
var (
	KeyMintDenom               = []byte("MintDenom")
	KeyInflationRateChange     = []byte("InflationRateChange")
	KeyInflationMax            = []byte("InflationMax")
	KeyInflationMin            = []byte("InflationMin")
	KeyGoalBonded              = []byte("GoalBonded")
	KeyBlocksPerYear           = []byte("BlocksPerYear")
	KeyDistributionProportions = []byte("DistributionProportions")
)

// ParamTable for minting module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(
	mintDenom string,
	inflationRateChange,
	inflationMax,
	inflationMin,
	goalBonded sdk.Dec,
	blocksPerYear uint64,
	proportions DistributionProportions,
) Params {
	return Params{
		MintDenom:               mintDenom,
		InflationRateChange:     inflationRateChange,
		InflationMax:            inflationMax,
		InflationMin:            inflationMin,
		GoalBonded:              goalBonded,
		BlocksPerYear:           blocksPerYear,
		DistributionProportions: proportions,
	}
}

// default minting module parameters
func DefaultParams() Params {
	return Params{
		MintDenom:           sdk.DefaultBondDenom,
		InflationRateChange: sdk.NewDecWithPrec(13, 2),
		InflationMax:        sdk.NewDecWithPrec(20, 2),
		InflationMin:        sdk.NewDecWithPrec(7, 2),
		GoalBonded:          sdk.NewDecWithPrec(67, 2),
		BlocksPerYear:       uint64(60 * 60 * 8766 / 5), // assuming 5 second block times
		DistributionProportions: DistributionProportions{
			Staking:       sdk.NewDecWithPrec(4, 1), // 0.4
			Incentives:    sdk.NewDecWithPrec(4, 1), // 0.4
			CommunityPool: sdk.NewDecWithPrec(2, 1), // 0.2
		},
	}
}

// validate params
func (p Params) Validate() error {
	if err := validateMintDenom(p.MintDenom); err != nil {
		return err
	}
	if err := validateInflationRateChange(p.InflationRateChange); err != nil {
		return err
	}
	if err := validateInflationMax(p.InflationMax); err != nil {
		return err
	}
	if err := validateInflationMin(p.InflationMin); err != nil {
		return err
	}
	if err := validateGoalBonded(p.GoalBonded); err != nil {
		return err
	}
	if err := validateBlocksPerYear(p.BlocksPerYear); err != nil {
		return err
	}
	if p.InflationMax.LT(p.InflationMin) {
		return fmt.Errorf(
			"max inflation (%s) must be greater than or equal to min inflation (%s)",
			p.InflationMax, p.InflationMin,
		)
	}
	return validateDistributionProportions(p.DistributionProportions)
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMintDenom, &p.MintDenom, validateMintDenom),
		paramtypes.NewParamSetPair(KeyInflationRateChange, &p.InflationRateChange, validateInflationRateChange),
		paramtypes.NewParamSetPair(KeyInflationMax, &p.InflationMax, validateInflationMax),
		paramtypes.NewParamSetPair(KeyInflationMin, &p.InflationMin, validateInflationMin),
		paramtypes.NewParamSetPair(KeyGoalBonded, &p.GoalBonded, validateGoalBonded),
		paramtypes.NewParamSetPair(KeyBlocksPerYear, &p.BlocksPerYear, validateBlocksPerYear),
		paramtypes.NewParamSetPair(KeyDistributionProportions, &p.DistributionProportions, validateDistributionProportions),
	}
}

func validateMintDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if strings.TrimSpace(v) == "" {
		return errors.New("mint denom cannot be blank")
	}
	if err := sdk.ValidateDenom(v); err != nil {
		return err
	}

	return nil
}

func validateInflationRateChange(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("inflation rate change cannot be negative: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("inflation rate change too large: %s", v)
	}

	return nil
}

func validateInflationMax(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("max inflation cannot be negative: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("max inflation too large: %s", v)
	}

	return nil
}

func validateInflationMin(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("min inflation cannot be negative: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("min inflation too large: %s", v)
	}

	return nil
}

func validateGoalBonded(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("goal bonded cannot be negative: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("goal bonded too large: %s", v)
	}

	return nil
}

func validateBlocksPerYear(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("blocks per year must be positive: %d", v)
	}

	return nil
}

func validateDistributionProportions(i interface{}) error {
	v, ok := i.(DistributionProportions)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.Staking.IsNegative() {
		return errors.New("staking distribution ratio should not be negative")
	}

	if v.Incentives.IsNegative() {
		return errors.New("pool incentives distribution ratio should not be negative")
	}

	if v.CommunityPool.IsNegative() {
		return errors.New("community pool distribution ratio should not be negative")
	}

	totalProportions := v.Staking.Add(v.Incentives).Add(v.CommunityPool)

	if !totalProportions.Equal(sdk.NewDec(1)) {
		return errors.New("total distributions ratio should be 1")
	}

	return nil
}
