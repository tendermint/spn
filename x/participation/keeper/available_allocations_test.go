package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/participation/types"
)

func TestAvailableAllocationsGet(t *testing.T) {
	sdkCtx, tk, _ := testkeeper.NewTestSetup(t)

	invalidAddress := strconv.Itoa(1)
	allocationPrice := types.AllocationPrice{Bonded: sdk.NewInt(100)}

	tk.ParticipationKeeper.SetParams(sdkCtx, types.Params{
		AllocationPrice: allocationPrice,
	})

	validAddress := sample.Address()
	validAddressNoUse := sample.Address()
	addressNegativeDelegations := sample.Address()
	validAddressExtraUsed := sample.Address()

	tk.DelegateN(sdkCtx, validAddress, 100, 10)
	tk.DelegateN(sdkCtx, validAddressNoUse, 100, 10)
	tk.DelegateN(sdkCtx, addressNegativeDelegations, -100, 10)
	tk.DelegateN(sdkCtx, validAddressExtraUsed, 100, 10)

	tk.ParticipationKeeper.SetUsedAllocations(sdkCtx, types.UsedAllocations{
		Address:        validAddress,
		NumAllocations: 2,
	})

	// set used allocations to be greater than totalAllocations
	tk.ParticipationKeeper.SetUsedAllocations(sdkCtx, types.UsedAllocations{
		Address:        validAddressExtraUsed,
		NumAllocations: 11,
	})

	for _, tc := range []struct {
		desc       string
		address    string
		allocation uint64
		wantError  bool
	}{
		{
			desc:       "valid address with used allocations",
			address:    validAddress,
			allocation: 8, // (100 * 10 / 100) - 2 = 8
		},
		{
			desc:       "valid address with no used allocations",
			address:    validAddressNoUse,
			allocation: 10, // (100 * 10 / 100) - 0 = 10
		},
		{
			desc:       "return 0 when usedAllocations > totalAllocations",
			address:    validAddressExtraUsed,
			allocation: 0, // 11 > 10 - > return 0
		},
		{
			desc:      "invalid address returns error",
			address:   invalidAddress,
			wantError: true,
		},
		{
			desc:      "negative delegations will yield invalid allocation",
			address:   addressNegativeDelegations,
			wantError: true,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
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
