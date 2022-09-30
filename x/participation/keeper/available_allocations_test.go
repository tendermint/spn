package keeper_test

import (
	"strconv"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/participation/types"
)

func TestAvailableAllocationsGet(t *testing.T) {
	sdkCtx, tk, _ := testkeeper.NewTestSetup(t)

	invalidAddress := strconv.Itoa(1)
	allocationPrice := types.AllocationPrice{Bonded: sdkmath.NewInt(100)}

	params := types.DefaultParams()
	params.AllocationPrice = allocationPrice
	tk.ParticipationKeeper.SetParams(sdkCtx, params)

	validAddress := sample.Address(r)
	validAddressNoUse := sample.Address(r)
	addressNegativeDelegations := sample.Address(r)
	validAddressExtraUsed := sample.Address(r)

	tk.DelegateN(sdkCtx, r, validAddress, 100, 10)
	tk.DelegateN(sdkCtx, r, validAddressNoUse, 100, 10)
	tk.DelegateN(sdkCtx, r, addressNegativeDelegations, -100, 10)
	tk.DelegateN(sdkCtx, r, validAddressExtraUsed, 100, 10)

	tk.ParticipationKeeper.SetUsedAllocations(sdkCtx, types.UsedAllocations{
		Address:        validAddress,
		NumAllocations: sdkmath.NewInt(2),
	})

	// set used allocations to be greater than totalAllocations
	tk.ParticipationKeeper.SetUsedAllocations(sdkCtx, types.UsedAllocations{
		Address:        validAddressExtraUsed,
		NumAllocations: sdkmath.NewInt(11),
	})

	for _, tc := range []struct {
		name       string
		address    string
		allocation sdkmath.Int
		wantError  bool
	}{
		{
			name:       "should allow with valid address with used allocations",
			address:    validAddress,
			allocation: sdkmath.NewInt(8), // (100 * 10 / 100) - 2 = 8
		},
		{
			name:       "should allow with valid address with no used allocations",
			address:    validAddressNoUse,
			allocation: sdkmath.NewInt(10), // (100 * 10 / 100) - 0 = 10
		},
		{
			name:       "should return 0 when usedAllocations > totalAllocations",
			address:    validAddressExtraUsed,
			allocation: sdkmath.ZeroInt(), // 11 > 10 - > return 0
		},
		{
			name:      "should prevent with invalid address",
			address:   invalidAddress,
			wantError: true,
		},
		{
			name:      "should prevent with negative delegations",
			address:   addressNegativeDelegations,
			wantError: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			alloc, err := tk.ParticipationKeeper.GetAvailableAllocations(sdkCtx, tc.address)
			if tc.wantError {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.allocation, alloc)
			}
		})
	}
}
