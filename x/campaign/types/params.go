package types

import (
	"errors"
	"fmt"

	"cosmossdk.io/math"

	"gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	DefaultMinTotalSupply      = sdk.NewInt(100)                   // One hundred
	DefaultMaxTotalSupply      = sdk.NewInt(1_000_000_000_000_000) // One Quadrillion
	DefaultCampaignCreationFee = sdk.Coins(nil)                    // EmptyCoins

	KeyTotalSupplyRange    = []byte("TotalSupplyRange")
	KeyCampaignCreationFee = []byte("CampaignCreationFee")
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewTotalSupplyRange creates a new TotalSupplyRange instance
func NewTotalSupplyRange(minTotalSupply, maxTotalSupply math.Int) TotalSupplyRange {
	return TotalSupplyRange{
		MinTotalSupply: minTotalSupply,
		MaxTotalSupply: maxTotalSupply,
	}
}

// NewParams creates a new Params instance
func NewParams(minTotalSupply, maxTotalSupply math.Int, campaignCreationFee sdk.Coins) Params {
	return Params{
		TotalSupplyRange:    NewTotalSupplyRange(minTotalSupply, maxTotalSupply),
		CampaignCreationFee: campaignCreationFee,
	}
}

// DefaultParams returns default campaign parameters
func DefaultParams() Params {
	return NewParams(DefaultMinTotalSupply, DefaultMaxTotalSupply, DefaultCampaignCreationFee)
}

// String implements stringer interface
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyTotalSupplyRange, &p.TotalSupplyRange, validateTotalSupplyRange),
		paramtypes.NewParamSetPair(KeyCampaignCreationFee, &p.CampaignCreationFee, validateCampaignCreationFee),
	}
}

// ValidateBasic performs basic validation on campaign parameters.
func (p Params) ValidateBasic() error {
	if err := validateTotalSupplyRange(p.TotalSupplyRange); err != nil {
		return err
	}
	return p.CampaignCreationFee.Validate()
}

func validateTotalSupplyRange(i interface{}) error {
	v, ok := i.(TotalSupplyRange)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if err := v.ValidateBasic(); err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func validateCampaignCreationFee(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return v.Validate()
}
