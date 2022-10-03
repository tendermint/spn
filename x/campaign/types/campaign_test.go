package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	spntypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/testutil/sample"
	campaign "github.com/tendermint/spn/x/campaign/types"
)

var (
	invalidCampaignName  = "not_valid"
	invalidCampaignCoins = sdk.Coins{sdk.Coin{Denom: "invalid denom", Amount: sdkmath.ZeroInt()}}
)

func TestNewCampaign(t *testing.T) {
	campaignID := sample.Uint64(r)
	campaignName := sample.CampaignName(r)
	coordinator := sample.Uint64(r)
	totalSupply := sample.TotalSupply(r)
	metadata := sample.Metadata(r, 20)
	createdAt := sample.Duration(r).Milliseconds()

	t.Run("should allow creation of campaign", func(t *testing.T) {
		c := campaign.NewCampaign(
			campaignID,
			campaignName,
			coordinator,
			totalSupply,
			metadata,
			createdAt,
		)
		require.EqualValues(t, campaignID, c.CampaignID)
		require.EqualValues(t, campaignName, c.CampaignName)
		require.EqualValues(t, coordinator, c.CoordinatorID)
		require.EqualValues(t, createdAt, c.CreatedAt)
		require.False(t, c.MainnetInitialized)
		require.True(t, totalSupply.IsEqual(c.TotalSupply))
		require.EqualValues(t, campaign.EmptyShares(), c.AllocatedShares)
	})
}

func TestCampaign_Validate(t *testing.T) {
	var (
		invalidAllocatedShares            campaign.Campaign
		totalSharesReached                campaign.Campaign
		campaignInvalidSpecialAllocations campaign.Campaign
	)

	t.Run("should verify that invalid coins is invalid", func(t *testing.T) {
		require.False(t, invalidCampaignCoins.IsValid())
	})

	t.Run("should allow creation of valid allocations with totalshares reached", func(t *testing.T) {
		invalidAllocatedShares = sample.Campaign(r, 0)
		invalidAllocatedShares.AllocatedShares = campaign.NewSharesFromCoins(invalidCampaignCoins)
		totalSharesReached = sample.Campaign(r, 0)
		totalSharesReached.AllocatedShares = campaign.NewSharesFromCoins(sdk.NewCoins(
			sdk.NewCoin("foo", sdkmath.NewInt(spntypes.TotalShareNumber+1)),
		))
		reached, err := campaign.IsTotalSharesReached(totalSharesReached.AllocatedShares, spntypes.TotalShareNumber)
		require.NoError(t, err)
		require.True(t, reached)
	})

	t.Run("should allow creation of campaign with invalid special allocations", func(t *testing.T) {
		invalidSpecialAllocations := campaign.NewSpecialAllocations(
			sample.Shares(r),
			campaign.Shares(sdk.NewCoins(
				sdk.NewCoin("foo", sdkmath.NewInt(100)),
				sdk.NewCoin("s/bar", sdkmath.NewInt(200)),
			)),
		)
		require.Error(t, invalidSpecialAllocations.Validate())
		campaignInvalidSpecialAllocations = sample.Campaign(r, 0)
		campaignInvalidSpecialAllocations.SpecialAllocations = invalidSpecialAllocations
	})

	for _, tc := range []struct {
		desc     string
		campaign campaign.Campaign
		valid    bool
	}{
		{
			desc:     "should allow validation of valid campaign",
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
				sample.Metadata(r, 20),
				sample.Duration(r).Milliseconds(),
			),
			valid: false,
		},
		{
			desc: "should prevent validation of campaign with invalid total supply",
			campaign: campaign.NewCampaign(
				0,
				sample.CampaignName(r),
				sample.Uint64(r),
				invalidCampaignCoins,
				sample.Metadata(r, 20),
				sample.Duration(r).Milliseconds(),
			),
			valid: false,
		},
		{
			desc:     "should prevent validation of campaign with invalid allocated shares",
			campaign: invalidAllocatedShares,
			valid:    false,
		},
		{
			desc:     "should prevent validation of campaign with allocated shares greater than total shares",
			campaign: totalSharesReached,
			valid:    false,
		},
		{
			desc:     "should prevent validation of campaign with invalid special allocations",
			campaign: campaignInvalidSpecialAllocations,
			valid:    false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			require.EqualValues(t, tc.valid, tc.campaign.Validate(spntypes.TotalShareNumber) == nil)
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
			desc:  "should allow check of campaign with valid name",
			name:  "ThisIs-a-ValidCampaignName123",
			valid: true,
		},
		{
			desc:  "should prevent check of campaign with special character outside hyphen",
			name:  invalidCampaignName,
			valid: false,
		},
		{
			desc:  "should prevent check of campaign with empty name",
			name:  "",
			valid: false,
		},
		{
			desc:  "should prevent check of campaign with name exceeding max length",
			name:  sample.String(r, campaign.CampaignNameMaxLength+1),
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			require.EqualValues(t, tc.valid, campaign.CheckCampaignName(tc.name) == nil)
		})
	}
}
