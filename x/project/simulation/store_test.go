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
	simproject "github.com/tendermint/spn/x/project/simulation"
	projecttypes "github.com/tendermint/spn/x/project/types"
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

	t.Run("should return no coordinator", func(t *testing.T) {
		_, _, found := simproject.GetCoordSimAccount(r, ctx, tk.ProfileKeeper, accs)
		require.False(t, found)
	})

	populateCoordinators(t, r, ctx, *tk.ProfileKeeper, accs, 10)

	t.Run("should find coordinators", func(t *testing.T) {
		acc, coordID, found := simproject.GetCoordSimAccount(r, ctx, tk.ProfileKeeper, accs)
		require.True(t, found)
		require.Contains(t, accs, acc)
		_, found = tk.ProfileKeeper.GetCoordinator(ctx, coordID)
		require.True(t, found)
	})
}

func TestGetCoordSimAccountWithProjectID(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	r := sample.Rand()
	accs := sample.SimAccounts()

	t.Run("should find no project", func(t *testing.T) {
		_, _, found := simproject.GetCoordSimAccountWithProjectID(
			r,
			ctx,
			tk.ProfileKeeper,
			*tk.ProjectKeeper,
			accs,
			false,
			false,
		)
		require.False(t, found)
	})

	coords := populateCoordinators(t, r, ctx, *tk.ProfileKeeper, accs, 10)

	t.Run("should find one project with mainnet launch triggered", func(t *testing.T) {
		prjt := projecttypes.NewProject(
			0,
			sample.AlphaString(r, 5),
			coords[1],
			sample.TotalSupply(r),
			sample.Metadata(r, 20),
			sample.Duration(r).Milliseconds(),
		)
		prjt.MainnetInitialized = true
		chain := sample.Chain(r, 0, coords[1])
		chain.LaunchTriggered = true
		chain.IsMainnet = true
		prjt.MainnetID = tk.LaunchKeeper.AppendChain(ctx, chain)
		tk.ProjectKeeper.AppendProject(ctx, prjt)
		_, _, found := simproject.GetCoordSimAccountWithProjectID(
			r,
			ctx,
			tk.ProfileKeeper,
			*tk.ProjectKeeper,
			accs,
			false,
			true,
		)
		require.False(t, found)
	})

	t.Run("should find a project", func(t *testing.T) {
		prjt := projecttypes.NewProject(
			1,
			sample.AlphaString(r, 5),
			coords[0],
			sample.TotalSupply(r),
			sample.Metadata(r, 20),
			sample.Duration(r).Milliseconds(),
		)
		prjt.MainnetInitialized = true
		chain := sample.Chain(r, 0, coords[1])
		chain.LaunchTriggered = false
		chain.IsMainnet = true
		prjt.MainnetID = tk.LaunchKeeper.AppendChain(ctx, chain)
		tk.ProjectKeeper.AppendProject(ctx, prjt)
		acc, id, found := simproject.GetCoordSimAccountWithProjectID(
			r,
			ctx,
			tk.ProfileKeeper,
			*tk.ProjectKeeper,
			accs,
			false,
			true,
		)
		require.True(t, found)
		require.Contains(t, accs, acc)
		_, found = tk.ProjectKeeper.GetProject(ctx, id)
		require.True(t, found)
		require.EqualValues(t, id, prjt.ProjectID)
	})

	t.Run("should find a project with no mainnet initialized", func(t *testing.T) {
		prjt := projecttypes.NewProject(
			2,
			sample.AlphaString(r, 5),
			coords[1],
			sample.TotalSupply(r),
			sample.Metadata(r, 20),
			sample.Duration(r).Milliseconds(),
		)
		idNoMainnet := tk.ProjectKeeper.AppendProject(ctx, prjt)
		acc, id, found := simproject.GetCoordSimAccountWithProjectID(
			r,
			ctx,
			tk.ProfileKeeper,
			*tk.ProjectKeeper,
			accs,
			true,
			false,
		)
		require.True(t, found)
		require.Contains(t, accs, acc)
		_, found = tk.ProjectKeeper.GetProject(ctx, id)
		require.True(t, found)
		require.EqualValues(t, idNoMainnet, id)
		require.EqualValues(t, prjt.ProjectID, id)
		require.False(t, prjt.MainnetInitialized)
	})
}

func TestGetSharesFromProject(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	r := sample.Rand()

	t.Run("should find no project", func(t *testing.T) {
		_, found := simproject.GetSharesFromProject(r, ctx, *tk.ProjectKeeper, 0)
		require.False(t, found)
	})

	t.Run("should find no shares remaining for the project", func(t *testing.T) {
		prjt := projecttypes.NewProject(
			0,
			sample.AlphaString(r, 5),
			0,
			sample.TotalSupply(r),
			sample.Metadata(r, 20),
			sample.Duration(r).Milliseconds(),
		)
		shares, err := projecttypes.NewShares(fmt.Sprintf(
			"%[1]dfoo,%[1]dbar,%[1]dtoto",
			spntypes.TotalShareNumber,
		))
		require.NoError(t, err)
		prjt.AllocatedShares = shares
		prjtSharesReached := tk.ProjectKeeper.AppendProject(ctx, prjt)
		_, found := simproject.GetSharesFromProject(r, ctx, *tk.ProjectKeeper, prjtSharesReached)
		require.False(t, found)
	})

	t.Run("should find project with available shares", func(t *testing.T) {
		prjtID := tk.ProjectKeeper.AppendProject(ctx, projecttypes.NewProject(
			1,
			sample.AlphaString(r, 5),
			0,
			sample.TotalSupply(r),
			sample.Metadata(r, 20),
			sample.Duration(r).Milliseconds(),
		))
		shares, found := simproject.GetSharesFromProject(r, ctx, *tk.ProjectKeeper, prjtID)
		require.True(t, found)
		require.NotEqualValues(t, projecttypes.EmptyShares(), shares)
	})
}

