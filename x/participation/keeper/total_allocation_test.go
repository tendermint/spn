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

func TestTotalAllocationGet(t *testing.T) {
	sdkCtx, tk, _ := testkeeper.NewTestSetup(t)

	invalidAddress := strconv.Itoa(1)
	allocationPrice := types.AllocationPrice{Bonded: sdk.NewInt(100)}

	tk.ParticipationKeeper.SetParams(sdkCtx, types.Params{
		AllocationPrice: allocationPrice,
	})

	validAddress := sample.Address()
	addressNegativeDelegations := sample.Address()

	tk.DelegateN(sdkCtx, validAddress, 100, 10)
	tk.DelegateN(sdkCtx, addressNegativeDelegations, -100, 10)

	for _, tc := range []struct {
		desc       string
		address    string
		allocation uint64
		wantError  bool
	}{
		{
			desc:       "valid address",
			address:    validAddress,
			allocation: 10, // 100 * 10 / 100 = 10
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
			alloc, err := tk.ParticipationKeeper.GetTotalAllocation(sdkCtx, tc.address)
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
