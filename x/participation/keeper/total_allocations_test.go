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

func TestTotalAllocationsGet(t *testing.T) {
	sdkCtx, tk, _ := testkeeper.NewTestSetup(t)

	invalidAddress := strconv.Itoa(1)
	params := types.DefaultParams()
	params.AllocationPrice = types.AllocationPrice{Bonded: sdk.NewInt(100)}

	tk.ParticipationKeeper.SetParams(sdkCtx, params)

	validAddress := sample.Address(r)
	addressNegativeDelegations := sample.Address(r)

	tk.DelegateN(sdkCtx, r, validAddress, 100, 10)
	tk.DelegateN(sdkCtx, r, addressNegativeDelegations, -100, 10)

	for _, tc := range []struct {
		desc       string
		address    string
		allocation sdk.Int
		wantError  bool
	}{
		{
			desc:       "valid address",
			address:    validAddress,
			allocation: sdk.NewInt(10), // 100 * 10 / 100 = 10
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
			alloc, err := tk.ParticipationKeeper.GetTotalAllocations(sdkCtx, tc.address)
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
