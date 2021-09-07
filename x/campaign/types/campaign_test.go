package types_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	campaign "github.com/tendermint/spn/x/campaign/types"
	"testing"
)

func TestNewCampaign(t *testing.T) {
	campaignName := sample.CampaignName()
	coordinator := sample.Uint64()
	totalSupply := sample.Coins()
	dynamicShares := sample.Bool()

	cmpn := campaign.NewCampaign(campaignName, coordinator, totalSupply, dynamicShares)
	require.EqualValues(t,campaignName, cmpn.CampaignName)
	require.EqualValues(t,coordinator, cmpn.CoordinatorID)
	require.False(t, cmpn.MainnetInitialized)
	require.True(t, totalSupply.IsEqual(cmpn.TotalSupply))
	require.EqualValues(t,dynamicShares, cmpn.DynamicShares)
	require.EqualValues(t,campaign.EmptyShares(), cmpn.AllocatedShares)
	require.EqualValues(t,campaign.EmptyShares(), cmpn.TotalShares)
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
			name:  "not_valid",
			valid: false,
		},
		{
			desc:  "should not be empty",
			name:  "",
			valid: false,
		},
		{
			desc:  "should not exceed max length",
			name:  sample.String(campaign.CampaignNameMaxLength+1),
			valid: false,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			require.EqualValues(t, tc.valid, campaign.CheckCampaignName(tc.name) == nil)
		})
	}
}