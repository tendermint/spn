package keeper

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/x/campaign/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNMainnetVestingAccount(keeper *Keeper, ctx sdk.Context, n int) []types.MainnetVestingAccount {
	items := make([]types.MainnetVestingAccount, n)
	for i := range items {
		items[i].CampaignID = uint64(i)
		items[i].Address = strconv.Itoa(i)

		keeper.SetMainnetVestingAccount(ctx, items[i])
	}
	return items
}

func TestMainnetVestingAccountGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNMainnetVestingAccount(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMainnetVestingAccount(ctx,
			item.CampaignID,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestMainnetVestingAccountRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNMainnetVestingAccount(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMainnetVestingAccount(ctx,
			item.CampaignID,
			item.Address,
		)
		_, found := keeper.GetMainnetVestingAccount(ctx,
			item.CampaignID,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestMainnetVestingAccountGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNMainnetVestingAccount(keeper, ctx, 10)
	require.Equal(t, items, keeper.GetAllMainnetVestingAccount(ctx))
}
