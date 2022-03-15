package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/participation/types"
)

func createNDelegations(ctx sdk.Context, tk testkeeper.TestKeepers, addr string, n int) ([]stakingtypes.Delegation, sdk.Dec) {
	items := make([]stakingtypes.Delegation, n)

	totalShares := sdk.ZeroDec()
	for i := range items {
		items[i] = sample.Delegation(tk.T, addr)
		totalShares = totalShares.Add(items[i].Shares)
		tk.StakingKeeper.SetDelegation(ctx, items[i])
	}

	return items, totalShares
}

func TestTotalAllcationGet(t *testing.T) {
	sdkCtx, tk, _ := testkeeper.NewTestSetup(t)

	// expect error with invalid address
	invalidAddr := strconv.Itoa(1)
	_, err := tk.ParticipationKeeper.GetTotalAllocation(sdkCtx, invalidAddr)
	// check error strings since bech32 errors are not typed
	require.Contains(t, err.Error(), "decoding bech32 failed: invalid bech32 string length 1")

	allocationPrice := types.AllocationPrice{Bonded: sdk.NewInt(100)}

	tk.ParticipationKeeper.SetParams(sdkCtx, types.Params{
		AllocationPrice: allocationPrice,
	})

	addr := sample.Address()
	_, totalShares := createNDelegations(sdkCtx, tk, addr, 10)

	totalAlloc, err := tk.ParticipationKeeper.GetTotalAllocation(sdkCtx, addr)
	require.NoError(t, err)

	calcAlloc := totalShares.Quo(allocationPrice.Bonded.ToDec()).TruncateInt64()
	require.Equal(t, uint64(calcAlloc), totalAlloc)
}
