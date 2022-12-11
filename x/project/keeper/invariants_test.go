package keeper_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	tc "github.com/tendermint/spn/testutil/constructor"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/project/keeper"
	"github.com/tendermint/spn/x/project/types"
)

func TestAccountWithoutProjectInvariant(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("should allow valid case", func(t *testing.T) {
		project := sample.Project(r, 0)
		project.ProjectID = tk.ProjectKeeper.AppendProject(ctx, project)
		tk.ProjectKeeper.SetMainnetAccount(ctx, sample.MainnetAccount(r, project.ProjectID, sample.Address(r)))
		msg, broken := keeper.AccountWithoutProjectInvariant(*tk.ProjectKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("should prevent invalid case", func(t *testing.T) {
		tk.ProjectKeeper.SetMainnetAccount(ctx, sample.MainnetAccount(r, 100, sample.Address(r)))
		msg, broken := keeper.AccountWithoutProjectInvariant(*tk.ProjectKeeper)(ctx)
		require.True(t, broken, msg)
	})
}

func TestProjectSharesInvariant(t *testing.T) {
	t.Run("should allow valid case", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		// create projects with some allocated shares
		projectID1, projectID2 := uint64(1), uint64(2)
		project := sample.Project(r, projectID1)
		project.AllocatedShares = types.IncreaseShares(
			project.AllocatedShares,
			tc.Shares(t, "100foo,200bar"),
		)
		tk.ProjectKeeper.SetProject(ctx, project)

		project = sample.Project(r, projectID2)
		project.AllocatedShares = types.IncreaseShares(
			project.AllocatedShares,
			tc.Shares(t, "10000foo"),
		)
		tk.ProjectKeeper.SetProject(ctx, project)

		// mint vouchers
		voucherFoo, voucherBar := types.VoucherDenom(projectID1, "foo"), types.VoucherDenom(projectID1, "bar")
		tk.Mint(ctx, sample.Address(r), tc.Coins(t, fmt.Sprintf("50%s,100%s", voucherFoo, voucherBar)))

		// mint vouchers for another project
		voucherFoo = types.VoucherDenom(projectID2, "foo")
		tk.Mint(ctx, sample.Address(r), tc.Coins(t, fmt.Sprintf("5000%s", voucherFoo)))

		// add accounts with shares
		tk.ProjectKeeper.SetMainnetAccount(ctx, types.MainnetAccount{
			ProjectID: projectID1,
			Address:   sample.Address(r),
			Shares:    tc.Shares(t, "20foo,40bar"),
		})
		tk.ProjectKeeper.SetMainnetAccount(ctx, types.MainnetAccount{
			ProjectID: projectID1,
			Address:   sample.Address(r),
			Shares:    tc.Shares(t, "30foo,60bar"),
		})
		tk.ProjectKeeper.SetMainnetAccount(ctx, types.MainnetAccount{
			ProjectID: projectID2,
			Address:   sample.Address(r),
			Shares:    tc.Shares(t, "5000foo"),
		})

		msg, broken := keeper.ProjectSharesInvariant(*tk.ProjectKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("should allow project with empty allocated share is valid", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		tk.ProjectKeeper.SetProject(ctx, sample.Project(r, 3))

		msg, broken := keeper.ProjectSharesInvariant(*tk.ProjectKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("should prevent allocated shares cannot be converted to vouchers", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		projectID := uint64(4)
		project := sample.Project(r, projectID)
		coins := tc.Coins(t, "100foo,200bar")
		shares := make(types.Shares, len(coins))
		for i, coin := range coins {
			shares[i] = coin
		}
		project.AllocatedShares = types.IncreaseShares(
			project.AllocatedShares,
			shares,
		)
		tk.ProjectKeeper.SetProject(ctx, project)

		msg, broken := keeper.ProjectSharesInvariant(*tk.ProjectKeeper)(ctx)
		require.True(t, broken, msg)
	})

	t.Run("should prevent invalid allocated shares", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		projectID := uint64(4)
		project := sample.Project(r, projectID)
		project.AllocatedShares = types.IncreaseShares(
			project.AllocatedShares,
			tc.Shares(t, "100foo,200bar"),
		)
		tk.ProjectKeeper.SetProject(ctx, project)

		// mint vouchers
		voucherFoo, voucherBar := types.VoucherDenom(projectID, "foo"), types.VoucherDenom(projectID, "bar")
		tk.Mint(ctx, sample.Address(r), tc.Coins(t, fmt.Sprintf("99%s,200%s", voucherFoo, voucherBar)))

		msg, broken := keeper.ProjectSharesInvariant(*tk.ProjectKeeper)(ctx)
		require.True(t, broken, msg)
	})

	t.Run("should prevent project with special allocations not tracked by allocated shares", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		project := sample.Project(r, 3)
		project.SpecialAllocations.GenesisDistribution = types.IncreaseShares(
			project.SpecialAllocations.GenesisDistribution,
			sample.Shares(r),
		)
		tk.ProjectKeeper.SetProject(ctx, project)

		msg, broken := keeper.ProjectSharesInvariant(*tk.ProjectKeeper)(ctx)
		require.True(t, broken, msg)
	})
}
