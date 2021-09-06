package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	campaign "github.com/tendermint/spn/x/campaign/types"
	"testing"
)

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
