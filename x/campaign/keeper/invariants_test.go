package keeper_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	tc "github.com/tendermint/spn/testutil/constructor"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/keeper"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestAccountWithoutCampaignInvariant(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("valid case", func(t *testing.T) {
		campaign := sample.Campaign(r, 0)
		campaign.CampaignID = tk.CampaignKeeper.AppendCampaign(ctx, campaign)
		tk.CampaignKeeper.SetMainnetAccount(ctx, sample.MainnetAccount(r, campaign.CampaignID, sample.Address(r)))
		msg, broken := keeper.AccountWithoutCampaignInvariant(*tk.CampaignKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("invalid case", func(t *testing.T) {
		tk.CampaignKeeper.SetMainnetAccount(ctx, sample.MainnetAccount(r, 100, sample.Address(r)))
		msg, broken := keeper.AccountWithoutCampaignInvariant(*tk.CampaignKeeper)(ctx)
		require.True(t, broken, msg)
	})
}

func TestCampaignSharesInvariant(t *testing.T) {
	t.Run("valid case", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		// create campaigns with some allocated shares
		campaignID1, campaignID2 := uint64(1), uint64(2)
		campaign := sample.Campaign(r, campaignID1)
		campaign.AllocatedShares = types.IncreaseShares(
			campaign.AllocatedShares,
			tc.Shares(t, "100foo,200bar"),
		)
		tk.CampaignKeeper.SetCampaign(ctx, campaign)

		campaign = sample.Campaign(r, campaignID2)
		campaign.AllocatedShares = types.IncreaseShares(
			campaign.AllocatedShares,
			tc.Shares(t, "10000foo"),
		)
		tk.CampaignKeeper.SetCampaign(ctx, campaign)

		// mint vouchers
		voucherFoo, voucherBar := types.VoucherDenom(campaignID1, "foo"), types.VoucherDenom(campaignID1, "bar")
		tk.Mint(ctx, sample.Address(r), tc.Coins(t, fmt.Sprintf("50%s,100%s", voucherFoo, voucherBar)))

		// mint vouchers for another campaign
		voucherFoo = types.VoucherDenom(campaignID2, "foo")
		tk.Mint(ctx, sample.Address(r), tc.Coins(t, fmt.Sprintf("5000%s", voucherFoo)))

		// add accounts with shares
		tk.CampaignKeeper.SetMainnetAccount(ctx, types.MainnetAccount{
			CampaignID: campaignID1,
			Address:    sample.Address(r),
			Shares:     tc.Shares(t, "20foo,40bar"),
		})
		tk.CampaignKeeper.SetMainnetAccount(ctx, types.MainnetAccount{
			CampaignID: campaignID1,
			Address:    sample.Address(r),
			Shares:     tc.Shares(t, "30foo,60bar"),
		})
		tk.CampaignKeeper.SetMainnetAccount(ctx, types.MainnetAccount{
			CampaignID: campaignID2,
			Address:    sample.Address(r),
			Shares:     tc.Shares(t, "5000foo"),
		})

		msg, broken := keeper.CampaignSharesInvariant(*tk.CampaignKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("campaign with empty allocated share is valid", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		tk.CampaignKeeper.SetCampaign(ctx, sample.Campaign(r, 3))

		msg, broken := keeper.CampaignSharesInvariant(*tk.CampaignKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("allocated shares cannot be converted to vouchers", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		campaignID := uint64(4)
		campaign := sample.Campaign(r, campaignID)
		coins := tc.Coins(t, "100foo,200bar")
		shares := make(types.Shares, len(coins))
		for i, coin := range coins {
			shares[i] = coin
		}
		campaign.AllocatedShares = types.IncreaseShares(
			campaign.AllocatedShares,
			shares,
		)
		tk.CampaignKeeper.SetCampaign(ctx, campaign)

		msg, broken := keeper.CampaignSharesInvariant(*tk.CampaignKeeper)(ctx)
		require.True(t, broken, msg)
	})

	t.Run("invalid allocated shares", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		campaignID := uint64(4)
		campaign := sample.Campaign(r, campaignID)
		campaign.AllocatedShares = types.IncreaseShares(
			campaign.AllocatedShares,
			tc.Shares(t, "100foo,200bar"),
		)
		tk.CampaignKeeper.SetCampaign(ctx, campaign)

		// mint vouchers
		voucherFoo, voucherBar := types.VoucherDenom(campaignID, "foo"), types.VoucherDenom(campaignID, "bar")
		tk.Mint(ctx, sample.Address(r), tc.Coins(t, fmt.Sprintf("99%s,200%s", voucherFoo, voucherBar)))

		msg, broken := keeper.CampaignSharesInvariant(*tk.CampaignKeeper)(ctx)
		require.True(t, broken, msg)
	})

	t.Run("campaign with special allocations not tracked by allocated shares", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		campaign := sample.Campaign(r, 3)
		campaign.SpecialAllocations.GenesisDistribution = types.IncreaseShares(
			campaign.SpecialAllocations.GenesisDistribution,
			sample.Shares(r),
		)
		tk.CampaignKeeper.SetCampaign(ctx, campaign)

		msg, broken := keeper.CampaignSharesInvariant(*tk.CampaignKeeper)(ctx)
		require.True(t, broken, msg)
	})
}
