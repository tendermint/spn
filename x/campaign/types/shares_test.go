package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	campaign "github.com/tendermint/spn/x/campaign/types"
	"testing"
)

func TestEmptyShares(t *testing.T) {
	shares := campaign.EmptyShares()
	require.Equal(t, shares, campaign.Shares(sdk.NewCoins()))
}

func TestNewShares(t *testing.T) {
	_, err := campaign.NewShares("invalid")
	require.Error(t, err)

	shares, err := campaign.NewShares("100foo,200bar")
	require.NoError(t, err)
	require.Equal(t, shares, campaign.Shares(sdk.NewCoins(
		sdk.NewCoin(campaign.SharePrefix + "foo", sdk.NewInt(100)),
		sdk.NewCoin(campaign.SharePrefix + "bar", sdk.NewInt(200)),
	)))
}

func TestNewSharesFromCoins(t *testing.T) {
	shares := campaign.NewSharesFromCoins(sdk.NewCoins(
		sdk.NewCoin("foo", sdk.NewInt(100)),
		sdk.NewCoin("bar", sdk.NewInt(200)),
	))
	require.Equal(t, shares, campaign.Shares(sdk.NewCoins(
		sdk.NewCoin(campaign.SharePrefix + "foo", sdk.NewInt(100)),
		sdk.NewCoin(campaign.SharePrefix + "bar", sdk.NewInt(200)),
	)))
}

func TestCheckShares(t *testing.T) {
	require.NoError(t, campaign.CheckShares(campaign.Shares(sdk.NewCoins(
		sdk.NewCoin(campaign.SharePrefix + "foo", sdk.NewInt(100)),
		sdk.NewCoin(campaign.SharePrefix + "bar", sdk.NewInt(200)),
	))))
	require.NoError(t, campaign.CheckShares(campaign.Shares(sdk.NewCoins(
		sdk.NewCoin("foo", sdk.NewInt(100)),
		sdk.NewCoin(campaign.SharePrefix + "bar", sdk.NewInt(200)),
	))))
}

func TestIsTotalReached(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		shares campaign.Shares
		totalShares campaign.Shares
		reached    bool
	}{
		{
			desc:     "empty is false",
			shares: campaign.EmptyShares(),
			totalShares: campaign.EmptyShares(),
			reached: false,
		},
		{
			desc:     "no default total is reached",
			shares: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin("foo", sdk.NewInt(campaign.DefaultTotalShareNumber)),
				sdk.NewCoin("bar", sdk.NewInt(100)),
				)),
			totalShares: campaign.EmptyShares(),
			reached: false,
		},
		{
			desc:     "no custom total is reached",
			shares: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin("foo", sdk.NewInt(100)),
				sdk.NewCoin("bar", sdk.NewInt(50)),
			)),
			totalShares: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin("foo", sdk.NewInt(100)),
				sdk.NewCoin("bar", sdk.NewInt(100)),
				sdk.NewCoin("foobar", sdk.NewInt(100)),
			)),
			reached: false,
		},
		{
			desc:     "a default total is reached",
			shares: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin("foo", sdk.NewInt(campaign.DefaultTotalShareNumber+1)),
				sdk.NewCoin("bar", sdk.NewInt(100)),
			)),
			totalShares: campaign.EmptyShares(),
			reached: true,
		},
		{
			desc:     "a custom total is reached",
			shares: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin("foo", sdk.NewInt(campaign.DefaultTotalShareNumber)),
				sdk.NewCoin("bar", sdk.NewInt(101)),
			)),
			totalShares: campaign.NewSharesFromCoins(sdk.NewCoins(
				sdk.NewCoin("bar", sdk.NewInt(100)),
			)),
			reached: true,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			require.Equal(t, tc.reached, campaign.IsTotalReached(tc.shares, tc.totalShares))
		})
	}
}