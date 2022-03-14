package keeper_test

import (
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

	allocationPrice := types.AllocationPrice{Bonded: sdk.NewInt(100)}

	tk.ParticipationKeeper.SetParams(sdkCtx, types.Params{
		AllocationPrice: allocationPrice,
	})

	addr := sample.Address()
	_, totalShares := createNDelegations(sdkCtx, tk, addr, 10)

	accAddr, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)
	totalAlloc, err := tk.ParticipationKeeper.GetTotalAllocation(sdkCtx, accAddr)
	require.NoError(t, err)

	calcAlloc := totalShares.Quo(allocationPrice.Bonded.ToDec()).TruncateInt64()
	require.Equal(t, uint64(calcAlloc), totalAlloc)
}