func TestGetAccountWithVouchers(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	r := sample.Rand()
	accs := sample.SimAccounts()

	mint := func(addr sdk.AccAddress, coins sdk.Coins) {
		require.NoError(t, tk.BankKeeper.MintCoins(ctx, projecttypes.ModuleName, coins))
		require.NoError(t, tk.BankKeeper.SendCoinsFromModuleToAccount(ctx, projecttypes.ModuleName, addr, coins))
	}

	t.Run("should find no account", func(t *testing.T) {
		_, _, _, found := simproject.GetAccountWithVouchers(r, ctx, tk.BankKeeper, *tk.ProjectKeeper, accs, false)
		require.False(t, found)
	})

	t.Run("should find account with vouchers for a project with launch triggered", func(t *testing.T) {
		acc, _ := simtypes.RandomAcc(r, accs)
		project := sample.Project(r, 0)
		project.MainnetInitialized = true
		chain := sample.Chain(r, 0, 0)
		chain.LaunchTriggered = true
		chain.IsMainnet = true
		project.MainnetID = tk.LaunchKeeper.AppendChain(ctx, chain)
		project.ProjectID = tk.ProjectKeeper.AppendProject(ctx, project)
		mint(acc.Address, sample.Vouchers(r, project.ProjectID))
		prjtID, acc, coins, found := simproject.GetAccountWithVouchers(r, ctx, tk.BankKeeper, *tk.ProjectKeeper, accs, false)
		require.True(t, found)
		require.EqualValues(t, project.ProjectID, prjtID)
		require.False(t, coins.Empty())
		require.Contains(t, accs, acc)
	})

	t.Run("should find account with vouchers", func(t *testing.T) {
		acc, _ := simtypes.RandomAcc(r, accs)
		project := sample.Project(r, 1)
		project.MainnetInitialized = false
		project.ProjectID = tk.ProjectKeeper.AppendProject(ctx, project)
		mint(acc.Address, sample.Vouchers(r, project.ProjectID))
		prjtID, acc, coins, found := simproject.GetAccountWithVouchers(r, ctx, tk.BankKeeper, *tk.ProjectKeeper, accs, true)
		require.True(t, found)
		require.EqualValues(t, project.ProjectID, prjtID)
		require.False(t, coins.Empty())
		require.Contains(t, accs, acc)
	})
}

func TestGetAccountWithShares(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	r := sample.Rand()
	accs := sample.SimAccounts()

	t.Run("should find no account", func(t *testing.T) {
		_, _, _, found := simproject.GetAccountWithShares(r, ctx, *tk.ProjectKeeper, accs, false)
		require.False(t, found)
	})

	t.Run("should not find account not part of sim accounts", func(t *testing.T) {
		sampleAddr := sample.Address(r)
		tk.ProjectKeeper.SetMainnetAccount(ctx, projecttypes.MainnetAccount{
			ProjectID: 10,
			Address:    sampleAddr,
			Shares:     sample.Shares(r),
		})
		_, _, _, found := simproject.GetAccountWithShares(r, ctx, *tk.ProjectKeeper, accs, false)
		require.False(t, found)
		tk.ProjectKeeper.RemoveMainnetAccount(ctx, 10, sampleAddr)
	})

	t.Run("should find account from project with launched mainnet can be retrieved", func(t *testing.T) {
		acc, _ := simtypes.RandomAcc(r, accs)
		project := sample.Project(r, 0)
		project.MainnetInitialized = true
		chain := sample.Chain(r, 0, 0)
		chain.LaunchTriggered = true
		chain.IsMainnet = true
		project.MainnetID = tk.LaunchKeeper.AppendChain(ctx, chain)
		project.ProjectID = tk.ProjectKeeper.AppendProject(ctx, project)
		share := sample.Shares(r)
		tk.ProjectKeeper.SetMainnetAccount(ctx, projecttypes.MainnetAccount{
			ProjectID: project.ProjectID,
			Address:    acc.Address.String(),
			Shares:     share,
		})
		prjtID, acc, shareRetrieved, found := simproject.GetAccountWithShares(r, ctx, *tk.ProjectKeeper, accs, false)
		require.True(t, found)
		require.Contains(t, accs, acc)
		require.EqualValues(t, project.ProjectID, prjtID)
		require.EqualValues(t, share, shareRetrieved)
	})

	t.Run("should find account from project", func(t *testing.T) {
		acc, _ := simtypes.RandomAcc(r, accs)
		project := sample.Project(r, 1)
		project.MainnetInitialized = false
		project.ProjectID = tk.ProjectKeeper.AppendProject(ctx, project)
		share := sample.Shares(r)
		tk.ProjectKeeper.SetMainnetAccount(ctx, projecttypes.MainnetAccount{
			ProjectID: project.ProjectID,
			Address:    acc.Address.String(),
			Shares:     share,
		})
		prjtID, acc, shareRetrieved, found := simproject.GetAccountWithShares(r, ctx, *tk.ProjectKeeper, accs, true)
		require.True(t, found)
		require.Contains(t, accs, acc)
		require.EqualValues(t, project.ProjectID, prjtID)
		require.EqualValues(t, share, shareRetrieved)
	})
}
