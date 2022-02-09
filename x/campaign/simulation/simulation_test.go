package simulation_test

import (
	"fmt"
	"math/rand"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	simcampaign "github.com/tendermint/spn/x/campaign/simulation"
	campaigntypes "github.com/tendermint/spn/x/campaign/types"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

// populateCoordinators populates the profile keeper with some coordinators from simulation accounts
func populateCoordinators(
	t *testing.T,
	r *rand.Rand,
	ctx sdk.Context,
	pk profilekeeper.Keeper,
	accs []simtypes.Account,
	coordNb int,
) (coordIDs []uint64) {
	require.LessOrEqual(t, coordNb, len(accs))
	r.Shuffle(len(accs), func(i, j int) {
		accs[i], accs[j] = accs[j], accs[i]
	})
	for i := 0; i < coordNb; i++ {
		coordID := pk.AppendCoordinator(ctx, profiletypes.Coordinator{
			Address:     accs[i].Address.String(),
			Description: sample.CoordinatorDescription(),
		})

		coordIDs = append(coordIDs, coordID)
	}

	return
}

func TestGetCoordSimAccount(t *testing.T) {
	pk, ctx := testkeeper.Profile(t)
	r := sample.Rand()
	accs := sample.SimAccounts()

	t.Run("no coordinator", func(t *testing.T) {
		_, _, found := simcampaign.GetCoordSimAccount(r, ctx, pk, accs)
		require.False(t, found)
	})

	populateCoordinators(t, r, ctx, *pk, accs, 10)

	t.Run("find coordinators", func(t *testing.T) {
		acc, coordID, found := simcampaign.GetCoordSimAccount(r, ctx, pk, accs)
		require.True(t, found)
		require.Contains(t, accs, acc)
		_, found = pk.GetCoordinator(ctx, coordID)
		require.True(t, found)
	})
}

func TestGetCoordSimAccountWithCampaignID(t *testing.T) {
	ck, _, pk, _, _, _, _, ctx := testkeeper.AllKeepers(t)
	r := sample.Rand()
	accs := sample.SimAccounts()

	t.Run("no campaign", func(t *testing.T) {
		_, _, found := simcampaign.GetCoordSimAccountWithCampaignID(r, ctx, pk, *ck, accs, false, false)
		require.False(t, found)
	})

	coords := populateCoordinators(t, r, ctx, *pk, accs, 10)

	t.Run("find a campaign", func(t *testing.T) {
		camp := campaigntypes.NewCampaign(
			0,
			sample.AlphaString(5),
			coords[0],
			sample.Coins(),
			false,
		)
		camp.MainnetInitialized = true
		ck.AppendCampaign(ctx, camp)
		acc, id, found := simcampaign.GetCoordSimAccountWithCampaignID(r, ctx, pk, *ck, accs, false, false)
		require.True(t, found)
		require.Contains(t, accs, acc)
		_, found = ck.GetCampaign(ctx, id)
		require.True(t, found)
		require.EqualValues(t, id, camp.CampaignID)
	})

	t.Run("find a campaign with dynamic shares", func(t *testing.T) {
		camp := campaigntypes.NewCampaign(
			1,
			sample.AlphaString(5),
			coords[1],
			sample.Coins(),
			true,
		)
		camp.MainnetInitialized = true
		idDynamicShares := ck.AppendCampaign(ctx, camp)
		acc, id, found := simcampaign.GetCoordSimAccountWithCampaignID(r, ctx, pk, *ck, accs, true, false)
		require.True(t, found)
		require.Contains(t, accs, acc)
		camp, found = ck.GetCampaign(ctx, id)
		require.True(t, found)
		require.EqualValues(t, idDynamicShares, id)
		require.EqualValues(t, id, camp.CampaignID)
		require.True(t, camp.DynamicShares)
	})

	t.Run("find a campaign with no mainnet initialized", func(t *testing.T) {
		camp := campaigntypes.NewCampaign(
			2,
			sample.AlphaString(5),
			coords[2],
			sample.Coins(),
			false,
		)
		idNoMainnet := ck.AppendCampaign(ctx, camp)
		acc, id, found := simcampaign.GetCoordSimAccountWithCampaignID(r, ctx, pk, *ck, accs, false, true)
		require.True(t, found)
		require.Contains(t, accs, acc)
		_, found = ck.GetCampaign(ctx, id)
		require.True(t, found)
		require.EqualValues(t, idNoMainnet, id)
		require.EqualValues(t, id, camp.CampaignID)
		require.False(t, camp.MainnetInitialized)
	})

	t.Run("find a campaign with no mainnet initialized and with dynamic shares", func(t *testing.T) {
		camp := campaigntypes.NewCampaign(
			3,
			sample.AlphaString(5),
			coords[3],
			sample.Coins(),
			true,
		)
		idNoMainnetDynamicShares := ck.AppendCampaign(ctx, camp)
		acc, id, found := simcampaign.GetCoordSimAccountWithCampaignID(r, ctx, pk, *ck, accs, true, true)
		require.True(t, found)
		require.Contains(t, accs, acc)
		_, found = ck.GetCampaign(ctx, id)
		require.True(t, found)
		require.EqualValues(t, idNoMainnetDynamicShares, id)
		require.EqualValues(t, id, camp.CampaignID)
		require.False(t, camp.MainnetInitialized)
		require.True(t, camp.DynamicShares)
	})
}

func TestGetSharesFromCampaign(t *testing.T) {
	ck, ctx := testkeeper.Campaign(t)
	r := sample.Rand()

	t.Run("no campaign", func(t *testing.T) {
		_, found := simcampaign.GetSharesFromCampaign(r, ctx, *ck, 0)
		require.False(t, found)
	})

	t.Run("no shares remains for the campaign", func(t *testing.T) {
		camp := campaigntypes.NewCampaign(
			0,
			sample.AlphaString(5),
			0,
			sample.Coins(),
			false,
		)
		shares, err := campaigntypes.NewShares(fmt.Sprintf(
			"%[1]dfoo,%[1]dbar,%[1]dtoto",
			campaigntypes.DefaultTotalShareNumber,
		))
		require.NoError(t, err)
		camp.AllocatedShares = shares
		campSharesReached := ck.AppendCampaign(ctx, camp)
		_, found := simcampaign.GetSharesFromCampaign(r, ctx, *ck, campSharesReached)
		require.False(t, found)
	})

	t.Run("campaign with available shares", func(t *testing.T) {
		campID := ck.AppendCampaign(ctx, campaigntypes.NewCampaign(
			1,
			sample.AlphaString(5),
			0,
			sample.Coins(),
			false,
		))
		shares, found := simcampaign.GetSharesFromCampaign(r, ctx, *ck, campID)
		require.True(t, found)
		require.NotEqualValues(t, campaigntypes.EmptyShares(), shares)
	})
}

func TestGetAccountWithVouchers(t *testing.T) {
	_, _, _, _, _, bk, _, ctx := testkeeper.AllKeepers(t)
	r := sample.Rand()
	accs := sample.SimAccounts()

	mint := func(addr sdk.AccAddress, coins sdk.Coins) {
		require.NoError(t, bk.MintCoins(ctx, campaigntypes.ModuleName, coins))
		require.NoError(t, bk.SendCoinsFromModuleToAccount(ctx, campaigntypes.ModuleName, addr, coins))
	}

	t.Run("no account", func(t *testing.T) {
		_, _, _, found := simcampaign.GetAccountWithVouchers(ctx, bk, accs)
		require.False(t, found)
	})

	t.Run("vouchers from an account", func(t *testing.T) {
		acc, _ := simtypes.RandomAcc(r, accs)
		mint(acc.Address, sample.Vouchers(10))
		campID, acc, coins, found := simcampaign.GetAccountWithVouchers(ctx, bk, accs)
		require.True(t, found)
		require.EqualValues(t, 10, campID)
		require.False(t, coins.Empty())
		require.Contains(t, accs, acc)
	})
}

func TestGetAccountWithShares(t *testing.T) {
	ck, ctx := testkeeper.Campaign(t)
	r := sample.Rand()
	accs := sample.SimAccounts()

	t.Run("no account", func(t *testing.T) {
		_, _, _, found := simcampaign.GetAccountWithShares(r, ctx, *ck, accs)
		require.False(t, found)
	})

	t.Run("account not part of sim accounts", func(t *testing.T) {
		sampleAddr := sample.Address()
		ck.SetMainnetAccount(ctx, campaigntypes.MainnetAccount{
			CampaignID: 10,
			Address:    sampleAddr,
			Shares:     sample.Shares(),
		})
		_, _, _, found := simcampaign.GetAccountWithShares(r, ctx, *ck, accs)
		require.False(t, found)
		ck.RemoveMainnetAccount(ctx, 10, sampleAddr)
	})

	t.Run("account can be retrieved", func(t *testing.T) {
		acc, _ := simtypes.RandomAcc(r, accs)
		share := sample.Shares()
		ck.SetMainnetAccount(ctx, campaigntypes.MainnetAccount{
			CampaignID: 10,
			Address:    acc.Address.String(),
			Shares:     share,
		})
		campID, acc, shareRetrieved, found := simcampaign.GetAccountWithShares(r, ctx, *ck, accs)
		require.True(t, found)
		require.Contains(t, accs, acc)
		require.EqualValues(t, uint64(10), campID)
		require.EqualValues(t, share, shareRetrieved)
	})
}
