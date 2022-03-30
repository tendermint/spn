package simulation_test

import (
	"fmt"
	"math/rand"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stretchr/testify/require"

	spntypes "github.com/tendermint/spn/pkg/types"
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
			Description: sample.CoordinatorDescription(r),
			Active:      true,
		})

		coordIDs = append(coordIDs, coordID)
	}

	return
}

func TestGetCoordSimAccount(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	r := sample.Rand()
	accs := sample.SimAccounts()

	t.Run("no coordinator", func(t *testing.T) {
		_, _, found := simcampaign.GetCoordSimAccount(r, ctx, tk.ProfileKeeper, accs)
		require.False(t, found)
	})

	populateCoordinators(t, r, ctx, *tk.ProfileKeeper, accs, 10)

	t.Run("find coordinators", func(t *testing.T) {
		acc, coordID, found := simcampaign.GetCoordSimAccount(r, ctx, tk.ProfileKeeper, accs)
		require.True(t, found)
		require.Contains(t, accs, acc)
		_, found = tk.ProfileKeeper.GetCoordinator(ctx, coordID)
		require.True(t, found)
	})
}

func TestGetCoordSimAccountWithCampaignID(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	r := sample.Rand()
	accs := sample.SimAccounts()

	t.Run("no campaign", func(t *testing.T) {
		_, _, found := simcampaign.GetCoordSimAccountWithCampaignID(
			r,
			ctx,
			tk.ProfileKeeper,
			*tk.CampaignKeeper,
			accs,
			false,
		)
		require.False(t, found)
	})

	coords := populateCoordinators(t, r, ctx, *tk.ProfileKeeper, accs, 10)

	t.Run("find a campaign", func(t *testing.T) {
		camp := campaigntypes.NewCampaign(
			0,
			sample.AlphaString(r, 5),
			coords[0],
			sample.TotalSupply(r),
			sample.Metadata(r, 20),
		)
		camp.MainnetInitialized = true
		tk.CampaignKeeper.AppendCampaign(ctx, camp)
		acc, id, found := simcampaign.GetCoordSimAccountWithCampaignID(
			r,
			ctx,
			tk.ProfileKeeper,
			*tk.CampaignKeeper,
			accs,
			false,
		)
		require.True(t, found)
		require.Contains(t, accs, acc)
		_, found = tk.CampaignKeeper.GetCampaign(ctx, id)
		require.True(t, found)
		require.EqualValues(t, id, camp.CampaignID)
	})

	t.Run("find a campaign with no mainnet initialized", func(t *testing.T) {
		camp := campaigntypes.NewCampaign(
			1,
			sample.AlphaString(r, 5),
			coords[1],
			sample.TotalSupply(r),
			sample.Metadata(r, 20),
		)
		camp.MainnetInitialized = true
		idDynamicShares := tk.CampaignKeeper.AppendCampaign(ctx, camp)
		acc, id, found := simcampaign.GetCoordSimAccountWithCampaignID(
			r,
			ctx,
			tk.ProfileKeeper,
			*tk.CampaignKeeper,
			accs,
			false,
		)
		require.True(t, found)
		require.Contains(t, accs, acc)
		camp, found = tk.CampaignKeeper.GetCampaign(ctx, id)
		require.True(t, found)
		require.EqualValues(t, idDynamicShares, id)
		require.EqualValues(t, id, camp.CampaignID)
	})

	t.Run("find a campaign with no mainnet initialized", func(t *testing.T) {
		camp := campaigntypes.NewCampaign(
			2,
			sample.AlphaString(r, 5),
			coords[2],
			sample.TotalSupply(r),
			sample.Metadata(r, 20),
		)
		idNoMainnet := tk.CampaignKeeper.AppendCampaign(ctx, camp)
		acc, id, found := simcampaign.GetCoordSimAccountWithCampaignID(
			r,
			ctx,
			tk.ProfileKeeper,
			*tk.CampaignKeeper,
			accs,
			true,
		)
		require.True(t, found)
		require.Contains(t, accs, acc)
		_, found = tk.CampaignKeeper.GetCampaign(ctx, id)
		require.True(t, found)
		require.EqualValues(t, idNoMainnet, id)
		require.EqualValues(t, camp.CampaignID, id)
		require.False(t, camp.MainnetInitialized)
	})
}

