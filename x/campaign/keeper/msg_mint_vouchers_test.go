package keeper_test

import (
	"testing"

	spntypes "github.com/tendermint/spn/pkg/types"

	testkeeper "github.com/tendermint/spn/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgMintVouchers(t *testing.T) {
	var (
		sdkCtx, tk, ts = testkeeper.NewTestSetup(t)
		ctx            = sdk.WrapSDKContext(sdkCtx)

		coord           = sample.Address()
		coordNoCampaign = sample.Address()

		shares, _    = types.NewShares("1000foo,500bar,300foobar")
		sharesTooBig = types.NewSharesFromCoins(sdk.NewCoins(
			sdk.NewCoin("foo", sdk.NewInt(spntypes.TotalShareNumber+1)),
		))
	)

	// Create coordinators
	res, err := ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coord,
		Description: sample.CoordinatorDescription(),
	})
	require.NoError(t, err)
	coordID := res.CoordinatorID
	res, err = ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     coordNoCampaign,
		Description: sample.CoordinatorDescription(),
	})
	require.NoError(t, err)

	// Set campaign
	campaign := sample.Campaign(0)
	campaign.CoordinatorID = coordID
	campaign.CampaignID = tk.CampaignKeeper.AppendCampaign(sdkCtx, campaign)

	for _, tc := range []struct {
		name string
		msg  types.MsgMintVouchers
		err  error
	}{
		{
			name: "mint vouchers",
			msg: types.MsgMintVouchers{
				Coordinator: coord,
				CampaignID:  0,
				Shares:      shares,
			},
		},
		{
			name: "mint same vouchers again",
			msg: types.MsgMintVouchers{
				Coordinator: coord,
				CampaignID:  0,
				Shares:      shares,
			},
		},
		{
			name: "mint other vouchers",
			msg: types.MsgMintVouchers{
				Coordinator: coord,
				CampaignID:  0,
				Shares:      sample.Shares(),
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
			name: "non existing campaign",
			msg: types.MsgMintVouchers{
				Coordinator: coord,
				CampaignID:  1000,
				Shares:      shares,
			},
			err: types.ErrCampaignNotFound,
		},
		{
			name: "non existing coordinator",
			msg: types.MsgMintVouchers{
				Coordinator: sample.Address(),
				CampaignID:  0,
				Shares:      shares,
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "invalid coordinator",
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
