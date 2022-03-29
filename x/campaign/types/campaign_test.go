package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	spntypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/testutil/sample"
	campaign "github.com/tendermint/spn/x/campaign/types"
)

var (
	invalidCampaignName = "not_valid"
	invalidCoins        = sdk.Coins{sdk.Coin{Denom: "invalid denom", Amount: sdk.ZeroInt()}}
	invalidShares       = campaign.Shares{sdk.Coin{Denom: "invalid denom", Amount: sdk.ZeroInt()}}
)

func TestNewCampaign(t *testing.T) {
	campaignID := sample.Uint64(r)
	campaignName := sample.CampaignName(r)
	coordinator := sample.Uint64(r)
	totalSupply := sample.TotalSupply(r)
	dynamicShares := sample.Bool(r)
	metadata := sample.Metadata(r, 20)

	cmpn := campaign.NewCampaign(campaignID, campaignName, coordinator, totalSupply, dynamicShares, metadata)
	require.EqualValues(t, campaignID, cmpn.CampaignID)
	require.EqualValues(t, campaignName, cmpn.CampaignName)
	require.EqualValues(t, coordinator, cmpn.CoordinatorID)
	require.False(t, cmpn.MainnetInitialized)
	require.True(t, totalSupply.IsEqual(cmpn.TotalSupply))
	require.EqualValues(t, dynamicShares, cmpn.DynamicShares)
	require.EqualValues(t, campaign.EmptyShares(), cmpn.AllocatedShares)
	require.EqualValues(t, campaign.EmptyShares(), cmpn.TotalShares)
}

func TestCampaign_Validate(t *testing.T) {
	require.False(t, invalidCoins.IsValid())

	invalidAllocatedShares := sample.Campaign(r, 0)
	invalidAllocatedShares.AllocatedShares = campaign.NewSharesFromCoins(invalidCoins)

	invalidTotalShares := sample.Campaign(r, 0)
	invalidTotalShares.DynamicShares = true
	invalidTotalShares.TotalShares = campaign.NewSharesFromCoins(invalidCoins)

	noDynamicShares := sample.Campaign(r, 0)
	noDynamicShares.DynamicShares = false
	noDynamicShares.TotalShares = sample.Shares(r)

	totalSharesReached := sample.Campaign(r, 0)
	totalSharesReached.AllocatedShares = campaign.NewSharesFromCoins(sdk.NewCoins(
		sdk.NewCoin("foo", sdk.NewInt(spntypes.TotalShareNumber+1)),
	))
	require.True(t, campaign.IsTotalSharesReached(totalSharesReached.AllocatedShares, campaign.EmptyShares(),
		spntypes.TotalShareNumber))

	for _, tc := range []struct {
		desc     string
		campaign campaign.Campaign
		valid    bool
	}{
		{
			desc:     "valid campaign",
			campaign: sample.Campaign(r, 0),
			valid:    true,
		},
		{
			desc: "invalid campaign name",
			campaign: campaign.NewCampaign(
				0,
				invalidCampaignName,
				sample.Uint64(r),
				sample.TotalSupply(r),
				false,
				sample.Metadata(r, 20),
			),
			valid: false,
		},
		{
			desc: "invalid total supply",
			campaign: campaign.NewCampaign(
				0,
				sample.CampaignName(r),
				sample.Uint64(r),
				invalidCoins,
				false,
				sample.Metadata(r, 20),
			),
			valid: false,
		},
		{
			desc:     "invalid allocated shares",
			campaign: invalidAllocatedShares,
			valid:    false,
		},
		{
			desc:     "invalid total shares",
			campaign: invalidTotalShares,
			valid:    false,
		},
		{
			desc:     "total shares can't be set if no dynamic shares",
			campaign: noDynamicShares,
			valid:    false,
		},
		{
			desc:     "allocated shares bigger than total shares",
			campaign: totalSharesReached,
			valid:    false,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			require.EqualValues(t, tc.valid, tc.campaign.Validate() == nil)
		})
	}
}

func TestCheckCampaignName(t *testing.T) {
	for _, tc := range []struct {
		desc  string
		name  string
		valid bool
	}{
		{
			desc:  "valid name",
			name:  "ThisIs-a-ValidCampaignName123",
			valid: true,
		},
		{
			desc:  "should not contain special character outside hyphen",
			name:  invalidCampaignName,
			valid: false,
		},
		{
			desc:  "should not be empty",
			name:  "",
			valid: false,
		},
		{
			desc:  "should not exceed max length",
			name:  sample.String(r, campaign.CampaignNameMaxLength+1),
			valid: false,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			require.EqualValues(t, tc.valid, campaign.CheckCampaignName(tc.name) == nil)
		})
	}
}