func TestGetSharesFromCampaign(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	r := sample.Rand()

	t.Run("no campaign", func(t *testing.T) {
		_, found := simcampaign.GetSharesFromCampaign(r, ctx, *tk.CampaignKeeper, 0)
		require.False(t, found)
	})

	t.Run("no shares remains for the campaign", func(t *testing.T) {
		camp := campaigntypes.NewCampaign(
			0,
			sample.AlphaString(r, 5),
			0,
			sample.TotalSupply(r),
			sample.Metadata(r, 20),
		)
		shares, err := campaigntypes.NewShares(fmt.Sprintf(
			"%[1]dfoo,%[1]dbar,%[1]dtoto",
			spntypes.TotalShareNumber,
		))
		require.NoError(t, err)
		camp.AllocatedShares = shares
		campSharesReached := tk.CampaignKeeper.AppendCampaign(ctx, camp)
		_, found := simcampaign.GetSharesFromCampaign(r, ctx, *tk.CampaignKeeper, campSharesReached)
		require.False(t, found)
	})

	t.Run("campaign with available shares", func(t *testing.T) {
		campID := tk.CampaignKeeper.AppendCampaign(ctx, campaigntypes.NewCampaign(
			1,
			sample.AlphaString(r, 5),
			0,
			sample.TotalSupply(r),
			sample.Metadata(r, 20),
		))
		shares, found := simcampaign.GetSharesFromCampaign(r, ctx, *tk.CampaignKeeper, campID)
		require.True(t, found)
		require.NotEqualValues(t, campaigntypes.EmptyShares(), shares)
	})
}

func TestGetAccountWithVouchers(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	r := sample.Rand()
	accs := sample.SimAccounts()

	mint := func(addr sdk.AccAddress, coins sdk.Coins) {
		require.NoError(t, tk.BankKeeper.MintCoins(ctx, campaigntypes.ModuleName, coins))
		require.NoError(t, tk.BankKeeper.SendCoinsFromModuleToAccount(ctx, campaigntypes.ModuleName, addr, coins))
	}

	t.Run("no account", func(t *testing.T) {
		_, _, _, found := simcampaign.GetAccountWithVouchers(ctx, tk.BankKeeper, accs)
		require.False(t, found)
	})

	t.Run("vouchers from an account", func(t *testing.T) {
		acc, _ := simtypes.RandomAcc(r, accs)
		mint(acc.Address, sample.Vouchers(r, 10))
		campID, acc, coins, found := simcampaign.GetAccountWithVouchers(ctx, tk.BankKeeper, accs)
		require.True(t, found)
		require.EqualValues(t, 10, campID)
		require.False(t, coins.Empty())
		require.Contains(t, accs, acc)
	})
}

func TestGetAccountWithShares(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	r := sample.Rand()
	accs := sample.SimAccounts()

	t.Run("no account", func(t *testing.T) {
		_, _, _, found := simcampaign.GetAccountWithShares(r, ctx, *tk.CampaignKeeper, accs)
		require.False(t, found)
	})

	t.Run("account not part of sim accounts", func(t *testing.T) {
		sampleAddr := sample.Address(r)
		tk.CampaignKeeper.SetMainnetAccount(ctx, campaigntypes.MainnetAccount{
			CampaignID: 10,
			Address:    sampleAddr,
			Shares:     sample.Shares(r),
		})
		_, _, _, found := simcampaign.GetAccountWithShares(r, ctx, *tk.CampaignKeeper, accs)
		require.False(t, found)
		tk.CampaignKeeper.RemoveMainnetAccount(ctx, 10, sampleAddr)
	})

	t.Run("account can be retrieved", func(t *testing.T) {
		acc, _ := simtypes.RandomAcc(r, accs)
		share := sample.Shares(r)
		tk.CampaignKeeper.SetMainnetAccount(ctx, campaigntypes.MainnetAccount{
			CampaignID: 10,
			Address:    acc.Address.String(),
			Shares:     share,
		})
		campID, acc, shareRetrieved, found := simcampaign.GetAccountWithShares(r, ctx, *tk.CampaignKeeper, accs)
		require.True(t, found)
		require.Contains(t, accs, acc)
		require.EqualValues(t, uint64(10), campID)
		require.EqualValues(t, share, shareRetrieved)
	})
}
