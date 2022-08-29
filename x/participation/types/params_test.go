package types

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestParamsValidate(t *testing.T) {
	tests := []struct {
		name   string
		params Params
		err    error
	}{
		{
			name: "invalid allocation price",
			params: NewParams(
				AllocationPrice{
					Bonded: sdkmath.NewInt(-1),
				},
				DefaultParticipationTierList,
				DefaultRegistrationPeriod,
				DefaultWithdrawalDelay,
			),
			err: errors.New("value for 'bonded' must be greater than zero"),
		},
		{
			name: "invalid participation tier list",
			params: NewParams(
				DefaultAllocationPrice,
				[]Tier{
					{
						TierID:              0,
						RequiredAllocations: sdkmath.OneInt(),
						Benefits:            TierBenefits{MaxBidAmount: sdkmath.ZeroInt()},
					},
				},
				DefaultRegistrationPeriod,
				DefaultWithdrawalDelay,
			),
			err: errors.New("max bid amount must be greater than zero"),
		},
		{
			name: "invalid registration period",
			params: NewParams(
				DefaultAllocationPrice,
				DefaultParticipationTierList,
				-1,
				DefaultWithdrawalDelay,
			),
			err: errors.New("time frame must be positive"),
		},
		{
			name: "invalid withdrawal delay",
			params: NewParams(
				DefaultAllocationPrice,
				DefaultParticipationTierList,
				DefaultRegistrationPeriod,
				0,
			),
			err: errors.New("time frame must be positive"),
		},
		{
			name: "valid params",
			params: NewParams(
				DefaultAllocationPrice,
				DefaultParticipationTierList,
				DefaultRegistrationPeriod,
				DefaultWithdrawalDelay,
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.params.Validate()
			if tt.err != nil {
				require.Error(t, err, tt.err)
				require.Equal(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateAllocationPrice(t *testing.T) {
	tests := []struct {
		name            string
		allocationPrice interface{}
		err             error
	}{
		{
			name:            "invalid interface",
			allocationPrice: "test",
			err:             fmt.Errorf("invalid parameter type: string"),
		},
		{
			name:            "uninitialized bonded amount",
			allocationPrice: AllocationPrice{Bonded: sdkmath.Int{}},
			err:             errors.New("value for 'bonded' should be set"),
		},
		{
			name:            "bonded amount lower or equal than zero",
			allocationPrice: AllocationPrice{Bonded: sdkmath.ZeroInt()},
			err:             errors.New("value for 'bonded' must be greater than zero"),
		},
		{
			name:            "valid allocation price",
			allocationPrice: AllocationPrice{Bonded: sdkmath.OneInt()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateAllocationPrice(tt.allocationPrice)
			if tt.err != nil {
				require.Error(t, err, tt.err)
				require.Equal(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateParticipationTierList(t *testing.T) {
	tests := []struct {
		name                  string
		participationTierList interface{}
		err                   error
	}{
		{
			name:                  "invalid interface",
			participationTierList: "test",
			err:                   fmt.Errorf("invalid parameter type: string"),
		},
		{
			name: "duplicated tier id",
			participationTierList: []Tier{
				{
					TierID:              0,
					RequiredAllocations: sdkmath.OneInt(),
					Benefits:            TierBenefits{MaxBidAmount: sdkmath.OneInt()},
				},
				{
					TierID:              0,
					RequiredAllocations: sdkmath.NewInt(2),
					Benefits:            TierBenefits{MaxBidAmount: sdkmath.NewInt(2)},
				},
			},
			err: errors.New("duplicated tier ID: 0"),
		},
		{
			name: "invalid required allocations",
			participationTierList: []Tier{
				{
					TierID:              0,
					RequiredAllocations: sdkmath.ZeroInt(),
					Benefits:            TierBenefits{MaxBidAmount: sdkmath.OneInt()},
				},
			},
			err: errors.New("required allocations must be greater than zero"),
		},
		{
			name: "invalid tier benefits",
			participationTierList: []Tier{
				{
					TierID:              0,
					RequiredAllocations: sdkmath.OneInt(),
					Benefits:            TierBenefits{MaxBidAmount: sdkmath.ZeroInt()},
				},
			},
			err: errors.New("max bid amount must be greater than zero"),
		},
		{
			name:                  "empty participation tier list",
			participationTierList: []Tier{},
		},
		{
			name: "valid participation tier list",
			participationTierList: []Tier{
				{
					TierID:              0,
					RequiredAllocations: sdkmath.OneInt(),
					Benefits:            TierBenefits{MaxBidAmount: sdkmath.OneInt()},
				},
				{
					TierID:              1,
					RequiredAllocations: sdkmath.NewInt(2),
					Benefits:            TierBenefits{MaxBidAmount: sdkmath.NewInt(2)},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateParticipationTierList(tt.participationTierList)
			if tt.err != nil {
				require.Error(t, err, tt.err)
				require.Equal(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateTierBenefits(t *testing.T) {
	tests := []struct {
		name         string
		tierBenefits TierBenefits
		err          error
	}{
		{
			name:         "uninitialized max bid amount",
			tierBenefits: TierBenefits{MaxBidAmount: sdkmath.Int{}},
			err:          errors.New("max bid amount should be set"),
		},
		{
			name:         "max bid amount lower than zero",
			tierBenefits: TierBenefits{MaxBidAmount: sdkmath.NewInt(-1)},
			err:          errors.New("max bid amount must be greater than zero"),
		},
		{
			name:         "valid tier benefits",
			tierBenefits: TierBenefits{MaxBidAmount: sdkmath.OneInt()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateTierBenefits(tt.tierBenefits)
			if tt.err != nil {
				require.Error(t, err, tt.err)
				require.Equal(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateTimeDuration(t *testing.T) {
	tests := []struct {
		name      string
		timeFrame interface{}
		err       error
	}{
		{
			name:      "invalid interface",
			timeFrame: "test",
			err:       fmt.Errorf("invalid parameter type: string"),
		},
		{
			name:      "value not positive",
			timeFrame: time.Duration(-rand.Int63n(1000)),
			err:       errors.New("time frame must be positive"),
		},
		{
			name:      "valid time frame",
			timeFrame: time.Hour,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateTimeDuration(tt.timeFrame)
			if tt.err != nil {
				require.Error(t, err, tt.err)
				require.Equal(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
