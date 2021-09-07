package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	campaign "github.com/tendermint/spn/x/campaign/types"
	"testing"
)

const invalidCampaignName = "not_valid"

func TestNewCampaign(t *testing.T) {
	campaignName := sample.CampaignName()
	coordinator := sample.Uint64()
	totalSupply := sample.Coins()
	dynamicShares := sample.Bool()

	cmpn := campaign.NewCampaign(campaignName, coordinator, totalSupply, dynamicShares)
	require.EqualValues(t, campaignName, cmpn.CampaignName)
	require.EqualValues(t, coordinator, cmpn.CoordinatorID)
	require.False(t, cmpn.MainnetInitialized)
	require.True(t, totalSupply.IsEqual(cmpn.TotalSupply))
	require.EqualValues(t, dynamicShares, cmpn.DynamicShares)
	require.EqualValues(t, campaign.EmptyShares(), cmpn.AllocatedShares)
	require.EqualValues(t, campaign.EmptyShares(), cmpn.TotalShares)
}

func TestCampaign_Validate(t *testing.T) {
	invalidCoins := sdk.Coins{sdk.Coin{Denom: "invalid denom", Amount: sdk.NewInt(0)}}
	require.False(t, invalidCoins.IsValid())

	invalidAllocatedShares := sample.Campaign()
	invalidAllocatedShares.AllocatedShares = campaign.NewSharesFromCoins(invalidCoins)

	invalidTotalShares := sample.Campaign()
	invalidTotalShares.DynamicShares = true
	invalidTotalShares.TotalShares = campaign.NewSharesFromCoins(invalidCoins)

	noDynamicShares := sample.Campaign()
	noDynamicShares.DynamicShares = false
	noDynamicShares.TotalShares = sample.Shares()

	totalSharesReached := sample.Campaign()
	totalSharesReached.AllocatedShares = campaign.NewSharesFromCoins(sdk.NewCoins(
		sdk.NewCoin("foo", sdk.NewInt(campaign.DefaultTotalShareNumber+1)),
	))
	require.True(t, campaign.IsTotalSharesReached(totalSharesReached.AllocatedShares, campaign.EmptyShares()))

	for _, tc := range []struct {
		desc     string
		campaign campaign.Campaign
		valid    bool
	}{
		{
			desc:     "valid campaign",
			campaign: sample.Campaign(),
			valid:    true,
		},
		{
			desc: "invalid campaign name",
			campaign: campaign.NewCampaign(
				invalidCampaignName,
				sample.Uint64(),
				sample.Coins(),
				false,
			),
			valid: false,
		},
		{
			desc: "invalid total supply",
			campaign: campaign.NewCampaign(
				sample.CampaignName(),
				sample.Uint64(),
				invalidCoins,
				false,
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
			name:  sample.String(campaign.CampaignNameMaxLength + 1),
			valid: false,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			require.EqualValues(t, tc.valid, campaign.CheckCampaignName(tc.name) == nil)
		})
	}
}
