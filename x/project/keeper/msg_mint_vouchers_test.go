package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	spntypes "github.com/tendermint/spn/pkg/types"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/project/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgMintVouchers(t *testing.T) {
	var (
		sdkCtx, tk, ts  = testkeeper.NewTestSetup(t)
		ctx             = sdk.WrapSDKContext(sdkCtx)
		coordID         uint64
		coord           = sample.Address(r)
		coordNoCampaign = sample.Address(r)

		shares, _    = types.NewShares("1000foo,500bar,300foobar")
		sharesTooBig = types.NewSharesFromCoins(sdk.NewCoins(
			sdk.NewCoin("foo", sdkmath.NewInt(spntypes.TotalShareNumber+1)),
		))
	)

	t.Run("should allow creation of coordinators", func(t *testing.T) {
		res, err := ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
			Address:     coord,
			Description: sample.CoordinatorDescription(r),
		})
		require.NoError(t, err)
		coordID = res.CoordinatorID
		res, err = ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
			Address:     coordNoCampaign,
			Description: sample.CoordinatorDescription(r),
		})
		require.NoError(t, err)
	})

	// Set campaign
	campaign := sample.Campaign(r, 0)
	campaign.CoordinatorID = coordID
	campaign.CampaignID = tk.CampaignKeeper.AppendCampaign(sdkCtx, campaign)

	for _, tc := range []struct {
		name string
		msg  types.MsgMintVouchers
		err  error
	}{
		{
			name: "should allow minting  vouchers",
			msg: types.MsgMintVouchers{
				Coordinator: coord,
				CampaignID:  0,
				Shares:      shares,
			},
		},
		{
			name: "should allow minting same vouchers again",
			msg: types.MsgMintVouchers{
				Coordinator: coord,
				CampaignID:  0,
				Shares:      shares,
			},
		},
		{
			name: "should allow minting other vouchers",
			msg: types.MsgMintVouchers{
				Coordinator: coord,
				CampaignID:  0,
				Shares:      sample.Shares(r),
			},
		},
		{
			name: "should not mint more than total shares",
			msg: types.MsgMintVouchers{
				Coordinator: coord,
				CampaignID:  0,
				Shares:      sharesTooBig,
			},
			err: types.ErrTotalSharesLimit,
		},
		{
			name: "should fail with non existing campaign",
			msg: types.MsgMintVouchers{
				Coordinator: coord,
				CampaignID:  1000,
				Shares:      shares,
			},
			err: types.ErrCampaignNotFound,
		},
		{
			name: "should fail with non existing coordinator",
			msg: types.MsgMintVouchers{
				Coordinator: sample.Address(r),
				CampaignID:  0,
				Shares:      shares,
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "should fail with invalid coordinator",
			msg: types.MsgMintVouchers{
				Coordinator: coordNoCampaign,
				CampaignID:  0,
				Shares:      shares,
			},
			err: profiletypes.ErrCoordInvalid,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var previousCampaign types.Campaign
			var previousBalance sdk.Coins

			coordAddr, err := sdk.AccAddressFromBech32(tc.msg.Coordinator)
			require.NoError(t, err)

			// Get values before message execution
			if tc.err == nil {
				var found bool
				previousCampaign, found = tk.CampaignKeeper.GetCampaign(sdkCtx, tc.msg.CampaignID)
				require.True(t, found)

				previousBalance = tk.BankKeeper.GetAllBalances(sdkCtx, coordAddr)
			}

			// Execute message
			_, err = ts.CampaignSrv.MintVouchers(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			campaign, found := tk.CampaignKeeper.GetCampaign(sdkCtx, tc.msg.CampaignID)
			require.True(t, found)

			// Allocated shares of the campaign must be increased
			expectedShares := types.IncreaseShares(previousCampaign.AllocatedShares, tc.msg.Shares)
			require.True(t, types.IsEqualShares(expectedShares, campaign.AllocatedShares))

			// Check coordinator balance
			minted, err := types.SharesToVouchers(tc.msg.Shares, tc.msg.CampaignID)
			require.NoError(t, err)
			balance := tk.BankKeeper.GetAllBalances(sdkCtx, coordAddr)
			require.True(t, balance.IsEqual(previousBalance.Add(minted...)))
		})
	}
}
